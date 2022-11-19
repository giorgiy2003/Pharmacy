package Handler

import (
	Repository "myapp/internal/repository"

	"github.com/gin-gonic/gin"
)

func Shop(c *gin.Context) {
	/*Products, err := Logic.ReadAll()
	if err != nil {
		log.Println(err)
		c.HTML(400, "InfoPage", nil)
		return
	}*/
	c.HTML(200, "shop", gin.H{})
	/*for _, Product := range Products {
		c.HTML(200, "shop", gin.H {
			"Id":           Product.Id,
			"Product_name": Product.Product_name,
			"Manufacturer": Product.Manufacturer,
			"Category":     Product.Category,
			"Description":  Product.Description,
			"Price":  Product.Price,
		})
	}*/
	return
}

func ConnectDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := Repository.OpenTable(); err != nil {
			c.HTML(500, "Connection_failed", gin.H{
				"Error": err,
			})
			return
		}
	}
}

//Главная форма
func MainForm(c *gin.Context) {
	c.HTML(200, "index", nil)
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
