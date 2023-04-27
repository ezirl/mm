package main

import (
	_ "github.com/ezirl/mm/routers"
	"github.com/ezirl/mm/parser"
	"github.com/astaxie/beego"
)

func init() {
	go parser.Start()
	//go telegram.Start()
}

func main() {
	beego.Run()

}
