package parser

import (
	"fmt"
	"github.com/ezirl/massmedia/orm"
	"github.com/jinzhu/gorm"
	"github.com/mmcdole/gofeed"
	"html"
	"regexp"
	"time"
)

// очистка строки от тегов
func removeTags(s string) string {
	re := regexp.MustCompile("(<([^>]+)>)")
	s = re.ReplaceAllString(s, "")
	return s
}

type Site struct {
	gorm.Model
	Name     string
	Link     string
	Rss      string
	Category int
}

type Post struct {
	gorm.Model
	Title       string
	Link        string
	Description string
	Category    int
	SiteID      uint
}

func Start() {
	// подключение к бд и миграции
	//db, err := gorm.Open("mysql", "root@/parsnews?charset=utf8&parseTime=True&loc=Local")
	db := orm.GetDb()
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Site{})
	// основной бесконечный цикл
	for {
		// достаем юрлы сайтов из бд
		var sites []Site
		db.Find(&sites)
		work(sites, db)
		time.Sleep(time.Second * 5)
	}
}

func work(sites []Site, db *gorm.DB) {
	// парсинг rss
	fp := gofeed.NewParser()
	for _, site := range sites {
		//fmt.Println(site.ID, site.Name)
		//if site.ID != 13 {
		//	continue;
		//}
		feed, err := fp.ParseURL(site.Rss)
		if err != nil {
			_check(err)
			print(site.ID)
		} else {
			go writeRssToDb(site, feed, db)
		}
	}
	fmt.Println()
}

func writeRssToDb(site Site, feed *gofeed.Feed, db *gorm.DB) {
	// проход по rss и запись в базу
	var added int // количество добавленных постов
	if feed == nil {
		return
	}

	if len(feed.Items) < 6 {
		return
	}

	for i := 0; i <= 5; i++ {
		item := feed.Items[i]
		// проверка на уникальность юрл поста
		var post Post
		db.Where("link = ?", item.Link).Find(&post)
		if post.Link == "" {
			// запись в бд
			if item.Title != "" {
				_ = db.Create(&Post{
					Title:       html.UnescapeString(item.Title),
					Link:        item.Link,
					Description: html.UnescapeString(removeTags(item.Description)),
					SiteID:      site.ID,
					Category:    site.Category,
				}).Error
				added++ // инкремент добавленных постов
				time.Sleep(time.Second * 5)
			}
		}
	}
	fmt.Println(site.Name, " : ", added, " added")
}

func _check(e error) {
	if e != nil {
		fmt.Println(">>>>>>>>>>>>>error: ", e)
	}
}
