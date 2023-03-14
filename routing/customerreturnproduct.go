package routing

import "github.com/gin-gonic/gin"

// ReturnProductToShop(ctx contractapi.TransactionContextInterface, login string, shopNumber string, productName string)
func ReturnProductToShop(c *gin.Context) {
	u := GetUserData(c)
	shop := c.Query("shop")
	product := c.Request.FormValue("product")
	_, err := transact("returnProductToShop", u.Login, shop, product)
	if err != nil {
		RedirectToError(c, err)
		return
	}
	RedirectFromRequestToRolePage(c)
}
