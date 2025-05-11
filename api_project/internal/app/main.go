package app

import (
	"apiProject/internal/api"
	"log"
	"os"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Stores

	// Handlers
	workoutHandler := api.CreateWorkoutHandler()

	application := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}

	return application, nil
}
