package main

import (
	_ "github.com/ezirl/massmedia/routers"
	"github.com/astaxie/beego"
)

func init() {
	go parser.Start()
	//go telegram.Start()
}

func main() {
	beego.Run()

}