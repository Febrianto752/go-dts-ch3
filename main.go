package main

import (
	"log"

	"github.com/Febrianto752/go-dts-ch3/config"
	"github.com/Febrianto752/go-dts-ch3/handler"
	"github.com/Febrianto752/go-dts-ch3/repository"
	"github.com/Febrianto752/go-dts-ch3/route"
	"github.com/Febrianto752/go-dts-ch3/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("gagal mengambil .env %v", err)

	}

	db := config.InitializeDB()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	router := route.NewRouter(userHandler, productHandler)

	router.Run(":8080")
}
