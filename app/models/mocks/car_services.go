package mocks

import (
	"steradian/app/models/requests"
	"steradian/app/models/responses"

	"github.com/gin-gonic/gin"
)

type MockCarRepository struct {
	CreateCarfn         func(ctx *gin.Context, request *requests.CreateCarRequests) (int, error)
	SelectListDataCarfn func(ctx *gin.Context) ([]responses.SelectListDataCarResponses, error)
	SelectDataCarByIdfn func(ctx *gin.Context, carId string) (*responses.SelectListDataCarResponses, error)
	UpdateCarByCarIdfn  func(ctx *gin.Context, request *requests.UpdateCarByCarIdRequest) error
	DeleteCarByCarIdfn  func(ctx *gin.Context, carId string) error

	// CreateCar(ctx *gin.Context, request *requests.CreateCarRequests) (int, error)
	// SelectListDataCar(ctx *gin.Context) ([]responses.SelectListDataCarResponses, error)
	// SelectDataCarById(ctx *gin.Context, carId string) (*responses.SelectListDataCarResponses, error)
	// UpdateCarByCarId(ctx *gin.Context, request *requests.UpdateCarByCarIdRequest) error
	// DeleteCarByCarId(ctx *gin.Context, carId string) error
}

func (mR *MockCarRepository) CreateCar(ctx *gin.Context, request *requests.CreateCarRequests) (int, error) {
	return mR.CreateCarfn(ctx, request)
}
func (mR *MockCarRepository) SelectListDataCar(ctx *gin.Context) ([]responses.SelectListDataCarResponses, error) {
	return mR.SelectListDataCarfn(ctx)
}
func (mR *MockCarRepository) SelectDataCarById(ctx *gin.Context, carId string) (*responses.SelectListDataCarResponses, error) {
	return mR.SelectDataCarByIdfn(ctx, carId)
}
func (mR *MockCarRepository) UpdateCarByCarId(ctx *gin.Context, request *requests.UpdateCarByCarIdRequest) error {
	return mR.UpdateCarByCarIdfn(ctx, request)
}
func (mR *MockCarRepository) DeleteCarByCarId(ctx *gin.Context, carId string) error {
	return mR.DeleteCarByCarIdfn(ctx, carId)
}
