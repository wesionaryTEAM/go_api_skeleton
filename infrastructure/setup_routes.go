package infrastructure

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "TeamPlace API Server is Runninggg..."})
	})

	// Run server
	port := os.Getenv("SERVER_PORT")
	httpRouter.Run(port)
}
