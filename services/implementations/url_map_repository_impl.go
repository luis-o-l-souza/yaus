package implementations

import (
	"database/sql"
	"yaus/services/model"
	"yaus/services/repositories"
)

type urlMapRepository struct {
	db *sql.DB
}

func NewUrlMapRepository(db *sql.DB) repositories.UrlMapRepository {
	return &urlMapRepository{db: db}
}

func (r *urlMapRepository) Create(payload *model.UrlMapPayload) (string, error) {
	return "", nil
}

func (r *urlMapRepository) GetByShortUrl(shortUrl string) (*model.UrlMap, error) {
	return &model.UrlMap{}, nil
}
