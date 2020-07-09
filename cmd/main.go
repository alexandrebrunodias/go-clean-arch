package main

import (
	"github.com/alexandrebrundias/product-crud/application/delivery/http/middleware"
	"github.com/alexandrebrundias/product-crud/application/product"
	"github.com/alexandrebrundias/product-crud/infrastructure/database"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigFile("config.yml")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := database.NewDatabase().Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	e.Use(middleware.InitMiddleware().CORS)

	productRepoistory := product.NewRepository(db)
	productUsecase := product.NewUsecase(productRepoistory)
	product.NewHandler(e, productUsecase)

	log.Fatal(e.Start(":" + viper.GetString("port")))
}
