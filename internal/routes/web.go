package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oustaa/go-practice/internal/app"
)

func GetWEBRoutes(app *app.Application) *chi.Mux {
	webRoutes := chi.NewRouter()

	webRoutes.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("WEB Root"))
	})

	return webRoutes
}
