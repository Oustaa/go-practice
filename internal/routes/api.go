package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetAPIRoutes() *chi.Mux {
	apiRoutes := chi.NewRouter()

	apiRoutes.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Root"))
	})

	return apiRoutes
}
