package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/lutfifadlan/habit/internal/models"
	"github.com/lutfifadlan/habit/internal/pkg/logger"
	"github.com/lutfifadlan/habit/internal/repository"
)

type HabitHandler struct {
	repo   *repository.Repository
	logger *logger.Logger
}

func NewHabitHandler(repo *repository.Repository, logger *logger.Logger) *HabitHandler {
	return &HabitHandler{repo: repo, logger: logger}
}

func (h *HabitHandler) Create(c fiber.Ctx) error {
	var request models.CreateHabitRequest
	if err := c.Bind().Body(&request); err != nil {
		h.logger.Error("Failed to bind request body: error=%v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	currentTime := time.Now()
	habit := models.Habit{
		UserID:    request.UserID,
		Habit:     request.Habit,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	if err := h.repo.CreateHabit(&habit); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create habit",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(habit)
}

func (h *HabitHandler) GetByUserId(c fiber.Ctx) error {
	userIDStr := c.Params("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID <= 0 {
		h.logger.Error("Invalid userID parameter: input=%s, error=%v", userIDStr, err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID must be a positive integer",
		})
	}

	habits, err := h.repo.GetHabitsByUserId(userID)
	if err != nil {
		h.logger.Error("Failed to get habits: user_id=%s, error=%v", userIDStr, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve habits",
		})
	}

	if len(habits) == 0 {
		return c.JSON([]interface{}{})
	}

	return c.JSON(habits)
}
