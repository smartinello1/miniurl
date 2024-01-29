package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.Static("/app", "../frontend/miniurl/dist")
	r.Static("/assets", "../frontend/miniurl/dist/assets")
	r.StaticFile("./favicon.ico", "../frontend/miniurl/dist/favicon.ico")

	// r.Use(favicon.New("../frontend/miniurl/dist/favicon.ico"))

	r.NoRoute(func(ctx *gin.Context) {
		log.Print("Not Found route")
		ctx.File("../frontend/miniurl/dist/index.html")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"redirect": true,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
