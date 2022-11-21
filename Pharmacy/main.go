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
	router.GET("/", Handler.MainForm) //Главная страница
	router.GET("/shop", Handler.Shop) //Магазин
	router.GET("/shop_single", Handler.Shop_single) //Просмотр карточки товара
	router.GET("/cart", Handler.Cart) //Корзина
	router.GET("/about", Handler.About) //О нас
	router.GET("/checkout", Handler.Checkout) //Страница оформления
	router.GET("/contact", Handler.Contact) //Контакты
	router.GET("/Order", Handler.Make_Order) //Оформить заказ
	router.GET("/AddToCart", Handler.AddToCart) //Добавить в корзину
	router.GET("/DeleteFromCart", Handler.DeleteFromCart) //Убрать из корзины
	router.GET("/SendMessage", Handler.SendMessage)	//Оставить отзыв

	router.GET("/Medicines_by_category/:id", Handler.Medicines_by_category) // Поиск по категориям
	router.Run("localhost:8080")
}
