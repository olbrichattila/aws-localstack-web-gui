package database

import (
	"database/sql"
	"errors"
	"fmt"
	"webuiApi/app/repositories/domain"

	_ "github.com/mattn/go-sqlite3"
)

const create_table = `
CREATE TABLE IF NOT EXISTS settings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    region VARCHAR(64) NOT NULL,
    endpoint TEXT NOT NULL,
    "key" VARCHAR(255) NOT NULL,
    secret VARCHAR(255) NOT NULL
)`

var (
	errDatabase     = errors.New("database error")
	errInitDatabase = errors.New("database initialization error")
)

func newSQLite() (Database, error) {
	dataManager := &sqlite{}

	if err := dataManager.createTable(); err != nil {
		return nil, err
	}

	return dataManager, nil
}

type sqlite struct {
}

// GetSettings implements Database.
func (s *sqlite) GetSettings() (domain.Setting, error) {
	setting := domain.Setting{
		Credentials: domain.Credentials{},
	}

	db, err := s.initDb()
	if err != nil {
		return setting, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT region, endpoint, key, secret FROM settings LIMIT 1")
	err = row.Scan(&setting.Region, &setting.Endpoint, &setting.Credentials.Key, &setting.Credentials.Secret)
	if err == sql.ErrNoRows {
		return domain.Setting{}, nil
	}

	if err != nil {
		return domain.Setting{}, fmt.Errorf("%w: %v", errDatabase, err)
	}

	return setting, nil
}

// SaveSettings implements Database.
func (s *sqlite) SaveSettings(setting domain.Setting) error {
	db, err := s.initDb()
	if err != nil {
		return err
	}
	defer db.Close()
	// Try to update first
	result, err := db.Exec(`
		UPDATE settings 
		SET region = ?, endpoint = ?, key = ?, secret = ?, updated_at = CURRENT_TIMESTAMP 
		WHERE id = 1`,
		setting.Region, setting.Endpoint, setting.Credentials.Key, setting.Credentials.Secret)

	if err != nil {
		return fmt.Errorf("%w: %v", errDatabase, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w: %v", errDatabase, err)
	}

	// If no rows were updated, do an insert
	if rows == 0 {
		_, err = db.Exec(`
			INSERT INTO settings (region, endpoint, key, secret, created_at, updated_at)
			VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`,
			setting.Region, setting.Endpoint, setting.Credentials.Key, setting.Credentials.Secret)

		if err != nil {
			return fmt.Errorf("%w: %v", errDatabase, err)
		}
	}

	return nil
}

func (s *sqlite) initDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "settings.db")
	if err != nil {
		return nil, fmt.Errorf("%w, %w", errInitDatabase, err)
	}
	return db, nil
}

func (s *sqlite) createTable() error {
	db, err := s.initDb()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(create_table)
	if err != nil {
		return fmt.Errorf("crate table %w", err)
	}
	return nil
}
