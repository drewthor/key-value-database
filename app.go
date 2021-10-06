package main

import (
	"github.com/drewthor/key-value-database/controller"
	"github.com/drewthor/key-value-database/repository"
	"github.com/drewthor/key-value-database/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	DB := map[string]string{}

	r.Mount("/", controller.DatastoreController{DatastoreService: &service.DatastoreService{DatastoreDAO: &repository.DatastoreDAO{DB: DB}}}.Routes())

	http.ListenAndServe(":4000", r)
}
