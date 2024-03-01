package mocks

import (
	"steradian/app/models/requests"
	"steradian/app/models/responses"

	"github.com/gin-gonic/gin"
)

type MockCarService struct {
	CreateCarfn        func(ctx *gin.Context, request *requests.CreateCarRequests) responses.GenericResponse
	GetDetailCarfn     func(ctx *gin.Context, request *requests.GetDetailCarRequest) responses.GenericResponse
	UpdateCarByCarIdfn func(ctx *gin.Context, request *requests.UpdateCarByCarIdRequest) responses.GenericResponse
	DeleteCarByCarIdfn func(ctx *gin.Context, request *requests.DeleteCarRequest) responses.GenericResponse
	GetListCarfn       func(ctx *gin.Context) responses.GenericResponse
}

func (mC *MockCarService) CreateCar(ctx *gin.Context, request *requests.CreateCarRequests) responses.GenericResponse {
	if mC.CreateCarfn != nil {
		return mC.CreateCarfn(ctx, request)
	}
	return responses.GenericResponse{}
}
func (mC *MockCarService) GetDetailCar(ctx *gin.Context, request *requests.GetDetailCarRequest) responses.GenericResponse {
	if mC.GetDetailCarfn != nil {
		return mC.GetDetailCarfn(ctx, request)
	}
	return responses.GenericResponse{}
}
func (mC *MockCarService) UpdateCarByCarId(ctx *gin.Context, request *requests.UpdateCarByCarIdRequest) responses.GenericResponse {
	if mC.UpdateCarByCarIdfn != nil {
		return mC.UpdateCarByCarIdfn(ctx, request)
	}
	return responses.GenericResponse{}
}
func (mC *MockCarService) DeleteCarByCarId(ctx *gin.Context, request *requests.DeleteCarRequest) responses.GenericResponse {
	if mC.DeleteCarByCarIdfn != nil {
		return mC.DeleteCarByCarIdfn(ctx, request)
	}
	return responses.GenericResponse{}
}
func (mC *MockCarService) GetListCar(ctx *gin.Context) responses.GenericResponse {
	if mC.GetListCarfn != nil {
		return mC.GetListCarfn(ctx)
	}
	return responses.GenericResponse{}
}
