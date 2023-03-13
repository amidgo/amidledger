package routing

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BuyProduct(c *gin.Context) {
	id := c.Query("id")
	accept := c.Query("accept")
	fmt.Println(accept)
	_, err := transact("shopBuyProduct", id, accept)
	if err != nil {
		RedirectToError(c, err)
		return
	}
	RedirectFromRequestToRolePage(c)
}
