package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortenController struct {
	db *sql.DB
}

func NewShortenController(db *sql.DB) *ShortenController {
	return &ShortenController{db}
}

func (s *ShortenController) Shorten(c *gin.Context) {
	str := "INSERT INTO url_maps (short_url,long_url) VALUES ($1, $2)"
	_, err := s.db.Exec(str, "test", "test")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"shortedUrl": "urlTest",
		})
	}
}
