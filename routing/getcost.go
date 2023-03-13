package routing

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCost(c *gin.Context) {
	product := c.Request.FormValue("name")
	amount := c.Request.FormValue("amount")
	fmt.Println(product, amount)
	cost, err := call("shopCalcCost", product, amount)
	if err != nil {
		RedirectToError(c, err)
		return
	}
	fmt.Println(string(cost))
	c.HTML(http.StatusOK, "cost.html", gin.H{"Cost": string(cost)})
}
