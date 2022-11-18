package Logic

import (
	"errors"
	"fmt"
	Model "myapp/internal/model"
	Repository "myapp/internal/repository"
	"strings"
)

func ReadAll() ([]Model.Product, error) {
	row, err := Repository.Connection.Query(`SELECT * FROM "Product" ORDER BY "product_id"`)
	if err != nil {
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Id, &p.Product_name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
		if err != nil {
			return nil, err
		}
		productInfo = append(productInfo, p)
	}
	return productInfo, nil
}

func Create(p Model.Product) error {
	p.Product_name = strings.TrimSpace(p.Product_name)
	p.Manufacturer = strings.TrimSpace(p.Manufacturer)
	p.Category = strings.TrimSpace(p.Category)
	p.Description = strings.TrimSpace(p.Description)
	p.Price = strings.TrimSpace(p.Price)
	if p.Product_name == "" || p.Manufacturer == "" || p.Category == "" || p.Description == "" || p.Price == "" {
		return errors.New("невозможно добавить запись, не все поля заполнены!")
	}
	if _, err := Repository.Connection.Exec(`INSERT INTO "Product" ("product_name", "product_manufacturer", "product_category", "product_description","product_price" ) VALUES ($1, $2,$3,$4)`, p.Product_name, p.Manufacturer, p.Category, p.Description, p.Price); err != nil {
		return err
	}
	return nil
}

func ReadOne(product_name string) ([]Model.Product, error) {
	row, err := Repository.Connection.Query(`SELECT * FROM "Product" WHERE "product_name" = $1`, product_name)
	if err != nil {
		return nil, err
	}
	var productInfo = []Model.Product{}
	for row.Next() {
		var p Model.Product
		err := row.Scan(&p.Id, &p.Product_name, &p.Manufacturer, &p.Category, &p.Description, &p.Price)
		if err != nil {
			return nil, err
		}
		productInfo = append(productInfo, p)
	}
	return productInfo, nil
}

func Update(p Model.Product, id string) error {
	if err := dataExist(id); err != nil {
		return err
	}
	p.Product_name = strings.TrimSpace(p.Product_name)
	p.Manufacturer = strings.TrimSpace(p.Manufacturer)
	p.Category = strings.TrimSpace(p.Category)
	p.Description = strings.TrimSpace(p.Description)
	p.Price = strings.TrimSpace(p.Price)
	if p.Product_name == "" || p.Manufacturer == "" || p.Category == "" || p.Description == "" || p.Price == "" {
		return errors.New("невозможно редактировать запись, не все поля заполнены!")
	}
	if _, err := Repository.Connection.Exec(`UPDATE "Product" SET "product_name" = $1,"product_manufacturer" = $2,"product_category" = $3,"product_description" = $4  "product_price" = $5  WHERE "product_id" = $6`, p.Product_name, p.Manufacturer, p.Category, p.Description, p.Price, id); err != nil {
		return err
	}
	return nil
}

func Delete(id string) error {
	if err := dataExist(id); err != nil {
		return err
	}
	if _, err := Repository.Connection.Exec(`DELETE FROM "Product" WHERE "product_id" = $1`, id); err != nil {
		return err
	}
	return nil
}

func dataExist(id string) error {
	persons, err := ReadOne(id)
	if err != nil {
		return err
	}
	if len(persons) == 0 {
		return fmt.Errorf("записи с id = %s не существует", id)
	}
	return nil
}
