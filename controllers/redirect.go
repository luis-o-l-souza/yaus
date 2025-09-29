package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {

	c.Redirect(http.StatusPermanentRedirect, "https://google.com")
}
