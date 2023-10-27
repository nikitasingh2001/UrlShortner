package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
	"urlshortner/internal/constant"
	"urlshortner/internal/database"
	"urlshortner/internal/helper"
	"urlshortner/internal/logger"
	"urlshortner/internal/models"
)

func ShortenURL(longURL string) (*models.UrlDb, error) {
	existingRecord, _ := database.Mgr.GetUrlFromLongUrl(longURL, constant.UrlCollection)

	if existingRecord.UrlCode != "" {
		logger.Log.WithFields(logrus.Fields{"longURL": longURL}).Info("ShortenURL: URL already exists in the database")
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
				logger.Log.WithFields(logrus.Fields{"longURL": longURL, "error": err}).Error("ShortenURL: Failed to insert URL into the database")
				return nil, err
			}

			logger.Log.WithFields(logrus.Fields{"longURL": longURL}).Info("ShortenURL: Successfully shortened URL")
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
