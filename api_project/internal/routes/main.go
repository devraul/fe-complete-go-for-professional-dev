package routes

import (
	"apiProject/internal/app"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/health", healthCheck)

	return router
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available")
}
