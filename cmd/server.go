package main

import (
	"github.com/alexandrebrundias/product-crud/infrastructure/database"
	"github.com/alexandrebrundias/product-crud/product/delivery/http"
	"github.com/alexandrebrundias/product-crud/product/repository"
	"github.com/alexandrebrundias/product-crud/product/usecase"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigFile("../config.yml")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
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

	productRepoistory := repository.NewRepoistory(db)
	productUsecase := usecase.NewProductUsecase(productRepoistory)
	http.NewProductHandler(e, productUsecase)

	log.Fatal(e.Start(":" + viper.GetString("port")))
}
