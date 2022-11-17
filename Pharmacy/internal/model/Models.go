package Model

type Product struct {
	Id           int    `json:"id"`
	Product_name string `json:"product_name"`
	Manufacturer string `json:"manufacturer"`
	Category     string `json:"category"`
	Description  string `json:"description"`
}
