package routing

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProviderPage(c *gin.Context) {
	c.HTML(http.StatusOK, "provider.html", gin.H{"UserData": UserData{Login: c.Query("login"), Role: c.Query("role")}})
}
