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

	apiRoutes.Get("/tasks", taksHandler.GetTasksHandler)
	apiRoutes.Get("/tasks/{id}", taksHandler.GetTaskByIdHandler)
	apiRoutes.Post("/tasks", taksHandler.PostTasksHandler)
	apiRoutes.Delete("/tasks/{id}", taksHandler.DeleteTaskHandler)

	return apiRoutes
}
