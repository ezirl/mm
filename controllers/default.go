package controllers

import (
	"github.com/ezirl/mm/models"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	postsPerPage := 50
	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, models.CountPost())
	c.Data["posts"] = models.GetListPostsByOffsetAndLimit(paginator.Offset(), postsPerPage)
	c.TplName = "index.tpl"
	
}

func (c *MainController) GetPostsBySite() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	postsPerPage := 50
	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, models.CountPostBySite(id))
	c.Data["posts"], c.Data["site"] = models.GetPostsBySite(id, paginator.Offset(), postsPerPage)
	c.TplName = "index.tpl"
}

func (c *MainController) GetPostsByCountry() {
	link := c.Ctx.Input.Param(":link")
	postsPerPage := 50
	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, models.CountPost())
	posts := models.GetPostsByCountry(link, paginator.Offset(), postsPerPage)
	if posts == nil {
		c.Abort("404")
	}
	c.Data["posts"] = posts
	c.Data["country"] = models.GetCountry(link)
	c.TplName = "index.tpl"
}

func (c *MainController) GetAbout() {
	c.TplName = "about.tpl"
}

func (c *MainController) GetAds() {
	c.Ctx.WriteString("google.com, pub-8200898278193793, DIRECT, f08c47fec0942fa0")
}

func (c *MainController) FormContact() {
	c.TplName = "contact.tpl"
}

func (c *MainController) Contact() {
	var cont models.Contact
	cont.Name = c.GetString("name")
	cont.Rss = c.GetString("rss")
	cont.Link = c.GetString("link")
	cont.Email = c.GetString("email")
	cont.Comm = c.GetString("comm")
	cont.Country = c.GetString("country")
	if err := models.AddContact(&cont); err == nil {
		c.Data["status"] = "отправлено"
		c.TplName = "status.tpl"
	} else {
		c.Data["status"] = err.Error()
		c.TplName = "status.tpl"
	}
}

func (c *MainController) DelContact() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DelContact(id); err == nil {
		c.Data["status"] = "удалено"
		c.Data["link"] = "/kosmo"
		c.TplName = "status.tpl"
	} else {
		c.Data["status"] = err.Error()
		c.TplName = "status.tpl"
	}
}

func (c *MainController) Link() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	models.IncTrust(uint(id))
	c.Data["post"] = models.GetPost(id)
	c.TplName = "link.tpl"
}
