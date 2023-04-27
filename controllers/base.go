package controllers

import (
	"github.com/ezirl/massmedia/models"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	c.Layout = "layout.tpl"
	c.Data["popular"] = models.GetPopularPostByOffsetAndLimit(0, 20)
	c.Data["countries"] = models.GetCountries()
	c.Data["politic"] = models.GetPostsByCategory(1, 30)
	c.Data["economic"] = models.GetPostsByCategory(2, 30)
	c.Data["elections"] = models.GetPostsByCategory(3, 30)
	c.Data["world"] = models.GetPostsByCategory(4, 30)
	c.Data["science"] = models.GetPostsByCategory(5, 30)
}
