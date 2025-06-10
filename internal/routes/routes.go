package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/kaipov24/scaffold-api/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/Health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutByID)

	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)

	return r
}
