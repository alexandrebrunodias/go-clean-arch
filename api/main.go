package main

import (
	"github.com/alexandrebrundias/product-crud/api/common"
	"github.com/alexandrebrundias/product-crud/api/handler"
	"github.com/alexandrebrundias/product-crud/infrastructure/database"
	"github.com/alexandrebrundias/product-crud/product"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

func init() {
	abs, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigFile(abs + "/config.yml")
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
	e.Use(common.InitMiddleware().CORS)

	productRepoistory := product.NewRepository(db)
	productUsecase := product.NewUsecase(productRepoistory)
	handler.NewProductHandler(e, productUsecase)

	log.Fatal(e.Start(":" + viper.GetString("port")))
}
