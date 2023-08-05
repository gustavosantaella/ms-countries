package main

import (
	"net/http"
	"project/src/databases/supabase"
	"project/src/router"

	"github.com/gin-gonic/gin"
)

var a string

func main() {
	supabase.Connect()

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"pong": true,
		})
	})

	router.Apirouter(r)
	r.Run("0.0.0.0:4000")
}
