package database

import (
	"errors"
	"fmt"
	"webuiApi/app/repositories/domain"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olbrichattila/gofra/pkg/app/db"
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

func newSQLite() Database {
	return &sqlite{}
}

type sqlite struct {
	db db.DBer
}

// Construct implements Database.
func (s *sqlite) Construct(db db.DBer) {
	s.db = db
}

// GetSettings implements Database.
func (s *sqlite) GetSettings() (domain.Setting, error) {

	row, err := s.db.QueryOne("SELECT region, endpoint, key, secret FROM settings LIMIT 1")
	if err != nil && err.Error() == "row cannot be found" {
		return domain.Setting{}, nil
	}

	if err != nil {
		return domain.Setting{}, fmt.Errorf("%w: %v", errDatabase, err)
	}

	setting := domain.Setting{
		Region:   row["region"].(string),
		Endpoint: row["endpoint"].(string),
		Credentials: domain.Credentials{
			Key:    row["key"].(string),
			Secret: row["secret"].(string),
		},
	}

	return setting, nil
}

// SaveSettings implements Database.
func (s *sqlite) SaveSettings(setting domain.Setting) error {
	row, err := s.db.QueryOne("SELECT count(*) as cnt FROM settings")

	if err != nil {
		return fmt.Errorf("%w: %v", errDatabase, err)
	}

	cnt, ok := row["cnt"].(int64)
	if !ok {
		return fmt.Errorf("%w: %v", errInitDatabase, err)
	}

	if cnt == 0 {
		_, err = s.db.Execute(`
			INSERT INTO settings (region, endpoint, key, secret, created_at, updated_at)
			VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`,
			setting.Region, setting.Endpoint, setting.Credentials.Key, setting.Credentials.Secret)

		if err != nil {
			return fmt.Errorf("%w: %v", errDatabase, err)
		}

		return nil
	}

	_, err = s.db.Execute(`
		UPDATE settings 
		SET region = ?, endpoint = ?, key = ?, secret = ?, updated_at = CURRENT_TIMESTAMP 
		WHERE id = 1`,
		setting.Region, setting.Endpoint, setting.Credentials.Key, setting.Credentials.Secret)

	if err != nil {
		return fmt.Errorf("%w: %v", errDatabase, err)
	}

	return nil
}
