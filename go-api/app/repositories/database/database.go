// Manage data getSettings and saveSettings and init table
package database

import (
	"webuiApi/app/repositories/domain"

	"github.com/olbrichattila/gofra/pkg/app/db"
)

type Database interface {
	Construct(db db.DBer)
	GetSettings() (domain.Setting, error)
	SaveSettings(domain.Setting) error
}

// New creates a new database manager
func New() Database {
	// Add if required different data engines as a factory, switch/case per os.Env
	return newSQLite()
}
