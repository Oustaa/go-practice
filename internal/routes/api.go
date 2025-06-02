package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/oustaa/go-practice/internal/Handlers/api"
	"github.com/oustaa/go-practice/internal/app"
	"github.com/oustaa/go-practice/internal/store"
)

func GetAPIRoutes(app *app.Application) *chi.Mux {
	apiRoutes := chi.NewRouter()

	TaskStore := store.NewMySQLTasksService(app.DB)
	taksHandler := api.NewTaskshandler(TaskStore)

	apiRoutes.Get("/tasks", taksHandler.GetTasks)
	apiRoutes.Post("/tasks", taksHandler.PostTasks)

	return apiRoutes
}
