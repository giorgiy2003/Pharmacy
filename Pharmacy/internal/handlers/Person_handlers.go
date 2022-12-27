package Handler

import (
	"fmt"
	"log"
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
	err := Logic.Registration(UserName, UserEmail, UserPassword1, UserPassword2, Checkbox)
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
		"Role":     Logic.Role,
		"User_id":  Logic.User_id,
		"Products": Products,
	})
}

//Выйти из аккаунта
func Sign_out(c *gin.Context) {
	Logic.User_id = 0
	c.Redirect(http.StatusSeeOther, "/")
}

//Страница разработчика
func Admin(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	c.HTML(200, "Developer_page", nil)
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
		"Role":     Logic.Role,
		"User_id":  Logic.User_id,
		"Products": Products,
	})
}

//Просмотреть товар
func Shop_single(c *gin.Context) {
	id := c.Query("id")
	Products, err := Logic.ShopSingle(id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	proverka1, err := Logic.Proverka1(id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}

	proverka2, err := Logic.Proverka2(id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.HTML(200, "shop_single", gin.H{
		"Proverka1": proverka1,
		"Proverka2": proverka2,
		"Role":      Logic.Role,
		"User_id":   Logic.User_id,
		"Products":  Products,
	})
}

//Добавить в избранное
func AddToFavotites(c *gin.Context) {
	id := c.Param("id")
	err := Logic.AddToFavotites(id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/shop_single?id=%s", id))
}

//Убрать из избранного
func DeleteFromFavotites(c *gin.Context) {
	id := c.Param("id")
	err := Logic.DeleteFromFavotites(id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/shop_single?id=%s", id))
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
			"Role":    Logic.Role,
			"User_id": Logic.User_id,
			"Info":    "По Вашему запросу ничего не найдено",
		})
		return
	}
	c.HTML(200, "shop", gin.H{
		"Role":     Logic.Role,
		"User_id":  Logic.User_id,
		"Products": Products,
	})
}

//Корзина
func Cart(c *gin.Context) {
	Products, total, err := Logic.UserCart()
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	if len(Products) == 0 {
		c.HTML(200, "NullCart", gin.H{
			"Role":    Logic.Role,
			"User_id": Logic.User_id,
		})
		return
	}
	c.HTML(200, "cart", gin.H{
		"Role":     Logic.Role,
		"Products": Products,
		"Total":    total,
		"User_id":  Logic.User_id,
	})
}

//Товары в избранном
func Favourites(c *gin.Context) {
	Products, err := Logic.Favourites()
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "favouritesPage", gin.H{
		"Role":     Logic.Role,
		"Products": Products,
		"User_id":  Logic.User_id,
	})
}

//О нас
func About(c *gin.Context) {
	c.HTML(200, "about", gin.H{
		"Role":    Logic.Role,
		"User_id": Logic.User_id,
	})
}

//Контакты
func Contact(c *gin.Context) {
	c.HTML(200, "contact", gin.H{
		"Role":    Logic.Role,
		"User_id": Logic.User_id,
	})
}

//Страница оформления заказа
func Checkout(c *gin.Context) {
	Products, total, err := Logic.UserCart()
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	if Products == nil {
		c.HTML(200, "400", gin.H{
			"Error": "Невозможно оформить заказ, в корзине ничего нет!",
		})
		return
	}
	c.HTML(200, "checkout", gin.H{
		"Role":     Logic.Role,
		"Products": Products,
		"Total":    total,
		"User_id":  Logic.User_id,
	})
}

//Сделать заказ
func Order(c *gin.Context) {
	
	c_fname := c.Request.FormValue("c_fname")
	c_lname := c.Request.FormValue("c_lname")
	c_patronymic := c.Request.FormValue("c_patronymic")
	c_address := c.Request.FormValue("c_address")
	c_email_address := c.Request.FormValue("c_email_address")
	c_phone := c.Request.FormValue("c_phone")
	c_order_notes := c.Request.FormValue("c_order_notes")
	log.Println(c_fname, c_lname, c_patronymic, c_address, c_email_address, c_phone, c_order_notes)
	
	err := Logic.MakeOrder()
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "thankyou", gin.H{
		"Role":    Logic.Role,
		"User_id": Logic.User_id,
	})
}

//Доставки
func HistoryPage(c *gin.Context) {

	Orders, err := Logic.Orders()
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.HTML(200, "OrdersPage", gin.H{
		"Role":    Logic.Role,
		"Orders":  Orders,
		"User_id": Logic.User_id,
	})
}

//Информация о заказе
func Order_details(c *gin.Context) {
	order := c.Param("order")
	Orders, err := Logic.Order_details(order)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "Order_details", gin.H{
		"Role":    Logic.Role,
		"Orders":  Orders,
		"User_id": Logic.User_id,
	})
}

//Добавить в корзину
func AddToCart(c *gin.Context) {
	id := c.Param("id")
	err := Logic.AddToCart(id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/shop_single?id=%s", id))
}

//Убрать из корзины
func DeleteFromCart(c *gin.Context) {
	id := c.Param("id")
	err := Logic.DeleteFromCart(id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/cart")
}

//Уменьшить количество товара в корзине на странице товара
func MinusKoll(c *gin.Context) {
	id := c.Param("id")
	koll := c.Request.FormValue("koll")
	err := Logic.MinusKoll(id, koll)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/shop_single?id=%s", id))
}

//Увеличить количество товара в корзине на странице товара
func AddKoll(c *gin.Context) {
	id := c.Param("id")
	koll := c.Request.FormValue("koll")
	err := Logic.AddKoll(id, koll)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/shop_single?id=%s", id))
}

//Уменьшить количество товара в корзине
func MinusKollinCart(c *gin.Context) {
	id := c.Param("id")
	koll := c.Request.FormValue("koll")
	err := Logic.MinusKoll(id, koll)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/cart")
}

//Увеличить количество товара в корзине
func AddKollinCart(c *gin.Context) {
	id := c.Param("id")
	koll := c.Request.FormValue("koll")
	err := Logic.AddKoll(id, koll)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/cart")
}

//Оставить отзыв
func SendMessage(c *gin.Context) {
	c.HTML(200, "index", gin.H{
		"Role":    Logic.Role,
		"User_id": Logic.User_id,
	})
}

//Лекарства по категориям
func Medicines_by_category(c *gin.Context) {
	category := c.Param("category")
	Products, err := Logic.Medicines_by_category(category)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "shop", gin.H{
		"Role":     Logic.Role,
		"User_id":  Logic.User_id,
		"Category": category,
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
		"Role":     Logic.Role,
		"User_id":  Logic.User_id,
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
		"Role":     Logic.Role,
		"User_id":  Logic.User_id,
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
		"Role":     Logic.Role,
		"User_id":  Logic.User_id,
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
		"Role":     Logic.Role,
		"User_id":  Logic.User_id,
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
