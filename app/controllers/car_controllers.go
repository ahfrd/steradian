package controllers

import (
	"encoding/json"
	"net/http"
	"steradian/app/models/requests"
	"steradian/app/services"
	"steradian/helpers"

	"github.com/gin-gonic/gin"
)

type CarController struct {
	CarService services.CarService
}

func NewCarController(carServices *services.CarService) CarController {
	return CarController{
		CarService: *carServices,
	}
}

func (cc *CarController) CreateCar(ctx *gin.Context) {
	var bodyReq requests.CreateCarRequests
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	helpers.LogRequest(ctx, string(requestData))
	response := cc.CarService.CreateCar(ctx, &bodyReq)
	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	helpers.LogResponse(ctx, string(responseData))
	ctx.JSON(response.Code, response)
	return
}
func (cc *CarController) GetDetailCar(ctx *gin.Context) {
	var bodyReq requests.GetDetailCarRequest
	bodyReq.CarId = ctx.Param("carId")
	if err := ctx.ShouldBindQuery(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	helpers.LogRequest(ctx, string(requestData))
	response := cc.CarService.GetDetailCar(ctx, &bodyReq)
	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	helpers.LogResponse(ctx, string(responseData))
	ctx.JSON(response.Code, response)
	return
}

func (cc *CarController) UpdateCarByCarId(ctx *gin.Context) {
	var bodyReq requests.UpdateCarByCarIdRequest

	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	helpers.LogRequest(ctx, string(requestData))
	response := cc.CarService.UpdateCarByCarId(ctx, &bodyReq)
	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	helpers.LogResponse(ctx, string(responseData))
	ctx.JSON(response.Code, response)
	return
}

func (cc *CarController) DeleteCarByCarId(ctx *gin.Context) {
	var bodyReq requests.DeleteCarRequest
	bodyReq.CarId = ctx.Param("carId")
	if err := ctx.ShouldBindQuery(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	helpers.LogRequest(ctx, string(requestData))
	response := cc.CarService.DeleteCarByCarId(ctx, &bodyReq)
	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	helpers.LogResponse(ctx, string(responseData))
	ctx.JSON(response.Code, response)
	return
}

func (cc *CarController) GetListCar(ctx *gin.Context) {

	helpers.LogRequest(ctx, string("get-list-car"))
	response := cc.CarService.GetListCar(ctx)
	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	helpers.LogResponse(ctx, string(responseData))
	ctx.JSON(response.Code, response)
	return
}
