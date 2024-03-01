package mocks

import (
	"steradian/app/models/requests"
	"steradian/app/models/responses"

	"github.com/gin-gonic/gin"
)

type MockOrdersService struct {
	CreateOrdersfn           func(ctx *gin.Context, request *requests.CreateOrdersRequests) responses.GenericResponse
	GetDetailOrdersfn        func(ctx *gin.Context, request *requests.GetDetailOrdersRequest) responses.GenericResponse
	UpdateOrdersByOrdersIdfn func(ctx *gin.Context, request *requests.UpdateOrdersByOrdersIdRequest) responses.GenericResponse
	DeleteOrdersByOrdersIdfn func(ctx *gin.Context, request *requests.DeleteOrdersRequest) responses.GenericResponse
	GetListOrdersfn          func(ctx *gin.Context) responses.GenericResponse
}

func (mC *MockOrdersService) CreateOrders(ctx *gin.Context, request *requests.CreateOrdersRequests) responses.GenericResponse {
	if mC.CreateOrdersfn != nil {
		return mC.CreateOrdersfn(ctx, request)
	}
	return responses.GenericResponse{}
}
func (mC *MockOrdersService) GetDetailOrders(ctx *gin.Context, request *requests.GetDetailOrdersRequest) responses.GenericResponse {
	if mC.GetDetailOrdersfn != nil {
		return mC.GetDetailOrdersfn(ctx, request)
	}
	return responses.GenericResponse{}
}
func (mC *MockOrdersService) UpdateOrdersByOrdersId(ctx *gin.Context, request *requests.UpdateOrdersByOrdersIdRequest) responses.GenericResponse {
	if mC.UpdateOrdersByOrdersIdfn != nil {
		return mC.UpdateOrdersByOrdersIdfn(ctx, request)
	}
	return responses.GenericResponse{}
}
func (mC *MockOrdersService) DeleteOrdersByOrdersId(ctx *gin.Context, request *requests.DeleteOrdersRequest) responses.GenericResponse {
	if mC.DeleteOrdersByOrdersIdfn != nil {
		return mC.DeleteOrdersByOrdersIdfn(ctx, request)
	}
	return responses.GenericResponse{}
}
func (mC *MockOrdersService) GetListOrders(ctx *gin.Context) responses.GenericResponse {
	if mC.GetListOrdersfn != nil {
		return mC.GetListOrdersfn(ctx)
	}
	return responses.GenericResponse{}
}
