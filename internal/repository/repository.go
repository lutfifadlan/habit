package repository

import (
	"database/sql"

	"github.com/lutfifadlan/habit/internal/pkg/logger"
)

type Repository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewRepository(db *sql.DB, logger *logger.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}
