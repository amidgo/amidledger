package routing

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomerShopPageR(c *gin.Context) {
	u := GetUserData(c)
	name := c.Request.FormValue("shop")
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/customerShopPage?login=%s&role=%s&shop=%s", u.Login, u.Role, name))
}

func CustomerShopPage(c *gin.Context) {
	shop := c.Query("shop")
	b, err := call("getShop", shop)
	if err != nil {
		RedirectToError(c, err)
		return
	}
	shopS := new(ShopUser)
	json.Unmarshal(b, shopS)
	c.HTML(http.StatusOK, "customershop.html", gin.H{"UserData": GetUserData(c), "Shop": shopS})
}
