package routes

import (
	"github.com/go-chi/chi/v5"
)

func GetRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/api", GetAPIRoutes())
	r.Mount("/", GetWEBRoutes())

	return r
}
