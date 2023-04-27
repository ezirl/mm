package models

import (
	"github.com/ezirl/mm/orm"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name    string
	Link    string
	Rss     string
	Email   string
	Comm    string
	Country string
}

func init() {
	orm.GetDb().AutoMigrate(&Contact{})
}

func AddContact(cont *Contact) error {
	db := orm.GetDb()
	if err := db.Create(&cont).Error; err != nil {
		return err
	}
	return nil
}

func ListContacts() []Contact {
	db := orm.GetDb()
	var contacts []Contact
	db.Order("created_at desc").Find(&contacts)
	return contacts
}

func DelContact(id int64) error {
	db := orm.GetDb()
	if err := db.Where("id = ?", id).Delete(&Contact{}).Error; err != nil {
		return err
	}
	return nil
}

func DelSpam() error {
	db := orm.GetDb()
	if err := db.Where("comm like '%http%'").Delete(Contact{}).Error; err != nil {
		return err
	}
	return nil
}
