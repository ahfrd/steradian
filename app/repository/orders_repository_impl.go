package repository

import (
	"steradian/app/models/requests"
	"steradian/app/models/responses"

	"github.com/gin-gonic/gin"
)

type OrdersRepository interface {
	CreateOrders(ctx *gin.Context, request *requests.CreateOrdersRequests) error
	SelectListDataOrders(ctx *gin.Context) ([]responses.SelectListDataOrdersResponses, error)
	SelectDataOrdersById(ctx *gin.Context, ordersId string) (*responses.SelectListDataOrdersResponses, error)
	UpdateOrdersByOrdersId(ctx *gin.Context, request *requests.UpdateOrdersByOrdersIdRequest) error
	DeleteOrdersByOrdersId(ctx *gin.Context, ordersId string) error
}
