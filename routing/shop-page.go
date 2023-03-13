package routing

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserData struct {
	Login    string
	Role     string
	Balance  float64
	Fio      string
	Products *Product
}

func ShopPage(c *gin.Context) {
	shopsb, _ := call("getAllProducts")
	var productList []*Product
	json.Unmarshal(shopsb, &productList)
	b, err := call("getShop", c.Query("login"))
	if err != nil {
		RedirectToError(c, err)
		return
	}
	shopUser := new(ShopUser)
	json.Unmarshal(b, shopUser)
	fmt.Println(shopUser.Shop)
	c.HTML(http.StatusOK, "shop.html", gin.H{"Products": productList, "UserData": shopUser.User, "Shop": shopUser.Shop})
}
