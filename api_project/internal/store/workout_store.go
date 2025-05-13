package store

import "database/sql"

type Workout struct {
	ID              int            `json:"id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	DurationMinutes int            `json:"duration_minutes"`
	CaloriesBurned  int            `json:"calories_burned"`
	Entries         []WorkoutEntry `json:"entries"`
}

type WorkoutEntry struct {
	ID              int    `json:"id"`
	ExerciseName    string `json:"exercise_name"`
	Sets            int    `json:"sets"`
	Reps            *int   `json:"reps"` // Is a pointer to int (*int) to allow for null values
	DurationSeconds *int   `json:"duration_seconds"`
	Weight          *int   `json:"weight"`
	Notes           string `json:"notes"`
	OrderIndex      int    `json:"order_index"`
}

type PostgresWorkoutStore struct {
	db *sql.DB
}

func CreatePostgresWorkoutStore(db *sql.DB) *PostgresWorkoutStore {
	return &PostgresWorkoutStore{
		db: db,
	}
}

// In all methods, we'll accept WorkoutStore instead of PostgresWorkoutStore
// to not tie the interface to a specific implementation (postgres).
type WorkoutStore interface {
	CreateWorkout(*Workout) error
	GetWorkoutByID(int) (*Workout, error)
}
