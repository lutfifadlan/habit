package handlers

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/lutfifadlan/habit/internal/models"
	"github.com/lutfifadlan/habit/internal/repository"
)

type UserHandler struct {
	repo *repository.Repository
}

func NewUserHandler(repo *repository.Repository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (u *UserHandler) Create(c fiber.Ctx) error {
	var request models.CreateUserRequest
	if err := c.Bind().Body(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	currentTime := time.Now()
	user := models.User{
		UserName:  request.UserName,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	if err := u.repo.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
