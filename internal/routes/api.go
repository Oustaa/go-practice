package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oustaa/go-practice/internal/Handlers/api"
	"github.com/oustaa/go-practice/internal/app"
	"github.com/oustaa/go-practice/internal/store"
)

func GetAPIRoutes(app *app.Application) *chi.Mux {
	apiRoutes := chi.NewRouter()

	TaskStore := store.NewMySQLTasksService(app.DB)
	taksHandler := api.NewTaskshandler(TaskStore)

	apiRoutes.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Header().Set("Content-Type", "Application/json")
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{"Page Not Found"})
	})
	apiRoutes.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{"Method Not Allowed"})
	})

	apiRoutes.Get("/tasks", taksHandler.GetTasks)
	apiRoutes.Put("/tasks/{id}", taksHandler.PutTask)
	apiRoutes.Post("/tasks", taksHandler.PostTask)

	return apiRoutes
}
