package servers

import (
	"net/http"
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func SetupYourRoute(g *gin.Engine, exampleController *controllers.ExampleController) {
	u := g.Group("/")
	u.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
	u.GET("/example", exampleController.Get)
}
