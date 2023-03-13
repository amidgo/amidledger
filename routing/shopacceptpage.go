package routing

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShopAcceptPage(c *gin.Context) {
	uData := GetUserData(c)
	id := c.Query("id")
	tx := new(ShopTransaction)
	js, err := call("getShopTransaction", id)
	if err != nil {
		RedirectToError(c, err)
		return
	}
	json.Unmarshal(js, tx)
	c.HTML(http.StatusOK, "accept.html", gin.H{
		"Id":             id,
		"UserData":       uData,
		"TempFailAmount": tx.TempFailAmount,
		"CurrentPrice":   tx.CurrentCost,
	})
}
