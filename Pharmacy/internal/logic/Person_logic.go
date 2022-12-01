package Logic

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	Model "myapp/internal/model"
	Repository "myapp/internal/repository"
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
	row, err := Repository.Connection.Query(`SELECT * FROM "products" ORDER BY "product_id"`)
	if err != nil {
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
		if err != nil {
			return nil, err
		}
		productInfo = append(productInfo, p)
	}
	return productInfo, nil
}

//Вывести первые 6 записей
func ReadProductsWithLimit() ([]Model.Product, error) {
	row, err := Repository.Connection.Query(`SELECT * FROM "products" ORDER BY "product_id" LIMIT 6`)
	if err != nil {
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
		if err != nil {
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
		err := row.Scan(&p.Product_Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
		if err != nil {
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
		return nil, err
	}
	row, err := Repository.Connection.Query(`SELECT * FROM "products" WHERE "product_id" = $1`, id)
	if err != nil {
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
		if err != nil {
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
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
		if err != nil {
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
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
		if err != nil {
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
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
		if err != nil {
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
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
		if err != nil {
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
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Product_Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
		if err != nil {
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
		return err
	}

	var u Model.User
	for row.Next() {
		err := row.Scan(&u.Login)
		if err != nil {
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
		return err
	}
	return nil
}

//Корзина
func UserCart() ([]Model.UserCart, error) {

	if User_id == 0 {
		return nil, nil
	}

	row, err := Repository.Connection.Query(`
	SELECT products.product_id, products.product_image, products.product_name, products.product_manufacturer, products.product_category, products.product_description, products.product_price, shopping_cart.product_koll 
	FROM products JOIN "shopping_cart" on products.product_id = shopping_cart.product_id
	WHERE user_id = $1
	ORDER BY "time_of_adding" DESC
	`, User_id)
	if err != nil {
		return nil, err
	}

	var UserInfo = []Model.UserCart{}
	for row.Next() {
		var UserCart Model.UserCart
		err := row.Scan(&UserCart.Product_Id, &UserCart.Image, &UserCart.Name, &UserCart.Manufacturer, &UserCart.Category, &UserCart.Description, &UserCart.Price, &UserCart.Product_Koll)
		if err != nil {
			return nil, err
		}
		UserInfo = append(UserInfo, UserCart)
	}

	return UserInfo, nil
}

//Добавить в корзину
func AddToCart(id string) error {

	if User_id == 0 {
		return nil
	}

	product_id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	//Проверяем существует товар в корзине пользователя
	rows, err := Repository.Connection.Query(`SELECT "product_id" FROM "shopping_cart" WHERE user_id = $1 AND product_id = $2`, User_id, product_id)
	if err != nil {
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
		return err
	}
	if _, err := Repository.Connection.Exec(`DELETE FROM "shopping_cart" WHERE user_id = $1 AND product_id = $2`, User_id, product_id); err != nil {
		return err
	}
	return nil
}

//Проверка на наличие товара в корзине пользователя
func Proverka(id string) (string, error) {

	product_id, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}
	rows, err := Repository.Connection.Query(`SELECT "product_id" FROM "shopping_cart" WHERE user_id = $1 AND product_id = $2`, User_id, product_id)
	if err != nil {
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

//Увеличить количество товара в корзине
func AddKoll(id, koll string) error {

	if User_id == 0 {
		return nil
	}

	product_id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	product_koll, err := strconv.Atoi(koll)
	if err != nil {
		return err
	}
	product_koll++
	//Проверяем существует ли товар в корзине пользователя
	rows, err := Repository.Connection.Query(`SELECT "product_id" FROM "shopping_cart" WHERE user_id = $1 AND product_id = $2`, User_id, product_id)
	if err != nil {
		return err
	}
	//Ecли товар в корзине обновляем количество товара
	for rows.Next() {
		var UserCart Model.UserCart
		rows.Scan(&UserCart.Product_Id)
		if UserCart.Product_Id != 0 {
			if _, err := Repository.Connection.Exec(`UPDATE "shopping_cart" SET "product_koll" = $1 WHERE user_id = $2 AND product_id = $3`, product_koll, User_id, product_id); err != nil {
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
		return err
	}
	product_koll, err := strconv.Atoi(koll)
	if err != nil {
		return err
	}
	product_koll--
	if product_koll == 0 {
		err := DeleteFromCart(id)
		if err != nil {
			return err
		}
	}
	//Проверяем существует ли товар в корзине пользователя
	rows, err := Repository.Connection.Query(`SELECT "product_id" FROM "shopping_cart" WHERE user_id = $1 AND product_id = $2`, User_id, product_id)
	if err != nil {
		return err
	}

	//Ecли товар в корзине обновляем количество товара
	for rows.Next() {
		var UserCart Model.UserCart
		rows.Scan(&UserCart.Product_Id)
		if UserCart.Product_Id != 0 {
			if _, err := Repository.Connection.Exec(`UPDATE "shopping_cart" SET "product_koll" = $1 WHERE user_id = $2 AND product_id = $3`, product_koll, User_id, product_id); err != nil {
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
		return nil, err
	}

	if User_id != 0 {
		row, err := Repository.Connection.Query(`
		SELECT products.product_id, products.product_image, products.product_name, products.product_manufacturer, products.product_category, products.product_description, products.product_price, shopping_cart.product_koll 
		FROM products JOIN "shopping_cart" on products.product_id = shopping_cart.product_id
		WHERE user_id = $1 AND products.product_id = $2
		`, User_id, product_id)
		if err != nil {
			return nil, err
		}
		for row.Next() {
			err := row.Scan(&UserCart.Product_Id, &UserCart.Image, &UserCart.Name, &UserCart.Manufacturer, &UserCart.Category, &UserCart.Description, &UserCart.Price, &UserCart.Product_Koll)
			if err != nil {
				return nil, err
			}
		}
		UserInfo = append(UserInfo, UserCart)
	}

	if User_id == 0 || UserCart.Product_Koll == 0 {
		row, err := Repository.Connection.Query(`SELECT * FROM "products" WHERE "product_id" = $1`, product_id)
		if err != nil {
			return nil, err
		}
		UserInfo = nil
		for row.Next() {
			err := row.Scan(&UserCart.Product_Id, &UserCart.Image, &UserCart.Name, &UserCart.Manufacturer, &UserCart.Category, &UserCart.Description, &UserCart.Price)
			if err != nil {
				return nil, err
			}
			UserInfo = append(UserInfo, UserCart)
		}
	}
	return UserInfo, nil
}

func CreateProduct(p Model.Product) error {
	p.Name = strings.TrimSpace(p.Name)
	p.Image = strings.TrimSpace(p.Image)
	p.Manufacturer = strings.TrimSpace(p.Manufacturer)
	p.Category = strings.TrimSpace(p.Category)
	p.Description = strings.TrimSpace(p.Description)
	p.Price = strings.TrimSpace(p.Price)
	if p.Name == "" || p.Image == "" || p.Manufacturer == "" || p.Category == "" || p.Description == "" || p.Price == "" {
		return errors.New("невозможно добавить запись, не все поля заполнены!")
	}
	if _, err := Repository.Connection.Exec(`INSERT INTO "products" ("product_name","product_image", "product_manufacturer", "product_category", "product_description","product_price" ) VALUES ($1,$2,$3,$4,$5,$6)`, p.Name, p.Image, p.Manufacturer, p.Category, p.Description, p.Price); err != nil {
		return err
	}
	return nil
}

func UpdateProduct(p Model.Product, id string) error {
	if err := dataExist(id); err != nil {
		return err
	}
	p.Name = strings.TrimSpace(p.Name)
	p.Image = strings.TrimSpace(p.Image)
	p.Manufacturer = strings.TrimSpace(p.Manufacturer)
	p.Category = strings.TrimSpace(p.Category)
	p.Description = strings.TrimSpace(p.Description)
	p.Price = strings.TrimSpace(p.Price)
	if p.Name == "" || p.Image == "" || p.Manufacturer == "" || p.Category == "" || p.Description == "" || p.Price == "" {
		return errors.New("невозможно редактировать запись, не все поля заполнены!")
	}
	if _, err := Repository.Connection.Exec(`UPDATE "products" SET "product_name" = $1,"product_image" = $2,,"product_manufacturer" = $3,"product_category" = $4,"product_description" = $5  "product_price" = $6  WHERE "product_id" = $7`, p.Name, p.Image, p.Manufacturer, p.Category, p.Description, p.Price, id); err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id string) error {
	if err := dataExist(id); err != nil {
		return err
	}
	if _, err := Repository.Connection.Exec(`DELETE FROM "products" WHERE "product_id" = $1`, id); err != nil {
		return err
	}
	return nil
}

func dataExist(id string) error {
	persons, err := ReadOneProductById(id)
	if err != nil {
		return err
	}
	if len(persons) == 0 {
		return fmt.Errorf("записи с id = %s не существует", id)
	}
	return nil
}
