package Handler

import (
	Repository "myapp/internal/repository"

	"github.com/gin-gonic/gin"
)

func Shop(c *gin.Context) {
	/*Products, err := Logic.ReadAll()
	if err != nil {
		log.Println(err)
		c.HTML(400, "InfoPage", nil)
		return
	}*/
	c.HTML(200, "shop", gin.H{})
	/*for _, Product := range Products {
		c.HTML(200, "shop", gin.H {
			"Id":           Product.Id,
			"Product_name": Product.Product_name,
			"Manufacturer": Product.Manufacturer,
			"Category":     Product.Category,
			"Description":  Product.Description,
			"Price":  Product.Price,
		})
	}*/
	return
}

func ConnectDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := Repository.OpenTable(); err != nil {
			c.HTML(500, "InternalServerError", gin.H{
				"Error": err,
			})
			return
		}
	}
}

func MainForm(c *gin.Context) {
	c.HTML(200, "index", nil)
}

func Shop_single(c *gin.Context) {
	c.HTML(200, "shop-single", nil)
}

func Cart(c *gin.Context) {
	c.HTML(200, "cart", nil)
}

func About(c *gin.Context) {
	c.HTML(200, "about", nil)
}

func Checkout(c *gin.Context) {
	c.HTML(200, "checkout", nil)
}

func Contact(c *gin.Context) {
	c.HTML(200, "contact", nil)
}

func Thanks(c *gin.Context) {
	c.HTML(200, "thankyou", nil)
}
