package routes

import (
	"steradian/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpCarRoute(router *gin.Engine, carControllers *controllers.CarController) {

	router.POST("/insert-car", carControllers.CreateCar)
	router.GET("/get-list-car", carControllers.GetListCar)
	router.GET("/get-detail-car/:carId", carControllers.GetDetailCar)
	router.PUT("/update-car", carControllers.UpdateCarByCarId)
	router.DELETE("/delete-car/:carId", carControllers.DeleteCarByCarId)

}
func SetUpOrdersRoute(router *gin.Engine, ordersControllers *controllers.OrdersControllers) {
	router.POST("/insert-orders", ordersControllers.CreateOrders)
	router.GET("/get-list-orders", ordersControllers.GetListOrders)
	router.GET("/get-detail-orders/:ordersId", ordersControllers.GetDetailOrders)
	router.PUT("/update-orders", ordersControllers.UpdateOrdersByOrdersId)
	router.DELETE("/delete-orders/:ordersId", ordersControllers.DeleteOrdersByOrdersId)
}
