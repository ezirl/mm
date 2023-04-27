package controllers

import (
	"github.com/ezirl/massmedia/models"
	"strconv"
)

//  SiteController operations for Site
type SiteController struct {
	BaseController
}

// URLMapping ...
func (c *SiteController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Get", c.Get)
	c.Mapping("Delete", c.Delete)
}

func (c *SiteController) GetForm() {
	c.Data["Country"] = models.GetCountries()
	c.TplName = "addsite.tpl"
}

func (c *SiteController) Post() {
	var v models.Site
	v.Name = c.GetString("name")
	v.Link = c.GetString("link")
	v.Rss = c.GetString("rss")
	v.Ico = c.GetString("ico")
	id, _ := c.GetInt("country")
	v.Category, _ = c.GetInt("category")
	v.CountryID = uint(id)
	if err := models.AddSite(&v); err == nil {
		c.Data["status"] = "отправлено"
		c.Data["link"] = "/kosmo"
		c.TplName = "status.tpl"
	} else {
		c.Data["status"] = err.Error()
		c.TplName = "status.tpl"
	}
}

type Site struct {
	Name string `form:"name"`
	Link string `form:"link"`
	Rss  string `form:"rss"`
	Ico  string `form:"ico"`
}

func (c *SiteController) GetListSite() {
	c.Data["sites"] = models.ListSites()
	c.TplName = "sites.tpl"
}

func (c *SiteController) AddSite() {
	c.Data["Form"] = &Site{}
	c.Data["Country"] = models.GetCountries()
	c.TplName = "admin/site.tpl"
}

// @router /del/:id
func (c *SiteController) Del() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteSite(id); err != nil {
		c.Data["status"] = err.Error()
		c.Data["link"] = "/kosmo"
		c.TplName = "status.tpl"
	}
	c.Redirect("/kosmo", 200)
}

func (c *SiteController) Edit() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	c.TplName = "admin/siteupdate.tpl"
	c.Data["Site"] = models.GetSite(id)
	c.Data["Country"] = models.GetCountries()
}

func (c *SiteController) Put() {
	var v models.Site
	id, _ := c.GetInt("id")
	v.Name = c.GetString("name")
	v.Link = c.GetString("link")
	v.Rss = c.GetString("rss")
	v.Category, _ = c.GetInt("category")
	v.Ico = c.GetString("ico")
	idCountry, _ := c.GetInt("country")
	v.CountryID = uint(idCountry)
	if err := models.UpdateSite(id, &v); err == nil {
		c.Data["status"] = "отправлено"
		c.Data["link"] = "/kosmo"
		c.TplName = "status.tpl"
	} else {
		c.Data["status"] = err.Error()
		c.TplName = "status.tpl"
	}
}

// Delete ...
// @Title Delete
// @Description delete the Site
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SiteController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteSite(id); err == nil {
		c.Data["status"] = "отправлено"
		c.Data["link"] = "/kosmo"
		c.TplName = "status.tpl"
	} else {
		c.Data["status"] = err.Error()
		c.TplName = "status.tpl"
	}
}
