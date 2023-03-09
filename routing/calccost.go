package routing

import "github.com/gin-gonic/gin"

func GetCost(c *gin.Context) {
	product := c.Request.FormValue("product")
	amount := c.Request.FormValue("amount")
}
