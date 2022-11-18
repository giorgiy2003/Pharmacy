package main

import (
	Handler "myapp/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/css", "Frontend/css")
	router.Static("/scss", "Frontend/scss")
	router.Static("/fonts", "Frontend/fonts")
	router.Static("/images", "Frontend/images")
	router.Static("/js", "Frontend/js")
	router.LoadHTMLGlob("Frontend/*.html")
	router.Use(Handler.ConnectDB())
	router.GET("/", Handler.MainForm)
	router.GET("/shop", Handler.Shop)
	router.GET("/shop-single", Handler.Shop_single)
	router.GET("/cart", Handler.Cart)
	router.GET("/about", Handler.About)
	router.GET("/checkout", Handler.Checkout)
	router.GET("/contact", Handler.Contact)
	router.GET("/thankyou", Handler.Thanks)

  /*router.GET("/Form_handler_Painkillers", Handler.Form_handler_Painkillers)
	router.GET("/Form_handler_Immunostimulating", Handler.Form_handler_Immunostimulating)
	router.GET("/Form_handler_Antipyretic", Handler.Form_handler_Antipyretic)
	router.GET("/Form_handler_Flu", Handler.Form_handler_Flu)
	router.GET("/Form_handler_Fungal", Handler.Form_handler_Fungal)
	router.GET("/Form_handler_Allergies", Handler.Form_handler_Allergies)
	router.GET("/Form_handler_Antibiotics", Handler.Form_handler_Antibiotics)*/
	router.Run("localhost:8080")
}
