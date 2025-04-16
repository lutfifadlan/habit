package main

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v3"
	_ "github.com/mattn/go-sqlite3"

	"github.com/lutfifadlan/habit/internal/handlers"
	"github.com/lutfifadlan/habit/internal/migrations"
	"github.com/lutfifadlan/habit/internal/pkg/logger"
	"github.com/lutfifadlan/habit/internal/repository"
)

func main() {
	appLogger := logger.New()
	appLogger.Info("Starting habit tracker service")

	db, err := sql.Open("sqlite3", "./habits.db")
	if err != nil {
		appLogger.Error("Failed to initialize database: %v", err)
		os.Exit(1)
	}
	defer func() {
		if err := db.Close(); err != nil {
			appLogger.Error("Failed to close database: %v", err)
		}
	}()

	migrator := migrations.New(db, appLogger)
	if err := migrator.Run(); err != nil {
		appLogger.Error("Migrations failed: %v", err)
		os.Exit(1)
	}

	repo := repository.NewRepository(db, appLogger)
	habitHandler := handlers.NewHabitHandler(repo, appLogger)
	userHandler := handlers.NewUserHandler(repo)

	app := fiber.New(fiber.Config{
		AppName:       "Habit Tracker",
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		ErrorHandler: func(c fiber.Ctx, err error) error {
			appLogger.Error("HTTP error: %v", err)
			return fiber.DefaultErrorHandler(c, err)
		},
	})

	app.Use(func(c fiber.Ctx) error {
		appLogger.Info("%s %s", c.Method(), c.Path())
		return c.Next()
	})

	api := app.Group("/api/v1")
	{
		api.Post("/habits", habitHandler.Create)
		api.Get("/users/:user_id/habits", habitHandler.GetByUserId)
		api.Post("/users", userHandler.Create)
	}

	app.Get("/health", func(c fiber.Ctx) error {
		if err := db.Ping(); err != nil {
			appLogger.Error("Database ping failed: %v", err)
			return c.Status(503).SendString("Unhealthy")
		}
		return c.SendString("OK")
	})

	appLogger.Info("Server starting on :8080")
	if err := app.Listen(":8080"); err != nil {
		appLogger.Error("Server failed to start: %v", err)
		os.Exit(1)
	}
}
