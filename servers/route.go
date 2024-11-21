package servers

import (
	"net/http"
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupYourRoute(g *gin.Engine, serverController *controllers.ServerController) {
	u := g.Group("/")
	u.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
	u.GET("/server", serverController.Get)
}
