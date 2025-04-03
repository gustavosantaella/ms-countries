package main

import (
	"net/http"
	"project/src/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var a string

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"pong": true,
		})
	})

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	router.Apirouter(r)
	// PORT, _ := helpers.EnvGetProperty("PORT")
	r.Run("0.0.0.0:10000")
}
