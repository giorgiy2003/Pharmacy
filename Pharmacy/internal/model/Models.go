package Model

type Product struct {
	Product_Id           int    `json:"id"`
	Product_Image        string `json:"image"`
	Product_Name         string `json:"product_name"`
	Product_Manufacturer string `json:"manufacturer"`
	Product_Category     string `json:"category"`
	Product_Description  string `json:"description"`
	Product_Price        string `json:"price"`
}

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Login        string `json:"login"`
	HashPassword string `json:"password"`
	Role         string `json:"role"`
}

type UserCart struct {
	Product_Id           int    `json:"product_Id"`
	Product_Image        string `json:"image"`
	Product_Name         string `json:"product_name"`
	Product_Manufacturer string `json:"manufacturer"`
	Product_Category     string `json:"category"`
	Product_Description  string `json:"description"`
	Product_Price        string `json:"price"`
	Product_Koll         int    `json:"product_Koll"`
	Product_amount       int    `json:"product_amount"`
}

type Order struct {
	Order_Id       int    `json:"order_Id"`
	User_Id        int    `json:"user_Id"`
	Product_Id     int    `json:"product_Id"`
	Product_Image  string `json:"image"`
	Product_Name   string `json:"product_name"`
	Product_Price  string `json:"price"`
	Product_Koll   int    `json:"product_Koll"`
	Product_amount int    `json:"product_amount"`
	Order_time     string `json:"order_time"`
	Order_status   string `json:"order_status"`
	Track_number   string `json:"track_number"`
	Delivery_price int    `json:"delivery_price"`
	Total_price    int    `json:"total_price"`
	//Данные покупателя
	Customer_Name    string `json:"customer_Name"`
	Customer_Address string `json:"customer_Address"`
	Customer_Email   string `json:"customer_Email"`
	Customer_Phone   string `json:"customer_Phone"`
	Customer_Comment string `json:"customer_Comment"`
}
