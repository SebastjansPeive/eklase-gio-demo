// Package storage is an interface for interacting with a database.
package storage

import (
	"fmt"
	"log"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

var (
	// Statement for creating tables. Currently creates `students` table only.
	// In the future can be expanded with creation of other tables.
	createTableStmt = `

	CREATE TABLE IF NOT EXISTS students (
		id	INTEGER,
		name	TEXT,
		surname	TEXT,
		PRIMARY KEY(id AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS classes (
		id	INTEGER,
		year INTEGER,
		modifier	TEXT,
		PRIMARY KEY(id AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS groups (
		student_id	INTEGER,
		name TEXT,
		surname TEXT,
		year INTEGER,
		modifier	TEXT,
		PRIMARY KEY(student_id)
	);`
	// Statement for adding a new entry into `students` table.
	insertStudentsStmt = `INSERT INTO students (name, surname) VALUES(?, ?);
	INSERT INTO groups (name, surname) VALUES(?, ?);`
	// Statement for getting all entries from `students` table.
	selectStudentsStmt = `SELECT id, name, surname FROM students`
	insertClassesStmt  = `INSERT INTO classes (year, modifier) VALUES(?, ?)`
	selectClassesStmt  = `SELECT id, year, modifier FROM classes`
	selectGroupsStmt   = `SELECT name, surname, year, modifier FROM groups`
	assignClassToStudentStmt = `UPDATE groups SET year = ?, modifier = ? WHERE student_id = ?`
)

// StudentEntry represents a row for a single student in the DB.
type StudentEntry struct {
	ID 		int    `db:"id"`
	Name    string `db:"name"`
	Surname string `db:"surname"`
}

type ClassEntry struct {
	ID 		 int 	`db:"id"`
	Year     string `db:"year"`
	Modifier string `db:"modifier"`
}

type GroupEntry struct {
	StudentID int    `db:"student_id"`
	Name      sql.NullString `db:"name"`
	Surname   sql.NullString `db:"surname"`
	ClassID   int    `db:"class_id"`
	Year      sql.NullString `db:"year"`
	Modifier  sql.NullString `db:"modifier"`
}

// Storage is an interface for interacting with persistent storage.
type Storage struct {
	db *sqlx.DB
}

// New initializes a new DB given its path, or opens an existing DB, and
// initializes the handler. Returns an error if any of the steps fails.
func New(path string) (*Storage, error) {
	// Open a DB by the path.
	db, err := sqlx.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite DB: %v", db)
	}

	// Create new tables. Note that the tables may exist already.
	res, err := db.Exec(createTableStmt)
	if err != nil {
		return nil, fmt.Errorf("table creation failed. Query: %v\nError: %v", createTableStmt, err)
	}
	if cnt, err := res.RowsAffected(); err != nil {
		log.Printf("%d rows affected.", cnt)
	}

	return &Storage{db: db}, nil
}

func Must(s *Storage, err error) *Storage {
	if err != nil {
		log.Fatalf("unable to create storage: %v", err)
	}
	return s
}

// Close closes the database after it is no longer required.
func (s *Storage) Close() error {
	return s.db.Close()
}

// Students returns a slice of existing students.
func (s Storage) Students() ([]StudentEntry, error) {
	var entries []StudentEntry
	// Read rows from the `students` table and populate students field in the
	// handler.
	if err := s.db.Select(&entries, selectStudentsStmt); err != nil {
		return nil, fmt.Errorf("querying 'students' table failed. Query: %v\nError: %v", selectStudentsStmt, err)
	}
	return entries, nil
}

func (s Storage) Classes() ([]ClassEntry, error) {
	var entries2 []ClassEntry
	if err := s.db.Select(&entries2, selectClassesStmt); err != nil {
		return nil, fmt.Errorf("querying 'classes' table failed. Query: %v\nError: %v", selectClassesStmt, err)
	}
	return entries2, nil
}

func (s Storage) Groups() ([]GroupEntry, error) {
	var entries []GroupEntry
	if err := s.db.Select(&entries, selectGroupsStmt); err != nil {
		return nil, fmt.Errorf("querying 'groups' table failed. Query: %v\nError: %v", selectGroupsStmt, err)
	}
	return entries, nil
}

// AddStudent appends a new student entry to the database.
func (s *Storage) AddStudent(name, surname string) error {
	// Attempt to add an entry to the database first.
	// If it fails, the student field will not be modified.
	res, err := s.db.Exec(insertStudentsStmt, name, surname)
	if err != nil {
		return fmt.Errorf("table creation failed. Query: %v\nError: %v", createTableStmt, err)
	}
	if cnt, err := res.RowsAffected(); err != nil {
		log.Printf("%d rows affected.", cnt)
	}
	return nil
}

func (s *Storage) AddClass(year, modifier string) error {
	res, err := s.db.Exec(insertClassesStmt, year, modifier)
	if err != nil {
		return fmt.Errorf("table creation failed. Query: %v\nError: %v", createTableStmt, err)
	}
	if cnt, err := res.RowsAffected(); err != nil {
		log.Printf("%d rows affected.", cnt)
	}
	return nil
}

func (s *Storage) AssignClassToStudent(year, modifier string, student_id int) error {
	res, err := s.db.Exec(assignClassToStudentStmt, year, modifier, student_id)
	if err != nil {
		return fmt.Errorf("table creation failed. Query: %v\nError: %v", createTableStmt, err)
	}
	if cnt, err := res.RowsAffected(); err != nil {
		log.Printf("%d rows affected.", cnt)
	}
	return nil
}
