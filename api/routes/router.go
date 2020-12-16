package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/gogoleplate/api/controllers"
)

// InitializeRoutes set up the routes for the server
func InitializeRoutes(r *gin.Engine) {
	// HTML routes for the GUI
	r.GET("/", controllers.SayHello)

	// API routes
	api := r.Group("/api")
	{
		api.GET("/register", controllers.Register)
		api.GET("/login", controllers.Login)
	}
}
