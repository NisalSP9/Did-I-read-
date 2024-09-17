package routes

import (
	"github.com/NisalSP9/Did-I-read/models"
)

var ApplicationRoutes models.Routers

func init() {
	routes := []models.Routers{
		healthCheckRoutes,
		userRoutes,
		authRoutes,
	}

	for _, r := range routes {
		ApplicationRoutes = append(ApplicationRoutes, r...)
	}
}
