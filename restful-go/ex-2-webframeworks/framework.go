// This is exercise uses the following web frameworks:
// 1. chi - https://github.com/go-chi/chi

// to run this exercise:
// go test -run TestUserEndpoints
package main

import (
	"github.com/go-chi/chi/middleware"
	"net/http"

	"github.com/go-chi/chi"
)

func getUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	// check if username is empty -> return 400
	if username == "" {
		http.Error(w, http.StatusText(400)+", username parameter cannot be empty", 400)
		return
	}

	// return username and 200
	w.Write([]byte(username))
}

func updateUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	// check if username is empty -> return 400
	if username == "" {
		http.Error(w, http.StatusText(400)+", username parameter cannot be empty", 400)
		return
	}

	// return response and 200
	w.Write([]byte("user registered"))
}

func deleteUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	// check if username is empty -> return 400
	if username == "" {
		http.Error(w, http.StatusText(400)+", username parameter cannot be empty", 400)
		return
	}

	// return response and 200
	w.Write([]byte("user deleted")) // TODO: return deleted username
}
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Route("/user", func(r chi.Router) {
		r.Put("/update", updateUsername)
		r.Get("/{username}", getUsername)
		r.Post("/delete", deleteUsername)
	})
	http.ListenAndServe(":3000", r)
}
