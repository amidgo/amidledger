package routing

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserData struct {
	Login string
	Role  string
}

func ShopPage(c *gin.Context) {
	shopsb, _ := call("getAllProducts")
	var productList []*Product
	json.Unmarshal(shopsb, productList)
	c.HTML(http.StatusOK, "shop.html", gin.H{"Products": productList, "UserData": UserData{Login: c.Query("login"), Role: c.Query("role")}})
}
