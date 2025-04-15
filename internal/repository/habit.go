package repository

import (
	"encoding/json"

	"github.com/lutfifadlan/habit/internal/models"
)

func (r *Repository) CreateHabit(habit *models.Habit) error {
	result, err := r.db.Exec(`
		INSERT INTO habits (user_id, habit)
		VALUES (?, ?)
	`, habit.UserID, habit.Habit)

	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	habit.ID = int(id)
	r.logger.Info("Created habit: %v", habit)
	return nil
}

func (r *Repository) GetHabitsByUserId(userID int) ([]*models.Habit, error) {
	const query = `
		SELECT id, user_id, name, completion_dates, created_at, updated_at
		FROM habits
		WHERE user_id = ?
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		r.logger.Error("Failed to query habits: %v", err)
		return nil, err
	}
	defer rows.Close()

	var habits []*models.Habit
	for rows.Next() {
		var habit models.Habit
		var datesJSON string

		if err := rows.Scan(
			&habit.ID,
			&habit.UserID,
			&habit.Habit,
			&datesJSON,
			&habit.CreatedAt,
			&habit.UpdatedAt,
		); err != nil {
			r.logger.Error("Failed to scan habit row: %v", err)
			return nil, err
		}

		if datesJSON != "" {
			if err := json.Unmarshal([]byte(datesJSON), &habit.CompletionDates); err != nil {
				r.logger.Error("Failed to unmarshal completion dates: %v", err)
				return nil, err
			}
		}

		habits = append(habits, &habit)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Row iteration error: %v", err)
		return nil, err
	}

	return habits, nil
}
