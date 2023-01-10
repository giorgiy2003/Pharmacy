package Logic

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	Model "myapp/internal/model"
	Repository "myapp/internal/repository"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	Role    string
	User_id = 0
)

//Вывести все товары
func ReadAllProducts() ([]Model.Product, error) {
	row, err := Repository.Connection.Query(`SELECT product_id, product_image, product_name, product_price, product_manufacturer, product_category, product_description FROM "products" ORDER BY "product_id"`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Product_Image, &p.Product_Name, &p.Product_Price, &p.Product_Manufacturer, &p.Product_Category, &p.Product_Description)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		productInfo = append(productInfo, p)
	}
	return productInfo, nil
}

//Вывести первые 6 записей
func ReadProductsWithLimit() ([]Model.Product, error) {
	row, err := Repository.Connection.Query(`SELECT product_id, product_image, product_name, product_price FROM "products" ORDER BY "product_id" LIMIT 6`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Product_Image, &p.Product_Name, &p.Product_Price)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		productInfo = append(productInfo, p)
	}
	return productInfo, nil
}

//Популярные товары в течение 7 дней
func Popular_Products() ([]Model.Product, error) {
	row, err := Repository.Connection.Query(`
	SELECT  products.product_id, products.product_image, products.product_name, products.product_price
	FROM orders JOIN "products" on orders.product_id = products.product_id
	WHERE  order_time >= now() - interval  '7 days'
	GROUP BY products.product_id, orders.product_price 
	ORDER BY SUM(orders.product_koll) DESC
	LIMIT 10
	`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Product_Image, &p.Product_Name, &p.Product_Price)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		productInfo = append(productInfo, p)
	}
	return productInfo, nil
}

//Поиск товара по наименованию
func ReadOneProductByName(product_name string) ([]Model.Product, error) {
	row, err := Repository.Connection.Query(`SELECT * FROM "products" WHERE "product_name" = $1`, product_name)
	if err != nil {
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Product_Image, &p.Product_Name, &p.Product_Manufacturer, &p.Product_Category, &p.Product_Description, &p.Product_Price)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		productInfo = append(productInfo, p)
	}
	return productInfo, nil
}

//Поиск товара по id
func ReadOneProductById(product_id string) ([]Model.Product, error) {
	id, err := strconv.Atoi(product_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	row, err := Repository.Connection.Query(`SELECT product_id, product_image, product_name, product_manufacturer, product_category, product_description, product_price FROM "products" WHERE "product_id" = $1`, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Product_Image, &p.Product_Name, &p.Product_Manufacturer, &p.Product_Category, &p.Product_Description, &p.Product_Price)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		productInfo = append(productInfo, p)
	}
	return productInfo, nil
}

//Поиск товара по ID или наименованию
func SearhProduct(product_name string) ([]Model.Product, error) {
	product_name = strings.TrimSpace(product_name)
	product_name = strings.ToLower(product_name)
	product_name = strings.Title(product_name)

	Products, _ := ReadOneProductById(product_name)
	if len(Products) == 0 {
		Products, err := ReadOneProductByName(product_name)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return Products, nil
	}
	return Products, nil
}

//Выборка по категориям
func Medicines_by_category(category string) ([]Model.Product, error) {

	row, err := Repository.Connection.Query(`SELECT * FROM "products" WHERE "product_category" = $1`, category)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Product_Image, &p.Product_Name, &p.Product_Manufacturer, &p.Product_Category, &p.Product_Description, &p.Product_Price)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		productInfo = append(productInfo, p)
	}
	return productInfo, nil
}

//Сортировать товары по названию
func NameASC() ([]Model.Product, error) {
	row, err := Repository.Connection.Query(`SELECT * FROM "products" ORDER BY "product_name"`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Product_Image, &p.Product_Name, &p.Product_Manufacturer, &p.Product_Category, &p.Product_Description, &p.Product_Price)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		productInfo = append(productInfo, p)
	}
	return productInfo, nil
}

//Сортировать товары по названию в обратном порядке
func NameDESC() ([]Model.Product, error) {
	row, err := Repository.Connection.Query(`SELECT * FROM "products" ORDER BY "product_name" DESC`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Product_Image, &p.Product_Name, &p.Product_Manufacturer, &p.Product_Category, &p.Product_Description, &p.Product_Price)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		productInfo = append(productInfo, p)
	}
	return productInfo, nil
}

//Сортировать товары по цене
func PriceASC() ([]Model.Product, error) {
	row, err := Repository.Connection.Query(`SELECT * FROM "products" ORDER BY "product_price"`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Product_Image, &p.Product_Name, &p.Product_Manufacturer, &p.Product_Category, &p.Product_Description, &p.Product_Price)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		productInfo = append(productInfo, p)
	}
	return productInfo, nil
}

//Сортировать товары по цене в обратном порядке
func PriceDESC() ([]Model.Product, error) {
	row, err := Repository.Connection.Query(`SELECT * FROM "products" ORDER BY "product_price" DESC `)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Product_Image, &p.Product_Name, &p.Product_Manufacturer, &p.Product_Category, &p.Product_Description, &p.Product_Price)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		productInfo = append(productInfo, p)
	}
	return productInfo, nil
}

//Авторизация
func Autorization(login, password string) error {
	Login := strings.TrimSpace(login)

	//Hashed value of password
	Password := SHA_256_Encode(password) //Кодируем значение пароля пользователя

	row, err := Repository.Connection.Query(`SELECT "user_id", "user_role" FROM "users" WHERE user_login = $1 AND user_password = $2`, Login, Password)
	if err != nil {
		log.Println(err)
		return err
	}

	for row.Next() {
		row.Scan(&User_id, &Role)
	}

	//Если значения структуры пусты возращаем ошибку
	if User_id == 0 {
		return errors.New("Введён неверный логин или пароль!")
	}
	return nil
}

//Кодирование данных
func SHA_256_Encode(text string) string {
	algorithm := sha256.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

//Регистрация
func Registration(UserName, UserEmail, UserPassword1, UserPassword2, Checkbox string) error {
	UserName = strings.TrimSpace(UserName)
	UserEmail = strings.TrimSpace(UserEmail)

	if UserName == "" || UserEmail == "" || UserPassword1 == "" || UserPassword2 == "" {
		return errors.New("Ошибка: не все поля заполнены!")
	}

	if Checkbox != "true" {
		return errors.New("Для регистрации необходимо принять условия пользовательского соглашения!")
	}

	row, err := Repository.Connection.Query(`SELECT user_login FROM "users" WHERE user_login = $1`, UserEmail)
	if err != nil {
		log.Println(err)
		return err
	}

	var u Model.User
	for row.Next() {
		err := row.Scan(&u.Login)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	//Если значения структуры не пусты возращаем ошибку
	if u.Login != "" {
		return errors.New("Пользователь с введённым Email уже зарегистрирован!")
	}

	if UserPassword1 != UserPassword2 {
		return errors.New("Ошибка: пароли не совпадают!")
	}

	if len(UserPassword1) < 4 {
		return errors.New("Слишком короткий пароль, минимальный размер 4 символа!")
	}

	//Hashed value of password
	UserPassword1 = SHA_256_Encode(UserPassword1) //Кодируем значение пароля пользователя

	_, err = Repository.Connection.Exec(`INSERT INTO "users" ("user_name","user_login", "user_password", "user_role" ) VALUES ($1,$2,$3,$4)`, UserName, UserEmail, UserPassword1, "Пользователь")
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//Корзина
func UserCart() ([]Model.UserCart, int, error) {

	if User_id == 0 {
		return nil, 0, nil
	}

	row, err := Repository.Connection.Query(`
	SELECT products.product_id, products.product_image, products.product_name, products.product_manufacturer, products.product_category, products.product_description, 
	products.product_price, shopping_cart.product_koll, shopping_cart.product_koll * products.product_price AS product_amount
	FROM products JOIN "shopping_cart" on products.product_id = shopping_cart.product_id
	WHERE user_id = $1
	GROUP BY products.product_id, shopping_cart.product_koll, time_of_adding
	ORDER BY "time_of_adding" DESC
	`, User_id)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	var UserInfo = []Model.UserCart{}
	for row.Next() {
		var UserCart Model.UserCart
		err := row.Scan(&UserCart.Product_Id, &UserCart.Product_Image, &UserCart.Product_Name, &UserCart.Product_Manufacturer, &UserCart.Product_Category, &UserCart.Product_Description, &UserCart.Product_Price, &UserCart.Product_Koll, &UserCart.Product_amount)
		if err != nil {
			log.Println(err)
			return nil, 0, err
		}
		UserInfo = append(UserInfo, UserCart)
	}

	var total sql.NullInt64

	if UserInfo != nil {

		row, err = Repository.Connection.Query(`
		SELECT SUM(shopping_cart.product_koll * products.product_price) AS Product_total_price
		FROM products JOIN "shopping_cart" on products.product_id = shopping_cart.product_id
		WHERE user_id =$1
		`, User_id)
		if err != nil {
			log.Println(err)
			return nil, 0, err
		}

		for row.Next() {
			err := row.Scan(&total)
			if err != nil {
				log.Println(err)
				return nil, 0, err
			}
		}
	}
	return UserInfo, int(total.Int64), nil
}

//Сделать заказ
func MakeOrder(city, fname, lname, patronymic, address, email_address, phone, order_notes string) error {

	if User_id == 0 {
		return nil
	}

	fname = strings.TrimSpace(fname)
	lname = strings.TrimSpace(lname)
	patronymic = strings.TrimSpace(patronymic)
	address = strings.TrimSpace(address)
	email_address = strings.TrimSpace(email_address)
	phone = strings.TrimSpace(phone)
	order_notes = strings.TrimSpace(order_notes)

	if fname == "" || lname == "" || address == "" || phone == "" {
		return errors.New("Ошибка: не все поля заполнены!")
	}

	name := fmt.Sprintf("%s %s %s", fname, lname, patronymic)
	newAddress := fmt.Sprintf("%s, %s", city, address)

	row, err := Repository.Connection.Query(`
	SELECT products.product_id, products.product_image, products.product_name, products.product_manufacturer, products.product_category, products.product_description, 
	products.product_price, shopping_cart.product_koll, shopping_cart.product_koll * products.product_price AS product_amount
	FROM products JOIN "shopping_cart" on products.product_id = shopping_cart.product_id
	WHERE user_id = $1
	GROUP BY products.product_id, shopping_cart.product_koll, time_of_adding
	ORDER BY "time_of_adding"
	`, User_id)
	if err != nil {
		log.Println(err)
		return err
	}

	for row.Next() {
		var UserCart Model.UserCart
		err := row.Scan(&UserCart.Product_Id, &UserCart.Product_Image, &UserCart.Product_Name, &UserCart.Product_Manufacturer, &UserCart.Product_Category, &UserCart.Product_Description, &UserCart.Product_Price, &UserCart.Product_Koll, &UserCart.Product_amount)
		if err != nil {
			log.Println(err)
			return err
		}

		rows, err := Repository.Connection.Query(`SELECT MAX (order_id) FROM "orders" WHERE user_id = $1`, User_id)
		if err != nil {
			log.Println(err)
			return err
		}
		var order_id int
		for rows.Next() {
			rows.Scan(&order_id)
		}
		Track_number := make_Track_number(order_id)

		delivery := 0

		if UserCart.Product_amount < 1000 {
			delivery = 150
		}
		total := delivery + UserCart.Product_amount

		if _, err := Repository.Connection.Exec(`INSERT INTO "orders" ("user_id", "product_id", "product_koll", "product_price", "order_time", 
		"order_status", "order_track_number", "delivery_price", "total_price", "customer_name", "customer_address", "customer_phone", "customer_email", "customer_comment") 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)`,
			User_id, UserCart.Product_Id, UserCart.Product_Koll, UserCart.Product_Price, time.Now(), "Ожидает подтверждения", Track_number, delivery, total, name, newAddress, phone, email_address, order_notes); err != nil {
			log.Println(err)
			return err
		}

		//После добавления товара в таблицу заказы, удаляем его из корзины
		if _, err := Repository.Connection.Exec(`DELETE FROM "shopping_cart" WHERE user_id = $1 AND product_id = $2`, User_id, UserCart.Product_Id); err != nil {
			log.Println(err)
			return err
		}

	}
	return nil
}

//Создание трек-номера заказа
func make_Track_number(Order_id int) string {
	return fmt.Sprintf("%d-%d", User_id+1000000, Order_id+1)
}

//Максимальный product_id из таблицы product
func ProductMax() (int, error) {
	rows, err := Repository.Connection.Query(`SELECT MAX (product_id) FROM "products"`)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	var product_id int
	for rows.Next() {
		rows.Scan(&product_id)
	}
	return product_id, nil
}

//Информация о заказе по трек-номеру
func Order_details(Track_number string) ([]Model.Order, error) {

	if User_id == 0 {
		return nil, nil
	}

	row, err := Repository.Connection.Query(`
	SELECT orders.product_id, products.product_name, "product_koll", orders.product_price, "order_time", "order_status", orders.product_koll * orders.product_price AS product_amount, "order_track_number",
	"delivery_price", "total_price", "customer_name", "customer_address", "customer_phone", "customer_email"
	FROM "orders" JOIN "products" on products.product_id = orders.product_id
	WHERE user_id = $1 AND order_Track_number = $2
	`, User_id, Track_number)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var OrderInfo = []Model.Order{}
	for row.Next() {
		var (
			Order Model.Order
			time  time.Time
		)
		err := row.Scan(&Order.Product_Id, &Order.Product_Name, &Order.Product_Koll, &Order.Product_Price, &time, &Order.Order_status, &Order.Product_amount, &Order.Track_number, &Order.Delivery_price, &Order.Total_price, &Order.Customer_Name, &Order.Customer_Address, &Order.Customer_Phone, &Order.Customer_Email)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		Order.Order_time = time.Format("2006-01-02")
		OrderInfo = append(OrderInfo, Order)
	}
	return OrderInfo, nil
}

//Доставки
func Orders() ([]Model.Order, error) {

	if User_id == 0 {
		return nil, nil
	}

	row, err := Repository.Connection.Query(`
	SELECT products.product_id, products.product_image, products.product_name, orders.product_price, orders.product_koll, orders.order_time, orders.order_status, orders.product_koll * orders.product_price AS product_amount, orders.order_track_number
	FROM products JOIN "orders" on products.product_id = orders.product_id
	WHERE user_id = $1
	GROUP BY products.product_id, orders.product_koll, order_time, orders.product_price, orders.order_status, orders.order_track_number
	ORDER BY "order_time" DESC
	`, User_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var OrderInfo = []Model.Order{}
	for row.Next() {
		var (
			Order Model.Order
			time  time.Time
		)
		err := row.Scan(&Order.Product_Id, &Order.Product_Image, &Order.Product_Name, &Order.Product_Price, &Order.Product_Koll, &time, &Order.Order_status, &Order.Product_amount, &Order.Track_number)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		Order.Order_time = time.Format("2006-01-02")
		OrderInfo = append(OrderInfo, Order)
	}
	return OrderInfo, nil
}

//Добавить в корзину
func AddToCart(id string) error {

	if User_id == 0 {
		return nil
	}

	product_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return err
	}

	//Проверяем существует товар в корзине пользователя
	rows, err := Repository.Connection.Query(`SELECT "product_id" FROM "shopping_cart" WHERE user_id = $1 AND product_id = $2`, User_id, product_id)
	if err != nil {
		log.Println(err)
		return err
	}

	//Ecли товар уже в корзине выходим из функции
	for rows.Next() {
		var UserCart Model.UserCart
		rows.Scan(&UserCart.Product_Id)
		if UserCart.Product_Id != 0 {
			return nil
		}
	}
	//Ecли товара не было в корзине, добавляем его
	if _, err := Repository.Connection.Exec(`INSERT INTO "shopping_cart" ("user_id","product_id","product_koll", "time_of_adding") VALUES ($1,$2,$3,$4)`, User_id, product_id, 1, time.Now()); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//Убрать из корзины
func DeleteFromCart(id string) error {

	if User_id == 0 {
		return nil
	}

	product_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return err
	}
	if _, err := Repository.Connection.Exec(`DELETE FROM "shopping_cart" WHERE user_id = $1 AND product_id = $2`, User_id, product_id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//Проверка на наличие товара в корзине пользователя
func Proverka1(id string) (string, error) {

	product_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return "", err
	}
	rows, err := Repository.Connection.Query(`SELECT "product_id" FROM "shopping_cart" WHERE user_id = $1 AND product_id = $2`, User_id, product_id)
	if err != nil {
		log.Println(err)
		return "", err
	}

	for rows.Next() {
		var UserCart Model.UserCart
		rows.Scan(&UserCart.Product_Id)
		if UserCart.Product_Id != 0 {
			return "товар в корзине", nil
		}
	}
	return "", nil
}

//Проверка на наличие товара в избранном у пользователя
func Proverka2(id string) (string, error) {

	product_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return "", err
	}
	rows, err := Repository.Connection.Query(`SELECT "product_id" FROM "favourites" WHERE user_id = $1 AND product_id = $2`, User_id, product_id)
	if err != nil {
		log.Println(err)
		return "", err
	}

	for rows.Next() {
		var UserCart Model.UserCart
		rows.Scan(&UserCart.Product_Id)
		if UserCart.Product_Id != 0 {
			return "товар в избранном", nil
		}
	}
	return "", nil
}

//Товары в избранном
func Favourites() ([]Model.UserCart, error) {

	if User_id == 0 {
		return nil, nil
	}

	row, err := Repository.Connection.Query(`
	SELECT products.product_id, products.product_image, products.product_name, products.product_price
	FROM products JOIN "favourites" on products.product_id = favourites.product_id
	WHERE user_id = $1
	GROUP BY products.product_id, time_of_adding
	ORDER BY "time_of_adding" DESC
	`, User_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var UserInfo = []Model.UserCart{}
	for row.Next() {
		var UserCart Model.UserCart
		err := row.Scan(&UserCart.Product_Id, &UserCart.Product_Image, &UserCart.Product_Name, &UserCart.Product_Price)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		UserInfo = append(UserInfo, UserCart)
	}
	return UserInfo, nil
}

//Добавить в избранное
func AddToFavotites(id string) error {

	if User_id == 0 {
		return nil
	}

	product_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return err
	}

	//Проверяем существует ли товар в избранном
	rows, err := Repository.Connection.Query(`SELECT "product_id" FROM "favourites" WHERE user_id = $1 AND product_id = $2`, User_id, product_id)
	if err != nil {
		log.Println(err)
		return err
	}

	//Ecли товар уже в избранном выходим из функции
	for rows.Next() {
		var UserCart Model.UserCart
		rows.Scan(&UserCart.Product_Id)
		if UserCart.Product_Id != 0 {
			return nil
		}
	}
	//Ecли товара не было в избранном, добавляем его
	if _, err := Repository.Connection.Exec(`INSERT INTO "favourites" ("user_id","product_id","time_of_adding") VALUES ($1,$2,$3)`, User_id, product_id, time.Now()); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//Убрать из избранного
func DeleteFromFavotites(id string) error {

	if User_id == 0 {
		return nil
	}

	product_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return err
	}
	if _, err := Repository.Connection.Exec(`DELETE FROM "favourites" WHERE user_id = $1 AND product_id = $2`, User_id, product_id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//Увеличить количество товара в корзине
func AddKoll(id, koll string) error {

	if User_id == 0 {
		return nil
	}

	product_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return err
	}
	product_koll, err := strconv.Atoi(koll)
	if err != nil {
		log.Println(err)
		return err
	}
	product_koll++
	//Проверяем существует ли товар в корзине пользователя
	rows, err := Repository.Connection.Query(`SELECT "product_id" FROM "shopping_cart" WHERE user_id = $1 AND product_id = $2`, User_id, product_id)
	if err != nil {
		log.Println(err)
		return err
	}
	//Ecли товар в корзине обновляем количество товара
	for rows.Next() {
		var UserCart Model.UserCart
		rows.Scan(&UserCart.Product_Id)
		if UserCart.Product_Id != 0 {
			if _, err := Repository.Connection.Exec(`UPDATE "shopping_cart" SET "product_koll" = $1 WHERE user_id = $2 AND product_id = $3`, product_koll, User_id, product_id); err != nil {
				log.Println(err)
				return err
			}
			return nil
		}
	}
	return nil
}

//Уменьшить количество товара в корзине
func MinusKoll(id, koll string) error {

	if User_id == 0 {
		return nil
	}
	product_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return err
	}
	product_koll, err := strconv.Atoi(koll)
	if err != nil {
		log.Println(err)
		return err
	}
	product_koll--
	if product_koll == 0 {
		err := DeleteFromCart(id)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	//Проверяем существует ли товар в корзине пользователя
	rows, err := Repository.Connection.Query(`SELECT "product_id" FROM "shopping_cart" WHERE user_id = $1 AND product_id = $2`, User_id, product_id)
	if err != nil {
		log.Println(err)
		return err
	}

	//Ecли товар в корзине обновляем количество товара
	for rows.Next() {
		var UserCart Model.UserCart
		rows.Scan(&UserCart.Product_Id)
		if UserCart.Product_Id != 0 {
			if _, err := Repository.Connection.Exec(`UPDATE "shopping_cart" SET "product_koll" = $1 WHERE user_id = $2 AND product_id = $3`, product_koll, User_id, product_id); err != nil {
				log.Println(err)
				return err
			}
			return nil
		}
	}
	return nil
}

//Просмотреть карточку товара
func ShopSingle(id string) ([]Model.UserCart, error) {

	var UserInfo = []Model.UserCart{}
	var UserCart Model.UserCart

	product_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if User_id != 0 {
		row, err := Repository.Connection.Query(`
		SELECT products.product_id, products.product_image, products.product_name, products.product_manufacturer, products.product_category, products.product_description, products.product_price, shopping_cart.product_koll 
		FROM products JOIN "shopping_cart" on products.product_id = shopping_cart.product_id
		WHERE user_id = $1 AND products.product_id = $2
		`, User_id, product_id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		for row.Next() {
			err := row.Scan(&UserCart.Product_Id, &UserCart.Product_Image, &UserCart.Product_Name, &UserCart.Product_Manufacturer, &UserCart.Product_Category, &UserCart.Product_Description, &UserCart.Product_Price, &UserCart.Product_Koll)
			if err != nil {
				log.Println(err)
				return nil, err
			}
		}
		UserInfo = append(UserInfo, UserCart)
	}

	if User_id == 0 || UserCart.Product_Koll == 0 {
		row, err := Repository.Connection.Query(`SELECT * FROM "products" WHERE "product_id" = $1`, product_id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		UserInfo = nil
		for row.Next() {
			err := row.Scan(&UserCart.Product_Id, &UserCart.Product_Image, &UserCart.Product_Name, &UserCart.Product_Manufacturer, &UserCart.Product_Category, &UserCart.Product_Description, &UserCart.Product_Price)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			UserInfo = append(UserInfo, UserCart)
		}
	}
	return UserInfo, nil
}

//Для администратора

//Заказы

//Список заказов
func Orders_Page(order_status string) ([]Model.Order, error) {

	if User_id == 0 {
		return nil, nil
	}

	row, err := Repository.Connection.Query(`
	SELECT orders.user_id, products.product_id,  products.product_name, orders.product_price, orders.product_koll, orders.order_time, orders.order_status, orders.product_koll * orders.product_price AS product_amount, 
	orders.order_track_number, orders.order_id ,orders.customer_address, orders.customer_phone, orders.customer_email,orders.customer_comment
	FROM products JOIN "orders" on products.product_id = orders.product_id
	WHERE orders.order_status = $1
	GROUP BY orders.user_id, products.product_id, orders.product_koll, order_time, orders.product_price, orders.order_status, orders.order_track_number, orders.order_id, orders.customer_address, orders.customer_phone, orders.customer_email,orders.customer_comment
	ORDER BY "order_time" ASC, orders.user_id
	`, order_status)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var OrderInfo = []Model.Order{}
	for row.Next() {
		var (
			Order Model.Order
			time  time.Time
		)
		err := row.Scan(&Order.User_Id, &Order.Product_Id, &Order.Product_Name, &Order.Product_Price, &Order.Product_Koll, &time, &Order.Order_status, &Order.Product_amount,
			&Order.Track_number, &Order.Order_Id, &Order.Customer_Address, &Order.Customer_Phone, &Order.Customer_Email, &Order.Customer_Comment)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		Order.Order_time = time.Format("2006-01-02")
		OrderInfo = append(OrderInfo, Order)
	}
	return OrderInfo, nil
}

//Поменять статус заказа
func Change_status(order_status, order_id string) error {

	if User_id == 0 {
		return nil
	}

	id, err := strconv.Atoi(order_id)
	if err != nil {
		log.Println(err)
		return err
	}

	if _, err := Repository.Connection.Exec(`UPDATE "orders" SET "order_status" = $1 WHERE "order_id" = $2`, order_status, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//Сотрудники

//Вывести всех сотрудников
func ReadAllWorkers() ([]Model.Worker, error) {
	row, err := Repository.Connection.Query(`SELECT worker_id, worker_firstname, worker_lastname, worker_email, worker_phone, post, salary_per_month FROM "workers" ORDER BY "worker_id"`)
	if err != nil {
		return nil, err
	}
	var personInfo = []Model.Worker{}
	for row.Next() {
		var p Model.Worker
		err := row.Scan(&p.Worker_Id, &p.Worker_FirstName, &p.Worker_LastName, &p.Worker_Email, &p.Worker_Phone, &p.Post, &p.Salary_per_month)
		if err != nil {
			return nil, err
		}
		personInfo = append(personInfo, p)
	}
	return personInfo, nil
}

//Добавить сотрудника
func CreateWorker(p Model.Worker) error {
	p.Worker_Email = strings.TrimSpace(p.Worker_Email)
	p.Worker_Phone = strings.TrimSpace(p.Worker_Phone)
	p.Worker_FirstName = strings.TrimSpace(p.Worker_FirstName)
	p.Worker_LastName = strings.TrimSpace(p.Worker_LastName)
	p.Post = strings.TrimSpace(p.Post)
	p.Salary_per_month = strings.TrimSpace(p.Salary_per_month)

	if p.Worker_FirstName == "" || p.Worker_LastName == "" || p.Worker_Phone == "" || p.Post == "" || p.Salary_per_month == "" {
		return errors.New("невозможно добавить запись, не все поля заполнены!")
	}

	Salary_per_month, err := strconv.Atoi(p.Salary_per_month)
	if err != nil {
		log.Println(err)
		return err
	}

	if _, err := Repository.Connection.Exec(`INSERT INTO "workers" ("worker_firstname", "worker_lastname", "worker_email", "worker_phone", "post", "salary_per_month", "time_of_adding") VALUES ($1,$2,$3,$4,$5,$6,$7)`, p.Worker_FirstName, p.Worker_LastName, p.Worker_Email, p.Worker_Phone, p.Post, Salary_per_month, time.Now()); err != nil {
		return err
	}
	return nil
}

//Редактировать запись сотрудника
func UpdateWorker(p Model.Worker, id string) error {

	if err := workerExist(id); err != nil {
		return err
	}
	p.Worker_Email = strings.TrimSpace(p.Worker_Email)
	p.Worker_Phone = strings.TrimSpace(p.Worker_Phone)
	p.Worker_FirstName = strings.TrimSpace(p.Worker_FirstName)
	p.Worker_LastName = strings.TrimSpace(p.Worker_LastName)
	p.Post = strings.TrimSpace(p.Post)
	p.Salary_per_month = strings.TrimSpace(p.Salary_per_month)
	if p.Worker_FirstName == "" || p.Worker_LastName == "" || p.Worker_Phone == "" || p.Post == "" || p.Salary_per_month == "" {
		return errors.New("невозможно редактировать запись, не все поля заполнены!")
	}
	if _, err := Repository.Connection.Exec(`UPDATE "workers" SET "worker_firstname" = $1,"worker_lastname" = $2,"worker_email" = $3,"worker_phone" = $4 ,"post" = $5 ,"salary_per_month" = $6  WHERE "worker_id" = $7`, p.Worker_FirstName, p.Worker_LastName, p.Worker_Email, p.Worker_Phone, p.Post, p.Salary_per_month, id); err != nil {
		return err
	}
	return nil
}

//Удалить сотрудника из базы
func DeleteWorker(id string) error {
	if err := workerExist(id); err != nil {
		return err
	}
	if _, err := Repository.Connection.Exec(`DELETE FROM "workers" WHERE "worker_id" = $1`, id); err != nil {
		return err
	}
	return nil
}

//Найти сотрудника по id
func ReadOneWorker(id string) ([]Model.Worker, error) {
	worker_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("Error: неверно введён параметр id")
	}
	row, err := Repository.Connection.Query(`SELECT worker_id, worker_firstname, worker_lastname, worker_email, worker_phone, post, salary_per_month FROM "workers" WHERE "worker_id" = $1`, worker_id)
	if err != nil {
		return nil, err
	}
	var personInfo = []Model.Worker{}
	for row.Next() {
		var p Model.Worker
		err := row.Scan(&p.Worker_Id, &p.Worker_FirstName, &p.Worker_LastName, &p.Worker_Email, &p.Worker_Phone, &p.Post, &p.Salary_per_month)
		if err != nil {
			return nil, err
		}
		personInfo = append(personInfo, p)
	}
	return personInfo, nil
}

//Проверка на наличие сотрудника в базе
func workerExist(id string) error {
	persons, err := ReadOneWorker(id)
	if err != nil {
		return err
	}
	if len(persons) == 0 {
		return fmt.Errorf("Error: записи с id = %s не существует", id)
	}
	return nil
}

//Отзывы

//Оставить отзыв
func CreateComment(p Model.Comment) error {

	p.Customer_FirstName = strings.TrimSpace(p.Customer_FirstName)
	p.Customer_LastName = strings.TrimSpace(p.Customer_LastName)
	p.Customer_Email = strings.TrimSpace(p.Customer_Email)
	p.Theme = strings.TrimSpace(p.Theme)
	p.Comment = strings.TrimSpace(p.Comment)

	if p.Customer_FirstName == "" || p.Customer_LastName == "" || p.Customer_Email == "" || p.Comment == "" {
		return errors.New("невозможно добавить запись, не все поля заполнены!")
	}

	if _, err := Repository.Connection.Exec(`INSERT INTO "comments" ("user_id", "customer_firstname", "customer_lastname", "customer_email", "theme", "message", "time_of_adding") VALUES ($1,$2,$3,$4,$5,$6,$7)`, User_id, p.Customer_FirstName, p.Customer_LastName, p.Customer_Email, p.Theme, p.Comment, time.Now()); err != nil {
		return err
	}
	return nil
}

//Вывести отзывы
func ReadAllComments() ([]Model.Comment, error) {
	row, err := Repository.Connection.Query(`SELECT "comment_id", "user_id", "customer_firstname", "customer_lastname", "customer_email", "theme", "message", "comment_status" FROM "comments" WHERE "comment_status" IS NULL ORDER BY "time_of_adding"`)
	if err != nil {
		return nil, err
	}
	var CommentInfo = []Model.Comment{}
	for row.Next() {
		var p Model.Comment
		var Comment_status sql.NullString
		err := row.Scan(&p.Comment_Id, &p.User_Id, &p.Customer_FirstName, &p.Customer_LastName, &p.Customer_Email, &p.Theme, &p.Comment, &Comment_status)
		if err != nil {
			return nil, err
		}
		if Comment_status.Valid {
			p.Comment_status = Comment_status.String
		} else {
			p.Comment_status = ""
		}
		CommentInfo = append(CommentInfo, p)
	}
	return CommentInfo, nil
}

//Вывести отзывы со статусом "Важно"
func ReadImportantComments() ([]Model.Comment, error) {
	row, err := Repository.Connection.Query(`SELECT "comment_id", "user_id", "customer_firstname", "customer_lastname", "customer_email", "theme", "message", "comment_status" FROM "comments" WHERE "comment_status"= $1 ORDER BY "time_of_adding"`, "Важно")
	if err != nil {
		return nil, err
	}
	var CommentInfo = []Model.Comment{}
	for row.Next() {
		var p Model.Comment
		err := row.Scan(&p.Comment_Id, &p.User_Id, &p.Customer_FirstName, &p.Customer_LastName, &p.Customer_Email, &p.Theme, &p.Comment, &p.Comment_status)
		if err != nil {
			return nil, err
		}
		CommentInfo = append(CommentInfo, p)
	}
	return CommentInfo, nil
}

//Именить статус товара на "Важно"
func Change_To_Important(id string) error {
	Comment_id, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("Error: неверно введён параметр id")
	}
	if _, err := Repository.Connection.Exec(`UPDATE "comments" SET "comment_status" = $1  WHERE "comment_id" = $2`, "Важно", Comment_id); err != nil {
		return err
	}
	return nil
}

//Удалить комментарий
func DeleteComment(id string) error {
	Comment_id, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("Error: неверно введён параметр id")
	}
	if _, err := Repository.Connection.Exec(`DELETE FROM "comments" WHERE "comment_id" = $1`, Comment_id); err != nil {
		return err
	}
	return nil
}

//Товары

//Добавить товар
func CreateProduct(p Model.Product) error {
	p.Product_Name = strings.TrimSpace(p.Product_Name)
	p.Product_Image = strings.TrimSpace(p.Product_Image)
	p.Product_Manufacturer = strings.TrimSpace(p.Product_Manufacturer)
	p.Product_Category = strings.TrimSpace(p.Product_Category)
	p.Product_Description = strings.TrimSpace(p.Product_Description)
	p.Product_Price = strings.TrimSpace(p.Product_Price)
	if p.Product_Name == "" || p.Product_Image == "" || p.Product_Manufacturer == "" || p.Product_Category == "" || p.Product_Description == "" || p.Product_Price == "" {
		return errors.New("невозможно добавить запись, не все поля заполнены!")
	}
	if _, err := Repository.Connection.Exec(`INSERT INTO "products" ("product_name","product_image", "product_manufacturer", "product_category", "product_description","product_price" ) VALUES ($1,$2,$3,$4,$5,$6)`, p.Product_Name, p.Product_Image, p.Product_Manufacturer, p.Product_Category, p.Product_Description, p.Product_Price); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//Редактировать карточку товара
func UpdateProduct(p Model.Product, id string) error {

	products, err := ProductExist(id)
	if err != nil {
		log.Println(err)
		return err
	}

	p.Product_Name = strings.TrimSpace(p.Product_Name)
	p.Product_Image = strings.TrimSpace(p.Product_Image)
	p.Product_Manufacturer = strings.TrimSpace(p.Product_Manufacturer)
	p.Product_Category = strings.TrimSpace(p.Product_Category)
	p.Product_Description = strings.TrimSpace(p.Product_Description)
	p.Product_Price = strings.TrimSpace(p.Product_Price)
	if p.Product_Name == "" || p.Product_Image == "" || p.Product_Manufacturer == "" || p.Product_Category == "" || p.Product_Description == "" || p.Product_Price == "" {
		return errors.New("невозможно редактировать запись, не все поля заполнены!")
	}

	for _, p := range products {
		err := os.Remove(fmt.Sprint("./Frontend/images/products/", p.Product_Image))
		if err != nil {
			log.Println(err)
		}
	}

	if _, err := Repository.Connection.Exec(`UPDATE "products" SET "product_name" = $1,"product_image" = $2,"product_manufacturer" = $3,"product_category" = $4,"product_description" = $5, "product_price" = $6  WHERE "product_id" = $7`, p.Product_Name, p.Product_Image, p.Product_Manufacturer, p.Product_Category, p.Product_Description, p.Product_Price, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//Удалить товар
func Form_handler_DeleteProductById(id string) error {
	products, err := ProductExist(id)
	if err != nil {
		log.Println(err)
		return err
	}
	for _, p := range products {
		err := os.Remove(fmt.Sprint("./Frontend/images/products/", p.Product_Image))
		if err != nil {
			log.Println(err)
		}
	}
	if _, err := Repository.Connection.Exec(`DELETE FROM "products" WHERE "product_id" = $1`, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ProductExist(id string) ([]Model.Product, error) {
	products, err := ReadOneProductById(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(products) == 0 {
		return nil, fmt.Errorf("Товара с id = %s не существует", id)
	}
	return products, nil
}
