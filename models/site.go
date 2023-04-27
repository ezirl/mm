package models

import (
	"github.com/ezirl/mm/orm"
	"github.com/jinzhu/gorm"
)

func init() {
	orm.GetDb().AutoMigrate(&Site{})
}

type Site struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Link      string `gorm:"not null"`
	Rss       string `gorm:"not null;unique"`
	Ico       string
	Country   Country
	Category  int
	CountryID uint
}

func ListSites() []Site {
	db := orm.GetDb()
	var sites []Site
	db.Preload("Country").
		Order("name desc").
		Find(&sites)
	return sites
}

func GetSite(id int64) Site {
	db := orm.GetDb()
	var site Site
	db.Preload("Country").First(&site, id)
	return site
}

func DeleteSite(id int64) error {
	db := orm.GetDb()
	err := db.Where("id = ?", id).Delete(&Site{}).Error
	err = db.Where("site_id = ?", id).Delete(&Post{}).Error
	if err != nil {
		return err
	}
	return nil
}

func AddSite(site *Site) error {
	db := orm.GetDb()
	if err := db.Create(&site).Error; err != nil {
		return err
	}
	return nil
}

func UpdateSite(id int, site *Site) error {
	db := orm.GetDb()
	var sit Site
	db.First(&sit, id)
	sit.Link = site.Link
	sit.Rss = site.Rss
	sit.Name = site.Name
	sit.Ico = site.Ico
	sit.Category = site.Category
	sit.CountryID = site.CountryID
	if err := db.Model(&sit).Update(&sit).Error; err != nil {
		return err
	}
	return nil
}
