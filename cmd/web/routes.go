package main

import (
	"net/http"

	"github.com/MyAusweis/bookings/cmd/pkg/config"
	"github.com/MyAusweis/bookings/cmd/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

/* using pat
func routes(app *config.AppConfig) http.Handler {

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
} */

// using chi
func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//fileServer := http.FileServer(http.Dir("./static/"))
	//mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	//return mux

	// Create a file server for your files
	fs := http.FileServer(http.Dir("./static/")) // folder static/images/ isi.jpg, pk ini imagesnya di bypass
	// We don't want that /assets/ prefix in our file paths, so let's strip it out.
	prefixHandler := http.StripPrefix("/static/", fs) //no 2
	mux.Handle("/static/*", prefixHandler)            //there is folder/static/* , should call prefixHandler  -no 1

	return mux

}
