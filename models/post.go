package models

import (
	"fmt"
	"time"

	"github.com/ezirl/mm/orm"
	"github.com/jinzhu/gorm"
)

func init() {
	orm.GetDb().AutoMigrate(&Post{})
}

type Post struct {
	gorm.Model
	Title       string `gorm:"not null;unique"`
	Link        string `gorm:"not null;unique"`
	Description string `gorm:"type:text"`
	Site        Site
	SiteID      uint
	Category    int
	Trust       int
	SiteName    string
}

func AddPost(post *Post) error {
	db := orm.GetDb()
	if err := db.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePost(id int, post Post) error {
	db := orm.GetDb()
	var pos Post
	db.First(&pos)
	fmt.Println(pos, id)
	pos.Link = post.Link
	pos.Title = post.Title
	pos.Description = post.Description
	pos.SiteName = post.SiteName
	pos.Category = 3
	if err := db.Model(&pos).Update(&pos).Error; err != nil {
		return err
	}
	return nil
}

func DelPost(id int64) error {
	db := orm.GetDb()
	err := db.Where("id = ?", id).Delete(&Post{}).Error
	if err != nil {
		return err
	}
	return nil
}

func IncTrust(id uint) {
	db := orm.GetDb()
	var post Post
	db.First(&post, id)
	post.Trust += 1
	db.Model(&post).Update(&post)
}

func GetPost(id int64) Post {
	db := orm.GetDb()
	var post Post
	db.Preload("Site").First(&post, id)
	return post
}

func CountPost() int64 {
	db := orm.GetDb()
	var count int64
	db.Table("posts").Count(&count)
	return count
}

func CountPostBySite(id int64) int64 {
	db := orm.GetDb()
	var count int64
	db.Table("posts").Where("site_id = ?", id).Count(&count)
	return count
}

func GetPostsBySite(id int64, offset int, postsPerPage int) ([]Post, Site) {
	var posts []Post
	db := orm.GetDb()
	var site Site
	db.First(&site, id)
	db.Preload("Site").
		Offset(offset).Limit(postsPerPage).
		Order("id desc").
		Where("site_id = ?", id).
		Find(&posts)
	return posts, site
}

func GetPostsByCountry(link string, offset int, postsPerPage int) []Post {
	var posts []Post
	db := orm.GetDb()
	var country Country
	if db.Where("link = ?", link).Find(&country).Error != nil {
		return nil
	}
	db.Preload("Site", "sites.country_id = ?", country.ID).
		Joins("JOIN sites on sites.id = posts.site_id").
		Joins("JOIN countries on countries.id = sites.country_id").
		Where("countries.id = ?", country.ID).
		Where("posts.category = ?", 0).
		Offset(offset).Limit(postsPerPage).
		Order("id desc").
		Find(&posts)
	return posts
}

func GetPosts() []Post {
	db := orm.GetDb()
	var posts []Post
	db.Order("id desc").
		Preload("Site").
		Find(&posts)
	return posts
}

func GetListPostsByOffsetAndLimit(offset int, postsPerPage int) []Post {
	db := orm.GetDb()
	var posts []Post
	db.Preload("Site").
		Where("category = ?", 0).
		Where("site_id != ?", 92).
		Order("id desc").
		Offset(offset).Limit(postsPerPage).
		Find(&posts)
	return posts
}

func GetPopularPost() []Post {
	db := orm.GetDb()
	var posts []Post
	today := time.Now().Format("2006-01-02")
	db.Where("created_at LIKE ?", today+"%").
		Where("category = ?", 0).
		Where("site_id != ?", 12).
		Where("site_id != ?", 11).
		Where("site_id != ?", 13).
		Where("site_id != ?", 226).
		Order("trust desc").
		Limit(10).
		Preload("Site").
		Find(&posts)
	return posts
}

func GetPopularPostByOffsetAndLimit(offset int, postsPerPage int) []Post {
	db := orm.GetDb()
	var posts []Post
	today := time.Now().Format("2006-01-02")
	db.Where("created_at LIKE ?", today+"%").
		Where("category = ?", 0).
		Where("site_id != ?", 12).
		Where("site_id != ?", 11).
		Where("site_id != ?", 13).
		Where("site_id != ?", 226).
		Offset(offset).Limit(postsPerPage).
		Order("trust desc").
		Preload("Site").
		Find(&posts)
	return posts
}

func GetPostsByCategory(id int, limit int) []Post {
	var posts []Post
	db := orm.GetDb()
	db.Preload("Site").
		Limit(limit).
		Order("id desc").
		Where("category = ?", id).
		Find(&posts)
	return posts
}
