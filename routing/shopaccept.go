package routing

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShopAcceptCost(c *gin.Context) {
	userData := GetUserData(c)
	productName := c.Request.FormValue("name")
	am := c.Request.FormValue("amount")
	js, err := transact("shopAcceptCost", userData.Login, productName, am)
	if err != nil {
		RedirectToError(c, err)
		return
	}
	tx := new(ShopTransaction)
	json.Unmarshal(js, tx)
	for i := 0; i < 6; i++ {
		js, err := transact("delivery", tx.Id)
		if err != nil {
			RedirectToError(c, err)
			return
		}
		json.Unmarshal(js, tx)
	}

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/shopAcceptCostPage?login=%s&role=%s&id=%s", userData.Login, userData.Role, tx.Id))
}
