package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
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
	database, err := sql.Open("sqlite3", "./programs.db")

	if err != nil {
		log.Fatal(err)
	}

	db.Database = database

	initStmt := `CREATE TABLE IF NOT EXISTS programs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL, 
		category TEXT NOT NULL, 
		days_per_week INTEGER NOT NULL);
		
		CREATE TABLE IF NOT EXISTS days (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		program_id INTEGER NOT NULL, 
		day_name TEXT NOT NULL, 
		FOREIGN KEY (program_id) REFERENCES programs(id)
		);
		
		CREATE TABLE IF NOT EXISTS exercises (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		day_id INTEGER NOT NULL,
		exercise_name TEXT NOT NULL, 
		sets INTEGER NOT NULL, 
		reps TEXT NOT NULL, 
		FOREIGN KEY (day_id) REFERENCES days(id))`

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

	programStmt := `INSERT INTO programs (title, category, days_per_week) VALUES (?, ?, ?)`
	res, err := tx.Exec(programStmt, program.Title, program.Category, program.DaysPerWeek)
	if err != nil {
		tx.Rollback()
		return err
	}

	programID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	dayStmt := `INSERT INTO days (program_id, day_name) VALUES (?, ?)`
	exerciseStmt := `INSERT INTO exercises (day_id, exercise_name, sets, reps) VALUES (?, ?, ?, ?)`

	for dayName, day := range program.ProgramDetails {
		res, err := tx.Exec(dayStmt, programID, dayName)
		if err != nil {
			tx.Rollback()
			return err
		}

		dayID, err := res.LastInsertId()
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

		dayStmt := `SELECT id, day_name FROM days WHERE program_id = ?`
		dayRows, err := db.Database.Query(dayStmt, programID)
		if err != nil {
			return nil, err
		}
		defer dayRows.Close()

		for dayRows.Next() {
			var dayID int
			var dayName string
			var day Day

			err := dayRows.Scan(&dayID, &dayName)
			if err != nil {
				return nil, err
			}

			exerciseStmt := `SELECT exercise_name, sets, reps FROM exercises WHERE day_id = ?`
			exerciseRows, err := db.Database.Query(exerciseStmt, dayID)
			if err != nil {
				return nil, err
			}
			defer exerciseRows.Close()

			for exerciseRows.Next() {
				var exercise Exercise

				err := exerciseRows.Scan(&exercise.Name, &exercise.Sets, &exercise.Reps)
				if err != nil {
					return nil, err
				}

				day.Exercises = append(day.Exercises, exercise)
			}

			program.ProgramDetails[dayName] = day
		}

		programs = append(programs, program)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return programs, nil
}

func (db *DB) GetProgram(programID int) (Program, error) {
	programStmt := `SELECT id, title, category, days_per_week FROM programs WHERE id = ?`
	row := db.Database.QueryRow(programStmt, programID)

	var program Program
	err := row.Scan(&program.Title, &program.Category, &program.DaysPerWeek)
	if err != nil {
		return Program{}, err
	}

	program.ProgramDetails = make(ProgramDetails)

	dayStmt := `SELECT id, day_name FROM days WHERE program_id = ?`
	dayRows, err := db.Database.Query(dayStmt, programID)
	if err != nil {
		return Program{}, err
	}
	defer dayRows.Close()

	for dayRows.Next() {
		var dayID int
		var dayName string
		var day Day

		err := dayRows.Scan(&dayID, &dayName)
		if err != nil {
			return Program{}, err
		}

		exerciseStmt := `SELECT exercise_name, sets, reps FROM exercises WHERE day_id = ?`
		exerciseRows, err := db.Database.Query(exerciseStmt, dayID)
		if err != nil {
			return Program{}, err
		}
		defer exerciseRows.Close()

		for exerciseRows.Next() {
			var exercise Exercise
			err := exerciseRows.Scan(&exercise.Name, &exercise.Sets, &exercise.Reps)
			if err != nil {
				return Program{}, err
			}
			day.Exercises = append(day.Exercises, exercise)
		}
		program.ProgramDetails[dayName] = day
	}
	return program, nil
}
