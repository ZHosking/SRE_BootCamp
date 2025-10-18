package models

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// InitDB opens/creates the database and sets up the table
func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		return nil, err
	}

	// Create table if not exists
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS students (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        age INTEGER NOT NULL
    );`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	log.Println("Database initialised")

	return db, nil
}

func GetAllStudents(db *sql.DB) ([]Student, error) {
	rows, err := db.Query("SELECT id, name, age FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Age); err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}

func GetStudentByID(db *sql.DB, id int) (*Student, error) {
	row := db.QueryRow("SELECT id, name, age FROM students WHERE id = ?", id)
	var s Student
	err := row.Scan(&s.ID, &s.Name, &s.Age)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func AddStudent(db *sql.DB, s Student) error {
	_, err := db.Exec("INSERT INTO students (name, age) VALUES (?, ?)", s.Name, s.Age)
	return err
}

func UpdateStudent(db *sql.DB, s Student) error {
	_, err := db.Exec("UPDATE students SET name = ?, age = ? WHERE id = ?", s.Name, s.Age, s.ID)
	return err
}

func DeleteStudent(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM students WHERE id = ?", id)
	return err
}
