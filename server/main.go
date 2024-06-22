package main

import (
	b64 "encoding/base64"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	loadDotEnv()

	r := gin.Default()

	// Serve static web app folders and files
	r.Static("/app", "../frontend/miniurl/dist")
	r.Static("/assets", "../frontend/miniurl/dist/assets")
	r.StaticFile("./favicon.ico", "../frontend/miniurl/dist/favicon.ico")

	// Handle 404 routes
	r.NoRoute(func(ctx *gin.Context) {
		log.Print("Not Found route")
		ctx.File("../frontend/miniurl/dist/index.html")
	})

	// Ping server
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Handle redirect, if mapped url found then redirect, otherwise redirect to not found page
	r.GET("/:miniurl", func(ctx *gin.Context) {
		log.Println("url: ", ctx.Param("miniurl"))
		urlSites := map[string]string{"test": "https://google.com"}
		redirectUrl := urlSites[ctx.Param("miniurl")]
		log.Println("redirectUrl: ", redirectUrl)
		if redirectUrl == "" {
			ctx.File("../frontend/miniurl/dist/index.html/")
			return
		}

		ctx.Redirect(301, redirectUrl)
	})

	// LOGIN API
	r.POST("/api/v1/login", func(ctx *gin.Context) {
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":        true,
				"errorMessage": err.Error(),
			})
		}
		bodyStr := string(body)

		bodyDecoded, _ := b64.StdEncoding.DecodeString(bodyStr)
		bodyDecodedStr := string(bodyDecoded)

		username := strings.Split(bodyDecodedStr, ":")[0]
		password := strings.Split(bodyDecodedStr, ":")[1]

		log.Println("username: ", username)
		log.Println("password: ", password)

		ctx.JSON(http.StatusOK, gin.H{
			"error":        false,
			"access_token": "testtoken",
			"userId":       "1234",
		})
	})

	r.Run()
}

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
