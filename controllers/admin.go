package controllers

import (
	"fmt"
	"strconv"

	"github.com/ezirl/mm/models"
	"github.com/astaxie/beego"
)

type AdminController struct {
	BaseController
}

type Kosmo struct {
	Login    string `form:"login"`
	Password string `form:"password,password"`
}

func (a *AdminController) Get() {
	if a.GetSession("kosmo") == beego.AppConfig.String("ps") {
		a.Data["Sites"] = models.ListSites()
		a.Data["Contacts"] = models.ListContacts()
		a.TplName = "admin/index.tpl"
	} else {
		a.Data["Form"] = &Kosmo{}
		a.TplName = "admin/login.tpl"
	}
}

func (a *AdminController) AddNews() {
	fmt.Println("news add work")
	news := models.Post{}
	news.Title = a.GetString("title")
	news.Link = a.GetString("url")
	news.Description = a.GetString("description")
	id, err := a.GetInt("site", 0)
	if err == nil && id != 0 {
		news.SiteID = uint(id)
		models.AddPost(&news)
	}
	a.Redirect("/", 303)
}

func (a *AdminController) Post() {
	login := a.GetString("login")
	password := a.GetString("password")
	if login == beego.AppConfig.String("login") && password == beego.AppConfig.String("ps") {
		_ = a.StartSession().Set("kosmo", password)
		a.Get()
	} else {
		a.Redirect("/kosmo", 401)
	}
}

func (c *AdminController) AddCountry() {
	c.TplName = "admin/country.tpl"
}

func (c *AdminController) AddCountryPost() {
	var v models.Country
	v.Name = c.GetString("name")
	v.Link = c.GetString("link")
	if err := models.AddCountry(&v); err == nil {
		c.Data["status"] = "отправлено"
		c.TplName = "status.tpl"
		c.Data["link"] = "/kosmo"
	} else {
		c.Data["status"] = err.Error()
		c.TplName = "status.tpl"
	}
}

func (c *AdminController) Election() {

	if c.GetSession("kosmo") == beego.AppConfig.String("ps") {
		c.Data["Elections"] = models.GetPostsByCategory(3, 20)
		c.TplName = "admin/election/index.tpl"
	} else {
		c.Redirect("/kosmo", 401)
	}
}

func (c *AdminController) AddPostGet() {
	c.TplName = "admin/election/add.tpl"
}

func (c *AdminController) AddPost() {
	var v models.Post
	v.Title = c.GetString("title")
	v.Link = c.GetString("link")
	v.SiteName = c.GetString("sitename")
	v.Description = c.GetString("description")
	if err := models.AddPost(&v); err == nil {
		c.Data["status"] = "отправлено"
		c.Data["link"] = "/kosmo/election"
		c.TplName = "status.tpl"
	} else {
		c.Data["status"] = err.Error()
		c.TplName = "status.tpl"
	}
}

func (c *AdminController) UpdatePost() {
	var v models.Post
	id, _ := c.GetInt("id")
	v.Title = c.GetString("title")
	v.Link = c.GetString("link")
	v.SiteName = c.GetString("sitename")
	v.Description = c.GetString("description")
	if err := models.UpdatePost(id, v); err == nil {
		c.Data["status"] = "отправлено"
		c.Data["link"] = "/kosmo/election"
		c.TplName = "status.tpl"
	} else {
		c.Data["status"] = err.Error()
		c.TplName = "status.tpl"
	}
}

func (c *AdminController) DelPost() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DelPost(id); err != nil {
		c.Data["status"] = err.Error()
		c.TplName = "status.tpl"
	}
	c.Redirect("/kosmo/election", 200)
}

func (c *AdminController) EditPost() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	c.TplName = "admin/election/update.tpl"
	c.Data["election"] = models.GetPost(id)
}

func (c *AdminController) RemoveSpam() {

	if err := models.DelSpam(); err == nil {
		c.Data["status"] = "удалено"
		c.Data["link"] = "/kosmo"
		c.TplName = "status.tpl"
	} else {
		c.Data["status"] = err.Error()
		c.TplName = "status.tpl"
	}
}
