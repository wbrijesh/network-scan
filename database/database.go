package database

import (
	"database/sql"
	"fmt"
	"os"
	"scan-network/utils"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func CreateDbIfNotExists(dbPath string) error {
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		utils.PrintInColor("Database file does not exist, creating it...", 136)
		file, err := os.Create(dbPath)
		if err != nil {
			return fmt.Errorf("error creating database file: %w", err)
		}
		file.Close()
	}
	return nil
}

func New(dbPath string) (*Database, error) {
	if err := CreateDbIfNotExists(dbPath); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) CreateOuiTable() error {
	_, err := d.db.Exec(`
		CREATE TABLE IF NOT EXISTS oui (
			id TEXT PRIMARY KEY,
			assignment TEXT UNIQUE,
			organisation TEXT
		)
	`)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}
	return nil
}

func (d *Database) InsertOUI(assignment, organisation string) error {
	id := utils.GenerateBUID()
	_, err := d.db.Exec("INSERT OR REPLACE INTO oui (id, assignment, organisation) VALUES (?, ?, ?)",
		id, assignment, organisation)
	if err != nil {
		return fmt.Errorf("error inserting OUI: %w", err)
	}
	return nil
}

func (d *Database) FindOrganisationByAssignment(assignment string) (string, error) {
	var organisation string
	err := d.db.QueryRow("SELECT organisation FROM oui WHERE assignment = ?", assignment).Scan(&organisation)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no organisation found for assignment: %s", assignment)
		}
		return "", fmt.Errorf("error querying database: %w", err)
	}
	return organisation, nil
}
