package routing

import "github.com/gin-gonic/gin"

func AcceptCost(c *gin.Context) {
	shop := c.Request.FormValue("shop")
	productName := c.Request.FormValue("product")
	amount := c.Request.FormValue("amount")

	_, err := transact("shopAcceptCost", shop, productName, amount)

	if err != nil {
		RedirectToError(c, err)
	}
}
