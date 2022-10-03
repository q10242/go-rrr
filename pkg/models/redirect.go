package models

import (
	"fmt"
	"go-rrr/pkg/config"
	"go-rrr/pkg/utils"
	"math/rand"
	"net/url"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Redirect struct {
	gorm.Model
	OriginUrl   string `gorm: ""json:"originUrl"`
	ShortUrl    string `json: shortUrl"`
	Probability int    `json: "probability"`
}

func init() {
	config.Connect()

	db = config.GetDB()
	db.AutoMigrate(&Redirect{})
}
func (r *Redirect) CreateRedirect() (*Redirect, error) {
	db.NewRecord(r)
	fmt.Println(r)

	OriginUrl, err := url.ParseRequestURI(r.OriginUrl)
	if err != nil {
		return nil, err
	}
	r.OriginUrl = OriginUrl.String()
	ShortString := randStringFinder()
	r.ShortUrl = ShortString
	db.Create(&r)
	return r, nil
}
func GetRedirectByShortUrl(ShortUrl string) (*Redirect, *gorm.DB) {
	var getRedirect Redirect
	db := db.Where("short_url=?", ShortUrl).Find(&getRedirect)
	if RedirectToRicRoll(int(getRedirect.Probability)) {
		getRedirect.OriginUrl = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
	}
	return &getRedirect, db
}

func randStringFinder() string {
	ShortUrl := utils.RandStringRunes(6)
	var getRedirect Redirect
	db := db.Where("short_url=?", ShortUrl).Find(&getRedirect)
	for !db.RecordNotFound() {
		var getRedirect Redirect
		ShortUrl := utils.RandStringRunes(6)
		db := db.Where("short_url=?", ShortUrl).Find(&getRedirect)
		if db.RecordNotFound() {
			break
		}
	}
	return ShortUrl
}

func Clean() {

	now := time.Now()

	fmt.Println(now)
	lastMounth := now.AddDate(0, -1, 0)
	var delRedirect Redirect
	db.Where("updated_at < ?", lastMounth).Unscoped().Delete(&delRedirect)
	fmt.Println("Cleaned records over 1 mounth")
}

func RedirectToRicRoll(Rate int) bool {
	result := rand.Intn(99)
	return Rate >= result
}
