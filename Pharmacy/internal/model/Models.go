package Model

type Product struct {
	Product_Id   int    `json:"id"`
	Image        string `json:"image"`
	Name         string `json:"product_name"`
	Manufacturer string `json:"manufacturer"`
	Category     string `json:"category"`
	Description  string `json:"description"`
	Price        string `json:"price"`
}

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Login        string `json:"login"`
	HashPassword string `json:"password"`
	Role         string `json:"role"`
}

type UserCart struct {
	User_Id      int    `json:"user_Id"`
	Product_Id   int    `json:"product_Id"`
	Image        string `json:"image"`
	Name         string `json:"product_name"`
	Manufacturer string `json:"manufacturer"`
	Category     string `json:"category"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	Product_Koll int    `json:"product_Koll"`
}
