package repositories

import "yaus/services/model"

type UrlMapRepository interface {
	Create(payload *model.UrlMapPayload) (string, error)
	GetByShortUrl(shortUrl string) (*model.UrlMap, error)
}
