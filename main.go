package main

import (
	_ "catapi-bee/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

