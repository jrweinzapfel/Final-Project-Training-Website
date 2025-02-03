package main

import (
	"log"

	"mattn-sqlite3/db"
)

func main() {
	database := &db.DB{}
	database.InitDB()

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
}
