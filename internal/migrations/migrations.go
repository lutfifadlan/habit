package migrations

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/lutfifadlan/habit/internal/pkg/logger"
)

//go:embed sql/*.sql
var sqlFiles embed.FS

type Migration struct {
	db     *sql.DB
	logger *logger.Logger
}

func New(db *sql.DB, logger *logger.Logger) *Migration {
	return &Migration{db: db, logger: logger}
}

func (m *Migration) Run() error {
	m.logger.Info("Running database migrations")

	files, err := sqlFiles.ReadDir("sql")
	if err != nil {
		m.logger.Error("Failed to read migration directory: %v", err)
		return fmt.Errorf("Failed to read migration files: %w", err)
	}

	for _, file := range files {
		m.logger.Info("Executing migration file", "file", file.Name())

		content, err := sqlFiles.ReadFile("sql/" + file.Name())
		if err != nil {
			m.logger.Error("Failed to read migration file: %v", err)
			return fmt.Errorf("Failed to read %s:%w", file.Name(), err)
		}

		if _, err := m.db.Exec(string(content)); err != nil {
			m.logger.Error("Migration failed", "file", file.Name(), "error", err)
			return fmt.Errorf("failed to execute %s: %w", file.Name(), err)
		}
	}

	m.logger.Info("Migrations completed successfully")
	return nil
}
