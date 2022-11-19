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
	router.GET("/shop_single", Handler.Shop_single)
	router.GET("/cart", Handler.Cart)
	router.GET("/about", Handler.About)
	router.GET("/checkout", Handler.Checkout)
	router.GET("/contact", Handler.Contact)
	router.GET("/Order", Handler.Make_Order)
	router.GET("/AddToCart", Handler.AddToCart)

  /*router.GET("/Painkillers_medicines", Handler.Painkillers_medicines)
	router.GET("/Immunostimulating_medicines", Handler.Immunostimulating_medicines)
	router.GET("/Antipyretic_medicines", Handler.Antipyretic_medicines)
	router.GET("/Flu_medicines", Handler.Flu_medicines)
	router.GET("/Fungal_medicines", Handler.Fungal_medicines)
	router.GET("/Allergies_medicines", Handler.Allergies_medicines)
	router.GET("/Antibiotics_medicines", Handler.Antibiotics_medicines)*/
	router.Run("localhost:8080")
}
