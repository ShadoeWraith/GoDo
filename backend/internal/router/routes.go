package router

import (
	"backend/internal/middleware"
	"backend/internal/tasks"
	"backend/internal/users"
	"backend/pkg/common"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", users.ListUsersHandler)
			r.Get("/{userID}", users.GetUserByIDHandler)
		})
		r.Route("/tasks", func(r chi.Router) {
			r.Use(middleware.Authenticate)

			r.Get("/", tasks.ListTasksHandler)
			r.Get("/{taskID}", tasks.GetSingleTaskHandler)
			r.Get("/deleted", tasks.ListDeletedTasksHandler)

			r.Post("/", tasks.CreateTaskHandler)

			r.Put("/{taskID}", tasks.UpdateTaskHandler)
			r.Put("/{taskID}/restore", tasks.RestoreDeletedTaskHandler)

			r.Delete("/{taskID}", tasks.DeleteTaskHandler)
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		common.JSONError(w, http.StatusNotFound, "The requested resource was not found.")
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		common.JSONError(w, http.StatusMethodNotAllowed, "Method not allowed.")
	})
}
