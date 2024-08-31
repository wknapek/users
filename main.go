package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"main/handlers"
	"main/security"
	"net/http"
	"os"
	"time"
)

func main() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Logger()
	var port string
	if len(os.Args) < 2 {
		port = ":3001"
	} else {
		port = ":" + port
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/login", security.Login)
	router.Get("/create", handlers.CreateUserHandler)
	router.Get("/get/{id}", handlers.GetUserHandler)
	router.Get("/getall", handlers.GetAllUsersHandler)
	log.Info().Msg("server  starting listening on port " + port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}

}
