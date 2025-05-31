package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetWEBRoutes() *chi.Mux {
	webRoutes := chi.NewRouter()

	webRoutes.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("WEB Root"))
	})

	return webRoutes
}
