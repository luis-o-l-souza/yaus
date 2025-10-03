package controllers

import (
	"net/http"

	"yaus/services/repositories"

	"github.com/gin-gonic/gin"
)

type RedirectController struct {
	u repositories.UrlMapRepository
}

func NewRedirectController(u repositories.UrlMapRepository) *RedirectController {
	return &RedirectController{u: u}
}

func (r *RedirectController) Redirect(c *gin.Context) {
	shortUrl, exists := c.Params.Get("shortedUrl")
	if exists {
		longUrl, err := r.u.GetByShortUrl(shortUrl)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.Redirect(http.StatusPermanentRedirect, longUrl.LongUrl)
	}
}
