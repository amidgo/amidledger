package routing

import "github.com/gin-gonic/gin"

func CreateProduct(c *gin.Context) {
	name := c.Request.FormValue("name")
	producer := c.Request.FormValue("producer")
	date := c.Request.FormValue("date")
	safetime := c.Request.FormValue("safetime")
	temp1 := c.Request.FormValue("temp1")
	temp2 := c.Request.FormValue("temp2")
	price := c.Request.FormValue("price")
	_, err := transact("createProduct", name, producer, date, safetime, temp1, temp2, price)
	if err != nil {
		RedirectToError(c, err)
		return
	}
	RedirectFromRequestToRolePage(c)
}
