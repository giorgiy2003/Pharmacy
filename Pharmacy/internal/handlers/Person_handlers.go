package Handler

import (
	"log"
	Logic "myapp/internal/logic"
	Repository "myapp/internal/repository"

	"github.com/gin-gonic/gin"
)

//Товары
func Shop(c *gin.Context) {
	Products, err := Logic.ReadAllProducts()
	log.Println(Products)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "shop", gin.H{
		"Products": Products,
	})
}

//Просмотреть товар
func Shop_single(c *gin.Context) {

	c.HTML(200, "shop_single", nil)
}

//Главная форма
func MainForm(c *gin.Context) {
	c.HTML(200, "index", nil)
}

//Корзина
func Cart(c *gin.Context) {
	c.HTML(200, "cart", nil)
}

//О нас
func About(c *gin.Context) {
	c.HTML(200, "about", nil)
}

//Контакты
func Contact(c *gin.Context) {
	c.HTML(200, "contact", nil)
}

//Страница оформления заказа
func Checkout(c *gin.Context) {
	c.HTML(200, "checkout", nil)
}

//Сделать заказ
func Make_Order(c *gin.Context) {
	c.HTML(200, "thankyou", nil)
}

//Добавить в корзину
func AddToCart(c *gin.Context) {

	c.HTML(200, "index", nil)
}

//Убрать из корзины
func DeleteFromCart(c *gin.Context) {

	c.HTML(200, "cart", nil)
}

//Оставить отзыв
func SendMessage(c *gin.Context) {

	c.HTML(200, "index", nil)
}

//Обезболивабщие лекарства
func Painkillers_medicines(c *gin.Context) {

	c.HTML(200, "shop", nil)
}

//Иммуностимулирующие лекарства
func Immunostimulating_medicines(c *gin.Context) {

	c.HTML(200, "shop", nil)
}

//Жаропонижающие лекарства
func Antipyretic_medicines(c *gin.Context) {

	c.HTML(200, "shop", nil)
}

//Лекарства от гриппа и простуды
func Flu_medicines(c *gin.Context) {

	c.HTML(200, "shop", nil)
}

//Лекарства от грибковых заболеваний
func Fungal_medicines(c *gin.Context) {

	c.HTML(200, "shop", nil)
}

//Лекарства от аллергии
func Allergies_medicines(c *gin.Context) {

	c.HTML(200, "shop", nil)
}

//Антибиотики
func Antibiotics_medicines(c *gin.Context) {

	c.HTML(200, "shop", nil)
}

func ConnectDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := Repository.OpenTable(); err != nil {
			c.HTML(500, "400", gin.H{
				"Error": err,
			})
			return
		}
	}
}
