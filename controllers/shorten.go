package controllers

import (
	"fmt"
	"net/http"
	"net/url"

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
		errorMsg := err.Error()
		if errorMsg == "EOF" {
			errorMsg = "Invalid JSON"
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMsg,
		})
		return
	}
	if isValid := validateUrl(body.LongUrl); !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid URL format",
		})
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

func validateUrl(inputUrl string) bool {
	_, err := url.ParseRequestURI(inputUrl)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
