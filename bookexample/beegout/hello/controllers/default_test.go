package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"
	"github.com/dghubble/sling"
)

func TestMainController_Get(t *testing.T) {
	beego.Router("/sample", &MainController{}, "get:Get")
	req, err := sling.New().Get("/sample").Request()
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(body))
}

func TestMainController_Post(t *testing.T) {
	beego.Router("/user", &MainController{}, "post:Post")
	body := &User{
		Username: "john",
		Grade:    10,
	}
	req, err := sling.New().Post("/user").BodyJSON(body).Request()
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)
	respbody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(respbody))
}
