package services

import (
	"steradian/app/models/requests"
	"steradian/app/models/responses"

	"github.com/gin-gonic/gin"
)

type OrdersService interface {
	CreateOrders(ctx *gin.Context, request *requests.CreateOrdersRequests) responses.GenericResponse
	GetDetailOrders(ctx *gin.Context, request *requests.GetDetailOrdersRequest) responses.GenericResponse
	UpdateOrdersByOrdersId(ctx *gin.Context, request *requests.UpdateOrdersByOrdersIdRequest) responses.GenericResponse
	DeleteOrdersByOrdersId(ctx *gin.Context, request *requests.DeleteOrdersRequest) responses.GenericResponse
	GetListOrders(ctx *gin.Context) responses.GenericResponse
}
