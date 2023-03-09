package main

import (
	"net/http"

	"github.com/amidgo/amidledger/routing"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()
	c.LoadHTMLGlob("templates/*")
	c.Use(cors.Default())
	routing.Init()

	c.GET("/error", func(ctx *gin.Context) {
		err := ctx.Query("err")
		ctx.HTML(http.StatusOK, "error.html", gin.H{"Error": err})
	})

	c.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})

	c.Run("0.0.0.0:1110")
}
