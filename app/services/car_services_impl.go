package services

import (
	"fmt"
	"net/http"
	"steradian/app/models/requests"
	"steradian/app/models/responses"
	"steradian/app/repository"

	"github.com/gin-gonic/gin"
)

type carServicesImpl struct {
	CarRepository repository.CarRepository
}

func NewCarServicesImpl(carRepository *repository.CarRepository) CarService {
	return &carServicesImpl{
		CarRepository: *carRepository,
	}
}

func (cs *carServicesImpl) CreateCar(ctx *gin.Context, request *requests.CreateCarRequests) responses.GenericResponse {
	var resData responses.GenericResponse
	insertCar := cs.CarRepository.CreateCar(ctx, request)
	if insertCar != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = fmt.Sprintf("maaf anda tidak bisa membuat data : %s", insertCar.Error())
		resData.Status = "error insert car "
		return resData
	}
	resData.Code = http.StatusOK
	resData.Message = "succses insert"
	resData.Status = "sucses"
	return resData
}

func (cs *carServicesImpl) GetListCar(ctx *gin.Context) responses.GenericResponse {
	var resData responses.GenericResponse
	selectListDataCar, err := cs.CarRepository.SelectListDataCar(ctx)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = fmt.Sprintf("maaf anda gagal mendapatkan data mobil : %s", err.Error())
		resData.Status = "failed get data"
		return resData
	}
	resData.Code = http.StatusOK
	resData.Message = "sucses get data"
	resData.Status = "sucses"
	resData.Data = selectListDataCar
	return resData
}
func (cs *carServicesImpl) GetDetailCar(ctx *gin.Context, request *requests.GetDetailCarRequest) responses.GenericResponse {
	var resData responses.GenericResponse
	getDetail, err := cs.CarRepository.SelectDataCarById(ctx, request.CarId)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = fmt.Sprintf("maaf anda gagal mendapatkan data mobil : %s", err.Error())
		resData.Status = "failed get data"
		return resData
	}
	resData.Code = http.StatusOK
	resData.Message = "sucses get data"
	resData.Status = "sucses"
	resData.Data = getDetail
	return resData
}

func (cs *carServicesImpl) UpdateCarByCarId(ctx *gin.Context, request *requests.UpdateCarByCarIdRequest) responses.GenericResponse {

	var resData responses.GenericResponse
	err := cs.CarRepository.UpdateCarByCarId(ctx, request)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = fmt.Sprintf("maaf anda gagal update data mobil : %s", err.Error())
		resData.Status = "failed get data"
		return resData
	}
	resData.Code = http.StatusOK
	resData.Message = "sucses updated data"
	resData.Status = "sucses"
	return resData
}
func (cs *carServicesImpl) DeleteCarByCarId(ctx *gin.Context, request *requests.DeleteCarRequest) responses.GenericResponse {

	var resData responses.GenericResponse
	err := cs.CarRepository.DeleteCarByCarId(ctx, request.CarId)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = fmt.Sprintf("maaf anda gagal delete data mobil : %s", err.Error())
		resData.Status = "failed get data"
		return resData
	}
	resData.Code = http.StatusOK
	resData.Message = "sucses delete data"
	resData.Status = "sucses"
	return resData
}
