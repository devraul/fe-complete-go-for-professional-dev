package main

import (
	"apiProject/internal/app"
	"fmt"
	"net/http"
	"time"
)

func main() {
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("Running app at http://localhost:8080")

	http.HandleFunc("/health", HealthCheck)

	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatal(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available")
}
