package routing

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FabricUser struct {
	Login    string     `json:"login"`
	Password string     `json:"password"`
	Role     string     `json:"role"`
	Balance  float64    `json:"balance"`
	Fio      string     `json:"fio"`
	Products []*Product `json:"products"`
}

type Product struct {
	Name     string  `json:"name"`
	Producer string  `json:"producer"`
	Date     string  `json:"date"`
	Safetime int     `json:"safetime"`
	Temp1    int     `json:"temp1"`
	Temp2    int     `json:"temp2"`
	Amount   float64 `json:"amount"`
	Price    float64 `json:"price"`
	Shop     string  `json:"shop"`
}

type ProductReturnTx struct {
	Id          string `json:"id"`
	ProductName string `json:"productname"`
	UserLogin   string `json:"userlogin"`
	ShopNumber  string `json:"shopnumber"`
}

type Shop struct {
	Number   string     `json:"number"`
	City     string     `json:"city"`
	Products []*Product `json:"products"`
}

type ShopTransaction struct {
	Id             string   `json:"id"`
	ShopNumber     string   `json:"shopnumber"`
	StartCost      float64  `json:"startcost"`
	CurrentCost    float64  `json:"currentcost"`
	Temp           []int    `json:"temp"`
	Product        *Product `json:"product"`
	TempFailAmount uint     `json:"tempfailamount"`
}

type ShopUser struct {
	User *FabricUser `json:"user"`
	Shop *Shop       `json:"shop"`
}

func RedirectToError(ctx *gin.Context, err error) {
	ctx.Redirect(http.StatusMovedPermanently, "/error?err="+err.Error())
}

func RedirectToRolePage(ctx *gin.Context, login string, role string) {
	ctx.Redirect(http.StatusMovedPermanently, "/user-page"+"?login="+login+"&role="+role)
}

func RedirectFromRequestToRolePage(ctx *gin.Context) {
	role := ctx.Query("role")
	login := ctx.Query("login")
	time.Sleep(time.Second)
	RedirectToRolePage(ctx, login, role)
}

func Login(c *gin.Context) {
	login := c.Request.FormValue("login")
	password := c.Request.FormValue("password")

	usr, err := call("login", login, password)
	if err != nil {
		RedirectToError(c, err)
		return
	}
	fuser := new(FabricUser)
	json.Unmarshal(usr, &fuser)
	RedirectToRolePage(c, fuser.Login, fuser.Role)
}
