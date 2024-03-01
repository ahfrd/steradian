package repository

import (
	"steradian/app/models/requests"
	"steradian/app/models/responses"

	"github.com/gin-gonic/gin"
)

type CarRepository interface {
	CreateCar(ctx *gin.Context, request *requests.CreateCarRequests) error
	SelectListDataCar(ctx *gin.Context) ([]responses.SelectListDataCarResponses, error)
	SelectDataCarById(ctx *gin.Context, carId string) (*responses.SelectListDataCarResponses, error)
	UpdateCarByCarId(ctx *gin.Context, request *requests.UpdateCarByCarIdRequest) error
	DeleteCarByCarId(ctx *gin.Context, carId string) error
}
