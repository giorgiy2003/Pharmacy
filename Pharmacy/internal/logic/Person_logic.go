package Logic

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	Model "myapp/internal/model"
	Repository "myapp/internal/repository"
	"strconv"
	"strings"
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
		err := row.Scan(&p.Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
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
		err := row.Scan(&p.Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
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
		err := row.Scan(&p.Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
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
		err := row.Scan(&p.Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
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
		err := row.Scan(&p.Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
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
		err := row.Scan(&p.Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
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
		err := row.Scan(&p.Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
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
		err := row.Scan(&p.Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
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
		err := row.Scan(&p.Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
		if err != nil {
			return nil, err
		}
		productInfo = append(productInfo, p)
	}
	return productInfo, nil
}

var Auth string
var User_id int
var Login, Password string

//Авторизация
func Autorization(login, password string) error {
	Login = strings.TrimSpace(login)

	//Hashed value of password
	Password = SHA_256_Encode(password) //Кодируем значение пароля пользователя

	row, err := Repository.Connection.Query(`SELECT "user_id", "user_login", "user_password" FROM "users" WHERE user_login = $1 AND user_password = $2`, Login, Password)
	if err != nil {
		return err
	}

	var User Model.User
	for row.Next() {
		row.Scan(&User.Id, &User.Login, &User.HashPassword)
	}

	//Если значения структуры пусты возращаем ошибку
	if User.Id == 0 {
		return errors.New("Введён неверный логин или пароль!")
	}

	Auth = "true"
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
func UserCart() ([]Model.Product, error) {
	row, err := Repository.Connection.Query(`SELECT "user_id" FROM "users" WHERE user_login = $1 AND user_password = $2`, Login, Password)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		row.Scan(&User_id)
	}
	rows, err := Repository.Connection.Query(`SELECT * FROM "shopping_cart" WHERE user_id = $1`, User_id)
	if err != nil {
		return nil, err
	}

	var UserInfo = []Model.User{}
	for rows.Next() {
		var User Model.User
		rows.Scan(&User.Product_Id, &User.Product_Koll)
		UserInfo = append(UserInfo, User)
	}

	var productInfo = []Model.Product{}
	for _, Info := range UserInfo {

		row, err := Repository.Connection.Query(`SELECT * FROM "products" WHERE "product_id" = $1`, Info.Id)
		if err != nil {
			return nil, err
		}
		
		for row.Next() {
			var p Model.Product
			err := row.Scan(&p.Id, &p.Image, &p.Name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
			if err != nil {
				return nil, err
			}
			productInfo = append(productInfo, p)
		}

	}

	log.Println(productInfo)
	return productInfo, nil
}

//Добавить в корзину
func AddToCart(id, koll string) error {
	product_id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if _, err := Repository.Connection.Exec(`INSERT INTO "shopping_cart" ("user_id","product_id", "quantity") VALUES ($1,$2,$3)`,User_id,product_id, koll); err != nil {
		return err
	}
	return nil
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
