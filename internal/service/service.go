package service

import (
	"fmt"
	"time"
	"urlshortner/internal/constant"
	"urlshortner/internal/database"
	"urlshortner/internal/helper"
	"urlshortner/internal/models"
)

func ShortenURL(longURL string) (*models.UrlDb, error) {
	existingRecord, _ := database.Mgr.GetUrlFromLongUrl(longURL, constant.UrlCollection)

	if existingRecord.UrlCode != "" {
		return &existingRecord, nil
	}

	for {
		code := helper.GenRandomString(6)

		record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)

		if record.UrlCode == "" {
			url := &models.UrlDb{
				CreatedAt: time.Now().Unix(),
				ExpiredAt: time.Now().Add(2 * time.Minute).Unix(),
				UrlCode:   code,
				LongUrl:   longURL,
				ShortUrl:  constant.BaseUrl + code,
			}

			_, err := database.Mgr.Insert(*url, constant.UrlCollection)
			if err != nil {
				return nil, err
			}

			return url, nil
		}
	}
}

func GetLongURL(code string) (*models.UrlDb, error) {
	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)

	if record.UrlCode == "" {
		return nil, fmt.Errorf("there is no URL found")
	}

	return &record, nil
}
