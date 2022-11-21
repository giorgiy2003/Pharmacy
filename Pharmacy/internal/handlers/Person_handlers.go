package Handler

import (
	Logic "myapp/internal/logic"
	Repository "myapp/internal/repository"

	"github.com/gin-gonic/gin"
)


//Главная форма
func MainForm(c *gin.Context) {
	Products, err := Logic.ReadProductsWithLimit()
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "index", gin.H{
		"Products": Products,
	})
}


//Товары
func Shop(c *gin.Context) {
	Products, err := Logic.ReadAllProducts()
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


//Лекарства по категориям
func Medicines_by_category(c *gin.Context) {
	id := c.Param("id")
	Products, err := Logic.Medicines_by_category(id)
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
