package routes

import (
	"github.com/jasongauvin/GAUVIN_JASON/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/GAUVIN_JASON/api/controllers"
)

// InitializeRoutes set up the routes for the server
func InitializeRoutes(r *gin.Engine) {
	// HTML routes for the GUI
	r.GET("/", controllers.SayHello)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// API routes
	api := r.Group("/api")
	// Check Authorization for api grouip
	api.Use(middleware.CheckAuthorization)

	{
		api.POST("/form", controllers.CreateForm)

	}
}
