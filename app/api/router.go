package api

import (
	"github.com/dogukanayd/go-rest/app/api/controllers"
	"github.com/gin-gonic/gin"
)

// Start new api server instance
func Start() *gin.Engine {
	g := gin.New()

	// Global middleware
	g.Use(gin.Recovery())
	g.GET("health", controllers.Health)

	v1 := g.Group("v1")
	{
		usr := v1.Group("user")
		{
			usr.GET("", controllers.User)
			usr.POST("sign-up")
			usr.POST("login")
			usr.PATCH("update")
			usr.DELETE("delete")
		}
	}

	return g
}
