package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	Database *sql.DB
}

type Program struct {
	Title          string         `json:"title"`
	Category       string         `json:"category"`
	DaysPerWeek    int            `json:"days_per_week"`
	ProgramDetails ProgramDetails `json:"program_details"`
}

type ProgramDetails map[string]Day

type Day struct {
	Exercises []Exercise `json:"exercises"`
}

type Exercise struct {
	Name string `json:"name"`
	Sets int    `json:"sets"`
	Reps string `json:"reps"`
}

func (db *DB) InitDB() {
	// Replace with your PostgreSQL connection string
	connStr := "user=postgres password=Newpassword dbname=programs sslmode=disable"
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.Database = database

	// Create tables if they don't exist
	initStmt := `
	CREATE TABLE IF NOT EXISTS programs (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		category TEXT NOT NULL,
		days_per_week INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS days (
		id SERIAL PRIMARY KEY,
		program_id INTEGER NOT NULL,
		day_name TEXT NOT NULL,
		FOREIGN KEY (program_id) REFERENCES programs(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS exercises (
		id SERIAL PRIMARY KEY,
		day_id INTEGER NOT NULL,
		exercise_name TEXT NOT NULL,
		sets INTEGER NOT NULL,
		reps TEXT NOT NULL,
		FOREIGN KEY (day_id) REFERENCES days(id) ON DELETE CASCADE
	);`

	_, err = db.Database.Exec(initStmt)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB) AddProgram(program Program) error {
	tx, err := db.Database.Begin()
	if err != nil {
		return err
	}

	programStmt := `INSERT INTO programs (title, category, days_per_week) VALUES ($1, $2, $3) RETURNING id`
	var programID int
	err = tx.QueryRow(programStmt, program.Title, program.Category, program.DaysPerWeek).Scan(&programID)
	if err != nil {
		tx.Rollback()
		return err
	}

	dayStmt := `INSERT INTO days (program_id, day_name) VALUES ($1, $2) RETURNING id`
	exerciseStmt := `INSERT INTO exercises (day_id, exercise_name, sets, reps) VALUES ($1, $2, $3, $4)`

	for dayName, day := range program.ProgramDetails {
		var dayID int
		err := tx.QueryRow(dayStmt, programID, dayName).Scan(&dayID)
		if err != nil {
			tx.Rollback()
			return err
		}

		for _, exercise := range day.Exercises {
			_, err := tx.Exec(exerciseStmt, dayID, exercise.Name, exercise.Sets, exercise.Reps)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit()
}

func (db *DB) GetPrograms() ([]Program, error) {
	programStmt := `SELECT id, title, category, days_per_week FROM programs`
	rows, err := db.Database.Query(programStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var programs []Program

	for rows.Next() {
		var program Program
		var programID int

		err := rows.Scan(&programID, &program.Title, &program.Category, &program.DaysPerWeek)
		if err != nil {
			return nil, err
		}

		program.ProgramDetails = make(ProgramDetails)

		dayStmt := `SELECT id, day_name FROM days WHERE program_id = $1`
		dayRows, err := db.Database.Query(dayStmt, programID)
		if err != nil {
			return nil, err
		}

		for dayRows.Next() {
			var dayID int
			var dayName string
			var day Day

			err := dayRows.Scan(&dayID, &dayName)
			if err != nil {
				return nil, err
			}

			exerciseStmt := `SELECT exercise_name, sets, reps FROM exercises WHERE day_id = $1`
			exerciseRows, err := db.Database.Query(exerciseStmt, dayID)
			if err != nil {
				return nil, err
			}

			for exerciseRows.Next() {
				var exercise Exercise
				err := exerciseRows.Scan(&exercise.Name, &exercise.Sets, &exercise.Reps)
				if err != nil {
					exerciseRows.Close()
					return nil, err
				}
				day.Exercises = append(day.Exercises, exercise)
			}
			exerciseRows.Close()

			program.ProgramDetails[dayName] = day
		}
		dayRows.Close()

		programs = append(programs, program)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return programs, nil
}

func (db *DB) GetProgram(programID int) (Program, error) {
	programStmt := `SELECT id, title, category, days_per_week FROM programs WHERE id = $1`
	row := db.Database.QueryRow(programStmt, programID)

	var program Program
	err := row.Scan(&program.Title, &program.Category, &program.DaysPerWeek)
	if err != nil {
		return Program{}, err
	}

	program.ProgramDetails = make(ProgramDetails)

	dayStmt := `SELECT id, day_name FROM days WHERE program_id = $1`
	dayRows, err := db.Database.Query(dayStmt, programID)
	if err != nil {
		return Program{}, err
	}

	for dayRows.Next() {
		var dayID int
		var dayName string
		var day Day

		err := dayRows.Scan(&dayID, &dayName)
		if err != nil {
			return Program{}, err
		}

		exerciseStmt := `SELECT exercise_name, sets, reps FROM exercises WHERE day_id = $1`
		exerciseRows, err := db.Database.Query(exerciseStmt, dayID)
		if err != nil {
			return Program{}, err
		}

		for exerciseRows.Next() {
			var exercise Exercise
			err := exerciseRows.Scan(&exercise.Name, &exercise.Sets, &exercise.Reps)
			if err != nil {
				exerciseRows.Close()
				return Program{}, err
			}
			day.Exercises = append(day.Exercises, exercise)
		}
		exerciseRows.Close()

		program.ProgramDetails[dayName] = day
	}
	dayRows.Close()
	return program, nil
}
