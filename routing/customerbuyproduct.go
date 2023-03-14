package routing

import (
	"github.com/gin-gonic/gin"
)

func CustomerBuyProduct(c *gin.Context) {
	u := GetUserData(c)
	shop := c.Query("shop")
	product := c.Request.FormValue("product")
	amount := c.Request.FormValue("amount")
	//BuyProductFromShop(ctx contractapi.TransactionContextInterface, login string, shopNumber string, productName string, amount float64)
	_, err := transact("buyProductFromShop", u.Login, shop, product, amount)
	if err != nil {
		RedirectToError(c, err)
		return
	}
	RedirectFromRequestToRolePage(c)
}
