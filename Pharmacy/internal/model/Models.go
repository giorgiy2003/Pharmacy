package Model

type Product struct {
	Id           int    `json:"id"`
	Image        string `json:"image"`
	Name         string `json:"product_name"`
	Manufacturer string `json:"manufacturer"`
	Category     string `json:"category"`
	Description  string `json:"description"`
	Price        string `json:"price"`
}

type User struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	HashPassword string `json:"password"`
	UserCard     []int  //Корзина товаров
}
