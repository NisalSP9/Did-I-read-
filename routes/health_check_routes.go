package routes

import (
	"github.com/NisalSP9/Did-I-read/api"
	"github.com/NisalSP9/Did-I-read/models"
)

var healthCheckRoutes = models.Routers{

	models.Router{
		Name:    "Health check",
		Method:  "GET",
		Path:    "/",
		Handler: api.HealthCheck,
	},
}
