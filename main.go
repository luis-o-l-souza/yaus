package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"yaus/controllers"
	"yaus/services/implementations"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("No env file")
	}
	var (
		host     = os.Getenv("DB_HOST")
		port     = 5432
		user     = os.Getenv("POSTGRES_DB_USER")
		password = os.Getenv("POSTGRES_DB_PWD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Print("Connected to database")
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	urlMapRepository := implementations.NewUrlMapRepository(db)
	shortenerController := controllers.NewShortenController(urlMapRepository)
	redirectController := controllers.NewRedirectController(urlMapRepository)
	r.POST("/shorten", shortenerController.Shorten)
	r.GET("/:shortedUrl", redirectController.Redirect)
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
