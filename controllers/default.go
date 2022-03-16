package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "thecatapi.com"
	c.Data["Email"] = "astaxie@gmail.com"
	resp, err := http.Get("https://api.thecatapi.com/v1/images/search")
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	//var sb string = string(body)
	log.Println(reflect.TypeOf(body).Kind())
	//log.Println(string(body[1:len(body)]))
	jsonStr := string(body[1 : len(body)-1])
	var x map[string]interface{}
	log.Println(jsonStr)
	//json.Unmarshal([]byte(jsonStr), &x)
	log.Println(x)
	c.Data["api"] = "https://api.thecatapi.com/v1/images/search"
	c.Data["jsonStr"] = jsonStr
	c.TplName = "index.tpl"
}
