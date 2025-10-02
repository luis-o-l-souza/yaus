package controllers

import (
	"net/http"
	"yaus/services/model"
	"yaus/services/repositories"

	"github.com/gin-gonic/gin"
)

type ShortenController struct {
	u repositories.UrlMapRepository
}

func NewShortenController(u repositories.UrlMapRepository) *ShortenController {
	return &ShortenController{u: u}
}

func (s *ShortenController) Shorten(c *gin.Context) {
	var body model.UrlMapPayload
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	shortedUrl, err := s.u.Create(body)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"shortedUrl": shortedUrl,
		})
		return
	}
}
