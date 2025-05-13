package api

import (
	"apiProject/internal/store"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
}

func CreateWorkoutHandler(workoutStore store.WorkoutStore) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore,
	}
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	// json handles http payload
	err := json.NewDecoder(r.Body).Decode(&workout) // content to be parsed is the request body

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create workout", http.StatusInternalServerError)
		return
	}

	createdWorkout, err := wh.workoutStore.CreateWorkout(&workout)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create workout", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdWorkout)
}

func (wh *WorkoutHandler) HandleGetWorkoutById(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL
	paramsWorkoutId := chi.URLParam(r, "id")

	if paramsWorkoutId == "" {
		http.NotFound(w, r)
	}

	// converting the ID to number just to make easier local  development
	workoutID, err := strconv.ParseInt(paramsWorkoutId, 10, 64)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "This is the workout id %d\n", workoutID)
}
