package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var db *gorm.DB
var err error

func main() {
	r := gin.Default()
	db, err = gorm.Open(mysql.Open("root:admin_12345@tcp(127.0.0.1:3306)/demodb_1"))
	if err != nil {
		fmt.Println("error", err)
	}
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/api/products", GetProducts)
	r.POST("/api/products", CreateProudct)
	r.PUT("/api/products/:id", UpdateProduct)
	r.Run(":8010")

}

func GetProducts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func CreateProudct(c *gin.Context) {
	var inputProduct CreateProductInput
	if err := c.ShouldBindJSON(&inputProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	product := Product{
		Name:  inputProduct.Name,
		Price: inputProduct.Price,
	}
	db = c.MustGet("db").(*gorm.DB)
	db.Create(&product)
	c.JSON(http.StatusOK, gin.H{"message": "product created"})
}

func UpdateProduct(c *gin.Context) {
	var product Product
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id=?", c.Param("id")).First(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "record not found"})
		return
	}
	var inputProduct UpdateProductInput
	if err = c.ShouldBindJSON(&inputProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&product).Updates(inputProduct)
	c.JSON(http.StatusOK, gin.H{"data": product})
}
