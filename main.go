package main

import (
	_ "web_profile/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

