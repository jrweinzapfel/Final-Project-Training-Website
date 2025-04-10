package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"mattn-sqlite3/db"
)

func main() {
	// Initialize the database
	database := &db.DB{}
	database.InitDB()

	// Create a sample program and add it to the database
	program := db.Program{
		Title:       "12-Week Hypertrophy Program",
		Category:    "Hypertrophy",
		DaysPerWeek: 4,
		ProgramDetails: db.ProgramDetails{
			"Day 1: Upper Body Push": db.Day{
				Exercises: []db.Exercise{
					{Name: "Barbell Bench Press", Sets: 4, Reps: "8-12"},
					{Name: "Incline Dumbbell Press", Sets: 3, Reps: "10-12"},
					{Name: "Overhead Shoulder Press (Barbell)", Sets: 3, Reps: "8-10"},
				},
			},
		},
	}

	err := database.AddProgram(program)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Register routes
	db.RegisterRoutes(r, database)

	// Start the server
	r.Run(":8080")
}
