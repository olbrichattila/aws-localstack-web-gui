// Manage data getSettings and saveSettings and init table
package database

import "api/internal/domain"

type Database interface {
	GetSettings() (domain.Setting, error)
	SaveSettings(domain.Setting) error
}

// New creates a new database manager
func New() (Database, error) {
	// Add if required different data engines as a factory, switch/case per os.Env
	return newSQLite()
}
