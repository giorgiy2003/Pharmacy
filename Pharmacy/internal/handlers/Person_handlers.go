package Handler

import (
	"fmt"
	Logic "myapp/internal/logic"
	Repository "myapp/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Форма авторизации
func Authorization(c *gin.Context) {
	c.HTML(200, "Authorization", nil)
}

//Обработичик авторизации
func Form_handler_Authorization(c *gin.Context) {
	login := c.Request.FormValue("Email")
	password := c.Request.FormValue("Password1")
	err := Logic.Autorization(login, password)
	if err != nil {
		c.HTML(200, "Authorization", gin.H{"err": err.Error()}) //Вывод ошибки
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}

//Форма регистрации
func Registration(c *gin.Context) {
	c.HTML(200, "Registration", nil)
}

//Обработичик регистрации
func Form_handler_Registration(c *gin.Context) {
	UserName := c.Request.FormValue("UserName")
	UserEmail := c.Request.FormValue("Email")
	UserPassword1 := c.Request.FormValue("Password1")
	UserPassword2 := c.Request.FormValue("Password2")
	Checkbox := c.Request.FormValue("Check")

	fmt.Println(UserName, UserEmail, UserPassword1,UserPassword2, Checkbox)

	err := Logic.Registration(UserName, UserEmail,UserPassword1,UserPassword2, Checkbox)
	if err != nil {
		c.HTML(200, "Registration", gin.H{"err": err.Error()}) //Вывод ошибки
		return
	}
	c.Redirect(http.StatusSeeOther, "/Authorization")
}

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
	id := c.Query("id")
	Products, err := Logic.ReadOneProductById(id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "shop_single", gin.H{
		"Products": Products,
	})
}

//Поиск товара
func SearhProduct(c *gin.Context) {
	productName := c.Request.FormValue("productName")
	Products, err := Logic.SearhProduct(productName)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	if len(Products) == 0 {
		c.HTML(200, "InfoPage", gin.H{
			"Info": "По Вашему запросу ничего не найдено",
		})
		return
	}
	c.HTML(200, "shop", gin.H{
		"Products": Products,
	})
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

//Использовать купон
func UseCoupon(c *gin.Context) {

	c.HTML(200, "checkout", nil)
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

//Фильтр товаров по наименованию
func NameASC(c *gin.Context) {
	Products, err := Logic.NameASC()
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

//Фильтр товаров по наименованию
func NameDESC(c *gin.Context) {
	Products, err := Logic.NameDESC()
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

//Фильтр товаров по цене
func PriceASC(c *gin.Context) {
	Products, err := Logic.PriceASC()
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

//Фильтр товаров по цене
func PriceDESC(c *gin.Context) {
	Products, err := Logic.PriceDESC()
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
