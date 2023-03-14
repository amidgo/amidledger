package routing

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandleReturn(c *gin.Context) {
	id := c.Request.FormValue("id")
	fmt.Println(c.Request.FormValue("result"))
	result := c.Request.FormValue("result") == "on"
	_, err := transact("acceptReturnProduct", id, GetUserData(c).Login, fmt.Sprint(result))
	if err != nil {
		RedirectToError(c, err)
		return
	}
	RedirectFromRequestToRolePage(c)
}
