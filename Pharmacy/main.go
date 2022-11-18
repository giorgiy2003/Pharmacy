package main

import (
	Handler "myapp/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/css", "Frontend/css")
	//router.Static("/scss", "Frontend/scss")
	//router.Static("/fonts", "Frontend/fonts")
	router.Static("/images", "Frontend/images")
	router.Static("/js", "Frontend/js")
	router.LoadHTMLGlob("Frontend/*.html")
	router.Use(Handler.ConnectDB())
	router.GET("/", Handler.MainForm)
	router.GET("/shop", Handler.Shop)
	

	//router.GET("/Form_handler_GetById", Handler.Form_handler_GetById)
	router.Run("localhost:8080")
}