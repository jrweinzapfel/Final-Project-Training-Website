package db

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *DB) {
	// Define routes
	router.GET("/programs", func(c *gin.Context) {
		programs, err := db.GetPrograms()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, programs)
	})

	router.GET("/program/:id", func(c *gin.Context) {
		id := c.Param("id")
		programID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid program ID"})
			return
		}

		program, err := db.GetProgram(programID)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				c.JSON(404, gin.H{"error": "Program not found"})
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(200, program)
	})
}
