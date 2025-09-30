package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RedirectController struct {
	db *sql.DB
}

func NewRedirectController(db *sql.DB) *RedirectController {
	return &RedirectController{db}
}

func (r *RedirectController) Redirect(c *gin.Context) {
	shortUrl, exists := c.Params.Get("shortedUrl")
	var longUrl string
	if exists {
		longUrlRow := r.db.QueryRow("SELECT u.long_url from url_maps u WHERE u.short_url = $1", shortUrl)
		err := longUrlRow.Scan(&longUrl)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		}
		c.Redirect(http.StatusPermanentRedirect, longUrl)
	}
}
