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
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Login         string `json:"login"`
	HashPassword  string `json:"password"`
	Role          string `json:"role"`
	Product_Id    int    `json:"product_Id"`
	Product_Image string `json:"product_Image"`
	Product_Name  string `json:"product_Name"`
	Product_Price string `json:"product_Price"`
	Product_Koll  string `json:"koll"`
}
