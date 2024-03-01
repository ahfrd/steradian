package controllers

import (
	"encoding/json"
	"net/http"
	"steradian/app/models/requests"
	"steradian/app/services"
	"steradian/helpers"

	"github.com/gin-gonic/gin"
)

type OrdersControllers struct {
	OrdersService services.OrdersService
}

func NewOrdersControllers(orderServices *services.OrdersService) OrdersControllers {
	return OrdersControllers{
		OrdersService: *orderServices,
	}
}

func (cc *OrdersControllers) CreateOrders(ctx *gin.Context) {
	var bodyReq requests.CreateOrdersRequests
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
	response := cc.OrdersService.CreateOrders(ctx, &bodyReq)
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
func (cc *OrdersControllers) GetDetailOrders(ctx *gin.Context) {
	var bodyReq requests.GetDetailOrdersRequest
	bodyReq.OrdersId = ctx.Param("ordersId")
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
	response := cc.OrdersService.GetDetailOrders(ctx, &bodyReq)
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

func (cc *OrdersControllers) UpdateOrdersByOrdersId(ctx *gin.Context) {
	var bodyReq requests.UpdateOrdersByOrdersIdRequest

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
	response := cc.OrdersService.UpdateOrdersByOrdersId(ctx, &bodyReq)
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

func (cc *OrdersControllers) DeleteOrdersByOrdersId(ctx *gin.Context) {
	var bodyReq requests.DeleteOrdersRequest

	bodyReq.OrdersId = ctx.Param("ordersId")
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
	response := cc.OrdersService.DeleteOrdersByOrdersId(ctx, &bodyReq)
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

func (cc *OrdersControllers) GetListOrders(ctx *gin.Context) {

	helpers.LogRequest(ctx, string("get-list-Orders"))
	response := cc.OrdersService.GetListOrders(ctx)
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
