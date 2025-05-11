package app

import (
	"apiProject/internal/api"
	"apiProject/internal/store"
	"database/sql"
	"log"
	"os"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	db, err := store.Open()

	if err != nil {
		return nil, err
	}

	// Stores

	// Handlers
	workoutHandler := api.CreateWorkoutHandler()

	application := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             db,
	}

	return application, nil
}
