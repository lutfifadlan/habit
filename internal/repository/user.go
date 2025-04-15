package repository

import "github.com/lutfifadlan/habit/internal/models"

func (r *Repository) CreateUser(user *models.User) error {
	result, err := r.db.Exec(`
		INSERT INTO users (username)
		VALUES (?)
	`, user.UserName)

	if err != nil {
		r.logger.Error("Failed to create user: %v", err)
		return err
	}

	id, _ := result.LastInsertId()
	user.ID = int(id)
	r.logger.Info("Created user: %+v", *user)
	return nil
}
