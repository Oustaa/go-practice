package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/oustaa/go-practice/internal/app"
)

func GetRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/api", GetAPIRoutes(app))
	r.Mount("/", GetWEBRoutes(app))

	return r
}
