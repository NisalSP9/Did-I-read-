package routes

import (
	"github.com/NisalSP9/Did-I-read/api"
	"github.com/NisalSP9/Did-I-read/models"
)

var authRoutes = models.Routers{
	models.Router{
		Name:    "Get user auth",
		Method:  "POST",
		Path:    "/api/auth/token",
		Handler: api.UserAuth,
	},
	models.Router{
		Name:    "Get refresh token",
		Method:  "Get",
		Path:    "/api/auth/token/refresh",
		Handler: api.RefreshToke,
	},
}
