package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type User struct {
	Username string `json:"username,omitempty"`
	Grade    int    `json:"grade,omitempty"`
}
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("hello world")
}

func (c *MainController) Post() {
	user := &User{}
	err := json.Unmarshal(c.Ctx.Input.CopyBody(1<<32), user)
	if err != nil {
		fmt.Printf("Error %v", err)
		return
	}
	c.Data["json"] = user
	c.ServeJSON()
}
