package routers

import (
	"github.com/ezirl/mm/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/ads.txt", &controllers.MainController{}, "get:GetAds")

	beego.Router("/", &controllers.MainController{})
	beego.Router("/?:link", &controllers.MainController{}, "get:GetPostsByCountry")
	beego.Router("/site/?:id", &controllers.MainController{}, "get:GetPostsBySite")
	beego.Router("/sites", &controllers.SiteController{}, "get:GetListSite")
	beego.Router("/addsite", &controllers.SiteController{}, "get:GetForm")
	beego.Router("/about", &controllers.MainController{}, "get:GetAbout")
	beego.Router("/contact", &controllers.MainController{}, "get:FormContact")
	beego.Router("/contact", &controllers.MainController{}, "post:Contact")
	beego.Router("/link/?:id", &controllers.MainController{}, "get:Link")

	//admin panel
	beego.Router("/kosmo", &controllers.AdminController{})
	beego.Router("/kosmo/addnews", &controllers.AdminController{}, "post:AddNews")
	beego.Router("/kosmo/country", &controllers.AdminController{}, "get:AddCountry")
	beego.Router("/kosmo/country", &controllers.AdminController{}, "post:AddCountryPost")
	beego.Router("/kosmo/contacts/spam", &controllers.AdminController{}, "get:RemoveSpam")
	beego.Router("/kosmo/site/contact/?:id", &controllers.MainController{}, "get:DelContact")
	beego.Router("/kosmo/site", &controllers.SiteController{}, "get:AddSite")
	beego.Router("/kosmo/site", &controllers.SiteController{}, "post:Post")
	beego.Router("/kosmo/site/edit", &controllers.SiteController{}, "post:Put")
	beego.Router("/kosmo/site/edit/?:id", &controllers.SiteController{}, "get:Edit")
	beego.Router("/kosmo/site/del/?:id", &controllers.SiteController{}, "get:Del")
	//posts elections
	beego.Router("/kosmo/election", &controllers.AdminController{}, "get:Election")
	beego.Router("/kosmo/election/add", &controllers.AdminController{}, "get:AddPostGet")
	beego.Router("/kosmo/election", &controllers.AdminController{}, "post:AddPost")
	beego.Router("/kosmo/election/edit", &controllers.AdminController{}, "post:UpdatePost")
	beego.Router("/kosmo/election/edit/?:id", &controllers.AdminController{}, "get:EditPost")
	beego.Router("/kosmo/election/del/?:id", &controllers.AdminController{}, "get:DelPost")

}
