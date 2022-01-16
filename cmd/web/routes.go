package main

import (
	"net/http"

	"github.com/ankitkr1924/BankManagementApp/pkg/config"
	"github.com/ankitkr1924/BankManagementApp/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/", handlers.Repo.About)

	return mux
}
