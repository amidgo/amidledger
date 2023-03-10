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

	c.POST("/login", routing.Login)

	c.GET("/shop-page", routing.ShopPage)

	c.GET("/provider-page", routing.ProviderPage)

	c.GET("/customer-page", routing.CustomerPage)

	c.GET("/getCost", routing.GetCost)

	c.GET("/shopAcceptCostPage", routing.ShopAcceptPage)

	c.GET("/customerShopPage", routing.CustomerShopPage)

	c.POST("/customerShopPageR", routing.CustomerShopPageR)

	c.POST("/shopAcceptCost", routing.ShopAcceptCost)

	c.POST("/createProduct", routing.CreateProduct)

	c.POST("/buyProduct", routing.BuyProduct)

	c.POST("/buyProductCustomer", routing.CustomerBuyProduct)

	c.POST("/returnProduct", routing.ReturnProductToShop)

	c.POST("/handleReturn", routing.HandleReturn)

	c.Run("0.0.0.0:1110")
}
