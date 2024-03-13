package mocks

import (
	"steradian/app/models/requests"
	"steradian/app/models/responses"

	"github.com/gin-gonic/gin"
)

type MockOrdersRepository struct {
	CreateOrdersfn           func(ctx *gin.Context, request *requests.CreateOrdersRequests) error
	SelectListDataOrdersfn   func(ctx *gin.Context) ([]responses.SelectListDataOrdersResponses, error)
	SelectDataOrdersByIdfn   func(ctx *gin.Context, ordersId string) (*responses.SelectListDataOrdersResponses, error)
	UpdateOrdersByOrdersIdfn func(ctx *gin.Context, request *requests.UpdateOrdersByOrdersIdRequest) error
	DeleteOrdersByOrdersIdfn func(ctx *gin.Context, ordersId string) error
}

func (mR *MockOrdersRepository) CreateOrders(ctx *gin.Context, request *requests.CreateOrdersRequests) error {
	return mR.CreateOrdersfn(ctx, request)
}
func (mR *MockOrdersRepository) SelectListDataOrders(ctx *gin.Context) ([]responses.SelectListDataOrdersResponses, error) {
	return mR.SelectListDataOrdersfn(ctx)
}
func (mR *MockOrdersRepository) SelectDataOrdersById(ctx *gin.Context, ordersId string) (*responses.SelectListDataOrdersResponses, error) {
	return mR.SelectDataOrdersByIdfn(ctx, ordersId)
}
func (mR *MockOrdersRepository) UpdateOrdersByOrdersId(ctx *gin.Context, request *requests.UpdateOrdersByOrdersIdRequest) error {
	return mR.UpdateOrdersByOrdersIdfn(ctx, request)
}
func (mR *MockOrdersRepository) DeleteOrdersByOrdersId(ctx *gin.Context, ordersId string) error {
	return mR.DeleteOrdersByOrdersIdfn(ctx, ordersId)
}
