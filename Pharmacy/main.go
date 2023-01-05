package main

import (
	Handler "myapp/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/css", "Frontend/css")
	router.Static("/fonts", "Frontend/fonts")
	router.Static("/images", "Frontend/images")
	router.Static("/js", "Frontend/js")
	router.LoadHTMLGlob("Frontend/*.html")
	router.Use(Handler.ConnectDB())
	router.GET("/Authorization", Handler.Authorization) //Авторизация ✅ 
	router.GET("/Form_handler_Authorization", Handler.Form_handler_Authorization) //Обработчик авторизации ✅
	router.GET("/Registration", Handler.Registration) //Регистрация пользователя ✅
	router.GET("/Form_handler_Registration", Handler.Form_handler_Registration) //Обработчик регистрации ✅
	router.GET("/Sign_out", Handler.Sign_out) //Выйти из аккаунта ✅
	router.GET("/", Handler.MainForm) //Главная страница ✅
	router.GET("/shop", Handler.Shop) //Магазин ✅
	router.GET("/NameASC", Handler.NameASC) //Фильтр товаров по наименованию ✅
	router.GET("/NameDESC", Handler.NameDESC) //Фильтр товаров по наименованию в обратном порядке ✅
	router.GET("/PriceASC", Handler.PriceASC) //Фильтр товаров по цене ✅
	router.GET("/PriceDESC", Handler.PriceDESC) //Фильтр товаров по цене в обратном порядке ✅
	router.GET("/shop_single", Handler.Shop_single) //Просмотр карточки товара ✅
	router.GET("/SearhProduct", Handler.SearhProduct) //Поиск товара по ID или названию ✅
	router.GET("/cart", Handler.Cart) //Корзина ✅
	router.GET("/about", Handler.About) //О нас ✅
	router.GET("/contact", Handler.Contact) //Контакты ✅
	router.GET("/SendMessage", Handler.SendMessage)	//Оставить отзыв 
	router.GET("/AddToCart/:id", Handler.AddToCart) //Добавить в корзину ✅
	router.GET("/DeleteFromCart/:id", Handler.DeleteFromCart) //Убрать из корзины ✅
	router.GET("/MinusKoll/:id", Handler.MinusKoll) //Уменьшить количество товара в корзине на странице товара ✅
	router.GET("/AddKoll/:id", Handler.AddKoll) //Уменьшить количество товара в корзине на странице товара ✅
	router.GET("/MinusKollinCart/:id", Handler.MinusKollinCart) //Уменьшить количество товара в корзине ✅
	router.GET("/AddKollinCart/:id", Handler.AddKollinCart) //Уменьшить количество товара в корзине ✅
	router.GET("/checkout", Handler.Checkout) //Страница оформления заказа ✅
	router.GET("/Order", Handler.Order) //Оформить заказ ✅
	router.GET("/favouritesPage", Handler.Favourites) //Товары в избранном ✅
	router.GET("/AddTofavourites/:id", Handler.AddToFavotites) //Добавить в избранное ✅
	router.GET("/DeleteFromfavourites/:id", Handler.DeleteFromFavotites) //Убрать из избранного ✅
	router.GET("/historyPage", Handler.HistoryPage) //История заказов ✅
	router.GET("/order_details/:order", Handler.Order_details) //Информация о заказе ✅
	router.GET("/Product_category/:category", Handler.Medicines_by_category) //Поиск по категориям ✅

	//Для администраторов

	//Заказы
	router.GET("/orders/:order_status", Handler.Orders_Page) //Вывести заказы ✅
	router.GET("/change_status/:order_id/:order_status", Handler.Change_status) //Поменять статус заказа ✅

	//Сотрудники
	router.GET("/Get_All_Workers", Handler.Get_All_Workers) //Вывести всех
	router.GET("/Add_Worker", Handler.Add_Worker) //Добавить сотрудника
	router.GET("/Remove_Worker", Handler.Remove_Worker) //Удалить запись
	router.GET("/Edit_Worker", Handler.Edit_Worker) //Редактировать запись сотрудника
	router.POST("/Form_handler_PostWorker", Handler.Form_handler_PostWorker)
	router.GET("/Form_handler_UpdateWorkerById", Handler.Form_handler_UpdateWorkerById)
	router.GET("/Form_handler_DeleteWorkerById", Handler.Form_handler_DeleteWorkerById)
	router.Run("localhost:8080")
}
