package services

import (
	"steradian/app/models/requests"
	"steradian/app/models/responses"

	"github.com/gin-gonic/gin"
)

type CarService interface {
	CreateCar(ctx *gin.Context, request *requests.CreateCarRequests) responses.GenericResponse
	GetDetailCar(ctx *gin.Context, request *requests.GetDetailCarRequest) responses.GenericResponse
	UpdateCarByCarId(ctx *gin.Context, request *requests.UpdateCarByCarIdRequest) responses.GenericResponse
	DeleteCarByCarId(ctx *gin.Context, request *requests.DeleteCarRequest) responses.GenericResponse
	GetListCar(ctx *gin.Context) responses.GenericResponse
}
