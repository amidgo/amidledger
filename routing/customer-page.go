package routing

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomerPage(c *gin.Context) {
	uData := GetUserData(c)
	b, err := call("getAllShops")
	if err != nil {
		RedirectToError(c, err)
		return
	}
	shops := make([]*Shop, 0)
	json.Unmarshal(b, &shops)
	c.HTML(http.StatusOK, "customer.html", gin.H{"UserData": uData, "Shops": shops})
}
