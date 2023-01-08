package Handler

import (
	"fmt"
	"log"
	Logic "myapp/internal/logic"
	Model "myapp/internal/model"
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

	Popular_Products, err := Logic.Popular_Products()
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "index", gin.H{
		"Role":             Logic.Role,
		"User_id":          Logic.User_id,
		"Products":         Products,
		"Popular_Products": Popular_Products,
	})
}

//Выйти из аккаунта
func Sign_out(c *gin.Context) {
	Logic.User_id = 0
	Logic.Role = ""
	c.Redirect(http.StatusSeeOther, "/")
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
	var delivery = 150

	Products, sumPrice, err := Logic.UserCart()
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
	if sumPrice >= 1000 {
		delivery = 0
	}
	total := delivery + sumPrice

	c.HTML(200, "checkout", gin.H{
		"Role":     Logic.Role,
		"Products": Products,
		"SumPrice": sumPrice,
		"Total":    total,
		"Delivery": delivery,
		"User_id":  Logic.User_id,
	})
}

//Сделать заказ
func Order(c *gin.Context) {
	c_city := c.Request.FormValue("c_city")
	c_fname := c.Request.FormValue("c_fname")
	c_lname := c.Request.FormValue("c_lname")
	c_patronymic := c.Request.FormValue("c_patronymic")
	c_address := c.Request.FormValue("c_address")
	c_email_address := c.Request.FormValue("c_email_address")
	c_phone := c.Request.FormValue("c_phone")
	c_order_notes := c.Request.FormValue("c_order_notes")

	err := Logic.MakeOrder(c_city, c_fname, c_lname, c_patronymic, c_address, c_email_address, c_phone, c_order_notes)
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

//Оставить отзыв
func SendMessage(c *gin.Context) {
	var Comment Model.Comment
	Comment.Customer_FirstName = c.Request.FormValue("c_fname")
	Comment.Customer_LastName = c.Request.FormValue("c_lname")
	Comment.Customer_Email = c.Request.FormValue("c_email")
	Comment.Theme = c.Request.FormValue("c_subject")
	Comment.Comment = c.Request.FormValue("c_message")

	err := Logic.CreateComment(Comment)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "thankyou2", gin.H{
		"Role":    Logic.Role,
		"User_id": Logic.User_id,
	})
}




//Для администратора


//Заказы
func Orders_Page(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	category := c.Param("order_status")
	Orders, err := Logic.Orders_Page(category)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "AdminOrdersPage", gin.H{
		"Orders": Orders,
		"Status": category,
	})
}

//Поменять статус заказа
func Change_status(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	order_id := c.Param("order_id")
	order_status := c.Param("order_status")
	err := Logic.Change_status(order_status, order_id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
}


//Сотрудники


//Cписок всех сотрудников
func Get_All_Workers(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	Workers, err := Logic.ReadAllWorkers()
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "AllWorkers", gin.H{
		"Workers": Workers,
	})
}

//Добавить сотрудника
func Form_handler_PostWorker(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	var newWorker Model.Worker
	newWorker.Worker_FirstName = c.Request.FormValue("Worker_FirstName")
	newWorker.Worker_LastName = c.Request.FormValue("Worker_LastName")
	newWorker.Worker_Email = c.Request.FormValue("Worker_Email")
	newWorker.Worker_Phone = c.Request.FormValue("Worker_Phone")
	newWorker.Post = c.Request.FormValue("Post")
	newWorker.Salary_per_month = c.Request.FormValue("Salary_per_month")

	err := Logic.CreateWorker(newWorker)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/Get_All_Workers")
}

//Редактировать запись сотрудника
func Form_handler_UpdateWorkerById(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	var newWorker Model.Worker
	id := c.Request.FormValue("id")
	newWorker.Worker_FirstName = c.Request.FormValue("Worker_FirstName")
	newWorker.Worker_LastName = c.Request.FormValue("Worker_LastName")
	newWorker.Worker_Email = c.Request.FormValue("Worker_Email")
	newWorker.Worker_Phone = c.Request.FormValue("Worker_Phone")
	newWorker.Post = c.Request.FormValue("Post")
	newWorker.Salary_per_month = c.Request.FormValue("Salary_per_month")

	err := Logic.UpdateWorker(newWorker, id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/Get_All_Workers")
}

//Удалить сотрудника из базы
func Form_handler_DeleteWorkerById(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	id := c.Request.FormValue("id")
	err := Logic.DeleteWorker(id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/Get_All_Workers")
}

func Add_Worker(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	c.HTML(200, "AddWorker", nil)
}
func Remove_Worker(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	c.HTML(200, "DeleteWorker", nil)
}
func Edit_Worker(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	c.HTML(200, "EditWorker", nil)
}


//Отзывы


//Все отзывы
func Get_All_Comments(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	Comments, err := Logic.ReadAllComments()
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "AllComments", gin.H{
		"Comments": Comments,
	})
}

//Отзывы со статусом "Важно"
func Get_Important_Comments(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	Comments, err := Logic.ReadImportantComments()
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "AllComments", gin.H{
		"Comments": Comments,
	})
}

//Удалить отзыв
func Remove_Comment(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	id := c.Param("comment_id")
	err := Logic.DeleteComment(id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/Get_All_Comments")
}

//Изменить статус комментария на "Важно"
func Change_To_Important(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	id := c.Param("comment_id")
	err := Logic.Change_To_Important(id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusSeeOther, "/Get_All_Comments")
}


//Продукты


//Cписок всех товаров
func Get_All_Products(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	Products, err := Logic.ReadAllProducts()
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "AllProducts", gin.H{
		"Products": Products,
	})
}

//Поиск товара
func Searh_Products(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	productName := c.Request.FormValue("productName")
	log.Println(productName)
	Products, err := Logic.SearhProduct(productName)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "AllProducts", gin.H{
		"Products": Products,
	})
}

//Лекарства по категориям
func Products_by_category(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	category := c.Param("Products_by_category")
	log.Println(category)
	Products, err := Logic.Medicines_by_category(category)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.HTML(200, "AllProducts", gin.H{
		"Products": Products,
	})
}

//Удалить товар
func Form_handler_DeleteProductById(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	id := c.Request.FormValue("product_id")
	err := Logic.Form_handler_DeleteProductById(id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/Get_All_Products")
}

//Добавить товар
func Form_handler_PostProduct(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	var newProduct Model.Product
	newProduct.Product_Name = c.Request.FormValue("Product_Name")
	newProduct.Product_Image = c.Request.FormValue("myFile")
	newProduct.Product_Manufacturer = c.Request.FormValue("Product_Manufacturer")
	newProduct.Product_Category = c.Request.FormValue("Product_Category")
	newProduct.Product_Price = c.Request.FormValue("Product_Price")
	newProduct.Product_Description = c.Request.FormValue("Product_Description")

	err := Logic.CreateProduct(newProduct)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/Get_All_Products")
}

//Редактировать запись товара
func Form_handler_UpdateProductById(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	var newProduct Model.Product
	id := c.Request.FormValue("id")
	newProduct.Product_Name = c.Request.FormValue("Product_Name")
	newProduct.Product_Image = c.Request.FormValue("myFile")
	newProduct.Product_Manufacturer = c.Request.FormValue("Product_Manufacturer")
	newProduct.Product_Category = c.Request.FormValue("Product_Category")
	newProduct.Product_Price = c.Request.FormValue("Product_Price")
	newProduct.Product_Description = c.Request.FormValue("Product_Description")

	err := Logic.UpdateProduct(newProduct, id)
	if err != nil {
		c.HTML(400, "400", gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/Get_All_Products")
}

func Edit_Product(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	c.HTML(200, "Edit_Product", nil)
}
func Add_Product(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	c.HTML(200, "Add_Product", nil)
}
func Remove_Product(c *gin.Context) {
	if Logic.Role != "Администратор" {
		c.HTML(404, "400", gin.H{
			"Error": "Страница не найдена",
		})
		return
	}
	c.HTML(200, "DeleteProduct", nil)
}