package routes

import (
	"github.com/NisalSP9/Did-I-read/api"
	"github.com/NisalSP9/Did-I-read/models"
)

var userRoutes = models.Routers{

	models.Router{
		Name:    "Create user",
		Method:  "POST",
		Path:    "/api/user",
		Handler: api.CreateUser,
	},
	models.Router{
		Name:    "Get user by id",
		Method:  "GET",
		Path:    "/api/user/{userid}",
		Handler: api.GetUserById,
	},
}
