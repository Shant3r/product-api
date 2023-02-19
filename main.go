package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/shant3r/product-api/internal/api"
	"github.com/shant3r/product-api/internal/db"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "usr"
	password = "pwd"
	dbname   = "products"
)

func main() {
	ctx := context.Background()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	repository := db.New(database)

	h := api.New(repository)

	// r.GET("/products", func(c *gin.Context) { h.GetProducts(ctx, c) })
	r.POST("/products", func(c *gin.Context) { h.AddProduct(ctx, c) })
	r.POST("/products/item", func(c *gin.Context) { h.AddProductItem(ctx, c) })
	r.POST("/products/price", func (c *gin.Context) { h.AddProductPrice(ctx,c)})
	r.Run()
}
