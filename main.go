package main

import (
	"log"
	"steradian/app/controllers"
	"steradian/app/repository"
	"steradian/app/services"
	"steradian/config"
	"steradian/helpers"
	"steradian/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	env := helpers.Env{}
	env.StartingCheck()

}
func main() {
	router := gin.Default()
	initDb, err := config.Database.ConnectDB(config.Database{})
	if err != nil {
		log.Panic(err)
	}
	carRepository := repository.NewCarRepositoryImpl(initDb)
	ordersRepository := repository.NewOrdersRepositoryImpl(initDb)

	carService := services.NewCarServiceImpl(&carRepository)
	ordersService := services.NewOrdersServicesImpl(&ordersRepository)

	carControllers := controllers.NewCarController(&carService)
	ordersControllers := controllers.NewOrdersControllers(&ordersService)

	routes.SetUpCarRoute(router, &carControllers)
	routes.SetUpOrdersRoute(router, &ordersControllers)
	router.Run(":8080")
}
