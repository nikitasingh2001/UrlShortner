package service

import (
	"fmt"
	"time"
	"urlshortner/internal/constant"
	"urlshortner/internal/database"
	"urlshortner/internal/helper"
	"urlshortner/internal/types"
)

func ShortenURL(longURL string) (*types.UrlDb, error) {
	existingRecord, _ := database.Mgr.GetUrlFromLongUrl(longURL, constant.UrlCollection)

	if existingRecord.UrlCode != "" {
		return &existingRecord, nil
	}

	for {
		code := helper.GenRandomString(6)

		record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)

		if record.UrlCode == "" {
			url := &types.UrlDb{
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

			go func(code string) {
				time.Sleep(2 * time.Minute)
				database.Mgr.DeleteUrlByCode(code, constant.UrlCollection)
			}(code)

			return url, nil
		}
	}
}

func GetLongURL(code string) (*types.UrlDb, error) {
	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)

	if record.UrlCode == "" {
		return nil, fmt.Errorf("there is no URL found")
	}

	return &record, nil
}
