package telegram

import (
	"errors"
	"fmt"
	"github.com/ezirl/massmedia/models"
	"github.com/ezirl/massmedia/orm"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var urlString = "https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s&parse_mode=html"
var token = "910620008:AAFa7YBGdD167lDicT_2_eswmFNs-6lWKm8"


func Start() {
	for {
		if time.Now().Format("04:05") == "00:00" {
			countries := models.GetCountries()
			for _, country := range countries {
				mes, err := CreateMessage(country.Link)
				if err == nil {
					SendMessage(country.Telegram, mes)
				}
				_err(err)
			}
		}
		time.Sleep(time.Second)
	}
}

func SendMessage(channel string, message string) {
	sendMess := fmt.Sprintf(urlString, token, channel, message)
	_, err := http.Get(sendMess)
	_err(err)
}

func CreateMessage(country string) (string, error) {
	times := Format(time.Now()) + time.Now().Format("15:04")
	posts, countryModel := getPopularPostsByCountryOrderTrust(country, 10)

	mess := `<a href="http://massmedia24.com/`+countryModel.Link+`">Massmedia24</a> `
	mess += times
	post := formatPosts(posts)
	if len(post) == 0 {
		return	"", errors.New("have not posts")
	}
	mess += post

	mess += "\n\nОставайся в курсе самых свежих новостей. Переходи на наш "
	mess += `<a href="http://massmedia24.com/`+countryModel.Link+`">сайт!</a>!`
	fmt.Println(mess)
	mess = url.QueryEscape(mess)
	return mess, nil
}

func formatPosts(posts []models.Post) string{
	var mess string
	for _, post := range posts {
		mess += "\n\n ⭐" + formatTitle(post.Title, post.Link)
	}
	return mess
}

func formatTitle(title string, link string) string {
	words := strings.Split(title, " ")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(words); i++{
		random := rand.Intn(len(words))
		if len(words[random]) > 5 {
			words[random] = `<a href="`+link+`">`+words[random]+`</a>`
			break
		}
	}
	res := strings.Join(words, " ")
	return res
}


func getPopularPostsByCountryOrderTrust(link string, postsPerPage int) ([]models.Post, models.Country) {
	var posts []models.Post
	db := orm.GetDb()
	var country models.Country
	if db.Where("link = ?", link).Find(&country).Error != nil {
		return nil, country
	}
	db.Preload("Site", "sites.country_id = ?", country.ID).
		Joins("JOIN sites on sites.id = posts.site_id").
		Joins("JOIN countries on countries.id = sites.country_id").
		Where("countries.id = ?", country.ID).
		//Where("posts.category = ?", 0).
		Where("posts.created_at >= date_sub(now(), INTERVAL 1 HOUR)").
		Limit(postsPerPage).
		Order("trust desc").
		Find(&posts)
	return posts, country
}

//time format&localization
func Format(t time.Time) string {
	return fmt.Sprintf("%d %s, ",
		t.Day(), months[t.Month()-1],
	)
}

var months = [...]string{
	"Января", "Февраля", "Марта", "Апреля", "Мая", "Июня", "Июля", "Августа", "Сентября", "Октября", "Ноября", "Декабря",
}

func _err(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
