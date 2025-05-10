package main

import (
	"apiProject/internal/app"
	"apiProject/internal/routes"
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "Go backend server port")
	flag.Parse()

	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Printf("Running app at http://localhost:%d\n", port)

	router := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()

	if err != nil {
		app.Logger.Fatal(err)
	}
}
