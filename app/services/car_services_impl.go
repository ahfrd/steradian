package services

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"steradian/app/models/requests"
	"steradian/app/models/responses"
	"steradian/app/repository"
	"steradian/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarServiceImpl struct {
	CarRepository repository.CarRepository
}

func NewCarServiceImpl(carRepository *repository.CarRepository) CarService {
	return &CarServiceImpl{
		CarRepository: *carRepository,
	}
}

func (cs *CarServiceImpl) CreateCar(ctx *gin.Context, request *requests.CreateCarRequests) responses.GenericResponse {
	var resData responses.GenericResponse
	lastInsertID, err := cs.CarRepository.CreateCar(ctx, request)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = fmt.Sprintf("maaf anda tidak bisa membuat data : %s", err.Error())
		resData.Status = "error insert car "
		return resData
	}
	linkToB64, err := helpers.LinkToBase64(request.ImageCar)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = fmt.Sprintf("maaf ada kesalahan : %s", err.Error())
		resData.Status = "error insert car "
		return resData
	}
	imgName := helpers.Base64ToImage(linkToB64, lastInsertID)
	updateRequest := &requests.UpdateCarByCarIdRequest{
		CarName:   request.CarName,
		CarId:     strconv.Itoa(lastInsertID),
		DayRate:   request.DayRate,
		MonthRate: request.MonthRate,
		ImageCar:  imgName,
	}
	_ = cs.CarRepository.UpdateCarByCarId(ctx, updateRequest)
	resData.Code = http.StatusOK
	resData.Message = "succses insert"
	resData.Status = "sucses"
	return resData
}

func (cs *CarServiceImpl) GetListCar(ctx *gin.Context) responses.GenericResponse {
	var resData responses.GenericResponse
	selectListDataCar, err := cs.CarRepository.SelectListDataCar(ctx)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = fmt.Sprintf("maaf anda gagal mendapatkan data mobil : %s", err.Error())
		resData.Status = "failed get data"
		return resData
	}
	for i, item := range selectListDataCar {
		bytes, err := ioutil.ReadFile(item.ImageChar)
		if err != nil {
			log.Fatal(err)
		}

		var base64Encoding string

		// Determine the content type of the image file
		mimeType := http.DetectContentType(bytes)

		// Prepend the appropriate URI scheme header depending
		// on the MIME type
		switch mimeType {
		case "image/jpeg":
			base64Encoding += "data:image/jpeg;base64,"
		case "image/png":
			base64Encoding += "data:image/png;base64,"
		case "image/gif":
			base64Encoding += "data:image/gif;base64,"
		}
		base64Encoding += base64.StdEncoding.EncodeToString(bytes)
		selectListDataCar[i].ImageChar = base64Encoding
	}
	resData.Code = http.StatusOK
	resData.Message = "sucses get data"
	resData.Status = "sucses"
	resData.Data = selectListDataCar
	return resData
}
func (cs *CarServiceImpl) GetDetailCar(ctx *gin.Context, request *requests.GetDetailCarRequest) responses.GenericResponse {
	var resData responses.GenericResponse
	getDetail, err := cs.CarRepository.SelectDataCarById(ctx, request.CarId)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = fmt.Sprintf("maaf anda gagal mendapatkan data mobil : %s", err.Error())
		resData.Status = "failed get data"
		return resData
	}
	bytes, err := ioutil.ReadFile(getDetail.ImageChar)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/gif":
		base64Encoding += "data:image/gif;base64,"
	}
	base64Encoding += base64.StdEncoding.EncodeToString(bytes)

	getDetail.ImageChar = base64Encoding

	resData.Code = http.StatusOK
	resData.Message = "sucses get data"
	resData.Status = "sucses"
	resData.Data = getDetail
	return resData
}

func (cs *CarServiceImpl) UpdateCarByCarId(ctx *gin.Context, request *requests.UpdateCarByCarIdRequest) responses.GenericResponse {

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
func (cs *CarServiceImpl) DeleteCarByCarId(ctx *gin.Context, request *requests.DeleteCarRequest) responses.GenericResponse {

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
