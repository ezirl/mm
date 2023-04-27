package models

import (
	"github.com/ezirl/massmedia/orm"
	"github.com/jinzhu/gorm"
)

func init() {
	orm.GetDb().AutoMigrate(&Country{})
}

type Country struct {
	gorm.Model
	Name string
	Link string
	Telegram string
}

func GetCountries() []Country {
	db := orm.GetDb()
	var countries []Country
	db.Find(&countries)
	return countries
}

func GetCountry(link string) Country {
	db := orm.GetDb()
	var country Country
	db.Where("link = ?", link).Find(&country)
	return country
}

func AddCountry(c *Country) error {
	db := orm.GetDb()
	if err := db.Create(&c).Error; err != nil {
		return err
	}
	return nil
}
