package services

import (
	"fmt"
	"net/http"
	"steradian/app/models/requests"
	"steradian/app/models/responses"
	"steradian/app/repository"
	"steradian/helpers"

	"github.com/gin-gonic/gin"
)

type OrdersServicesImpl struct {
	OrdersRepository repository.OrdersRepository
}

func NewOrdersServicesImpl(ordersRepository *repository.OrdersRepository) OrdersService {
	return &OrdersServicesImpl{
		OrdersRepository: *ordersRepository,
	}
}

func (cs *OrdersServicesImpl) CreateOrders(ctx *gin.Context, request *requests.CreateOrdersRequests) responses.GenericResponse {
	var resData responses.GenericResponse
	request.DropoffDate = helpers.FormatDateString(request.DropoffDate)
	request.PickupDate = helpers.FormatDateString(request.PickupDate)
	request.OrderDate = helpers.FormatDateString(request.OrderDate)

	insertOrders := cs.OrdersRepository.CreateOrders(ctx, request)
	if insertOrders != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = fmt.Sprintf("maaf anda tidak bisa membuat data : %s", insertOrders.Error())
		resData.Status = "error insert Orders "
		return resData
	}
	resData.Code = http.StatusOK
	resData.Message = "succses insert"
	resData.Status = "sucses"
	return resData
}

func (cs *OrdersServicesImpl) GetListOrders(ctx *gin.Context) responses.GenericResponse {
	var resData responses.GenericResponse
	selectListDataOrders, err := cs.OrdersRepository.SelectListDataOrders(ctx)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = "maaf anda gagal mendapatkan data mobil"
		resData.Status = "failed get data"
		return resData
	}
	resData.Code = http.StatusOK
	resData.Message = "sucses get data"
	resData.Status = "sucses"
	resData.Data = selectListDataOrders
	return resData
}
func (cs *OrdersServicesImpl) GetDetailOrders(ctx *gin.Context, request *requests.GetDetailOrdersRequest) responses.GenericResponse {
	var resData responses.GenericResponse
	getDetail, err := cs.OrdersRepository.SelectDataOrdersById(ctx, request.OrdersId)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = fmt.Sprintf("maaf anda gagal mendapatkan data orders : %s", err.Error())
		resData.Status = "failed get data"
		return resData
	}
	resData.Code = http.StatusOK
	resData.Message = "sucses get data"
	resData.Status = "sucses"
	resData.Data = getDetail
	return resData
}

func (cs *OrdersServicesImpl) UpdateOrdersByOrdersId(ctx *gin.Context, request *requests.UpdateOrdersByOrdersIdRequest) responses.GenericResponse {

	var resData responses.GenericResponse
	request.DropoffDate = helpers.FormatDateString(request.DropoffDate)
	request.PickupDate = helpers.FormatDateString(request.PickupDate)
	request.OrderDate = helpers.FormatDateString(request.OrderDate)
	err := cs.OrdersRepository.UpdateOrdersByOrdersId(ctx, request)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = fmt.Sprintf("maaf anda gagal update data orders %s", err.Error())
		resData.Status = "failed update data"
		return resData
	}
	resData.Code = http.StatusOK
	resData.Message = "sucses updated data"
	resData.Status = "sucses"
	return resData
}
func (cs *OrdersServicesImpl) DeleteOrdersByOrdersId(ctx *gin.Context, request *requests.DeleteOrdersRequest) responses.GenericResponse {

	var resData responses.GenericResponse
	err := cs.OrdersRepository.DeleteOrdersByOrdersId(ctx, request.OrdersId)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = fmt.Sprintf("maaf anda gagal delete data orders %s", err.Error())
		resData.Status = "failed get data"
		return resData
	}
	resData.Code = http.StatusOK
	resData.Message = "sucses delete data"
	resData.Status = "sucses"
	return resData
}
