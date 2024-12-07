package main

import (
	"github.com/Shachindra/altlife/server/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infof("Starting AltLife Version: v1.0")
	ginApp := gin.Default()
	// cors middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	ginApp.Use(cors.New(config))

	ginApp.Use(static.Serve("/", static.LocalFile("../ui", false)))

	ginApp.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": 404, "result": "Invalid Endpoint Request"})
	})
	api.ApplyRoutes(ginApp)
	// ginApp.Run(":" + os.Getenv("HTTP_PORT"))
	ginApp.Run(":" + "9080")
}
