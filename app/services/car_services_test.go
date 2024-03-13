package services_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"steradian/app/models/mocks"
	"steradian/app/models/requests"
	"steradian/app/models/responses"
	"steradian/app/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockCarRepository is a mock implementation of repository.CarRepository for testing purposes

func TestCreateCar_Success(t *testing.T) {
	// Prepare a mock request
	services := &services.CarServiceImpl{
		CarRepository: &mocks.MockCarRepository{
			CreateCarfn: func(ctx *gin.Context, request *requests.CreateCarRequests) (int, error) {
				return 8, nil
			},
		},
	}
	requestBody := &requests.CreateCarRequests{
		DayRate:   20000,
		MonthRate: 100000,
		ImageCar:  "https://images.drive.com.au/caradvice/image/private/c_fill,f_auto,g_auto,h_1080,q_auto:eco,w_1920/v1/30f70253683fac6506730f0efe3ad514",
		CarName:   "Dodge",
	}
	router := gin.Default()
	router.POST("/insert-car", func(ctx *gin.Context) {
		var req requests.CreateCarRequests
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res := services.CreateCar(ctx, &req)

		ctx.JSON(http.StatusOK, res)
	})
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/insert-car", nil)
	body, _ := json.Marshal(requestBody)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	router.ServeHTTP(w, ctx.Request)

	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	expectedRes := responses.GenericResponse{
		Code:    200,
		Status:  "sucses",
		Message: "succses insert",
	}
	actualRes := responses.GenericResponse{}
	_ = json.Unmarshal(w.Body.Bytes(), &actualRes)
	assert.Equal(t, expectedRes, actualRes)
}

func TestDetailCar_Success(t *testing.T) {
	// Prepare a mock request
	services := &services.CarServiceImpl{
		CarRepository: &mocks.MockCarRepository{
			SelectDataCarByIdfn: func(ctx *gin.Context, CarId string) (*responses.SelectListDataCarResponses, error) {
				return &responses.SelectListDataCarResponses{
					CarId:     "3",
					CarName:   "Mustang",
					DayRate:   120000,
					MonthRate: 300000,
					ImageChar: "app/upload/7.jpeg",
				}, nil
			},
		},
	}
	requestBody := &requests.GetDetailCarRequest{
		CarId: "3",
	}
	router := gin.Default()
	router.GET("/get-detail-car/3", func(ctx *gin.Context) {
		var req requests.GetDetailCarRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res := services.GetDetailCar(ctx, &req)

		ctx.JSON(http.StatusOK, res)
	})
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/get-detail-car/3", nil)
	body, _ := json.Marshal(requestBody)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	router.ServeHTTP(w, ctx.Request)

	assert.Equal(t, http.StatusOK, w.Code)

}
func TestListCar_Success(t *testing.T) {
	// Prepare a mock request
	services := &services.CarServiceImpl{
		CarRepository: &mocks.MockCarRepository{
			SelectListDataCarfn: func(ctx *gin.Context) ([]responses.SelectListDataCarResponses, error) {
				return nil, nil
			},
		},
	}
	requestBody := &requests.GetDetailCarRequest{
		CarId: "3",
	}
	router := gin.Default()
	router.GET("/get-list-car", func(ctx *gin.Context) {

		res := services.GetListCar(ctx)

		ctx.JSON(http.StatusOK, res)
	})
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/get-list-car", nil)
	body, _ := json.Marshal(requestBody)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	router.ServeHTTP(w, ctx.Request)

	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	expectedRes := responses.GenericResponse{
		Code:    200,
		Status:  "sucses",
		Message: "sucses get data",
	}
	actualRes := responses.GenericResponse{}
	_ = json.Unmarshal(w.Body.Bytes(), &actualRes)
	assert.Equal(t, expectedRes, actualRes)
}
func TestUpdateCar_Success(t *testing.T) {
	// Prepare a mock request
	services := &services.CarServiceImpl{
		CarRepository: &mocks.MockCarRepository{
			UpdateCarByCarIdfn: func(ctx *gin.Context, request *requests.UpdateCarByCarIdRequest) error {
				return nil
			},
		},
	}
	requestBody := requests.UpdateCarByCarIdRequest{
		CarId:     "3",
		DayRate:   20000,
		MonthRate: 100000,
		ImageCar:  "https://images.drive.com.au/caradvice/image/private/c_fill,f_auto,g_auto,h_1080,q_auto:eco,w_1920/v1/30f70253683fac6506730f0efe3ad514",
		CarName:   "Dodge",
	}
	router := gin.Default()
	router.PUT("/update-car", func(ctx *gin.Context) {

		res := services.UpdateCarByCarId(ctx, &requestBody)

		ctx.JSON(http.StatusOK, res)
	})
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodPut, "/update-car", nil)
	body, _ := json.Marshal(requestBody)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	router.ServeHTTP(w, ctx.Request)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestDeleteCar_Success(t *testing.T) {
	// Prepare a mock request
	services := &services.CarServiceImpl{
		CarRepository: &mocks.MockCarRepository{
			DeleteCarByCarIdfn: func(ctx *gin.Context, CarId string) error {
				return nil
			},
		},
	}
	requestBody := requests.DeleteCarRequest{
		CarId: "3",
	}
	router := gin.Default()
	router.DELETE("/delete-Car/3", func(ctx *gin.Context) {

		res := services.DeleteCarByCarId(ctx, &requestBody)

		ctx.JSON(http.StatusOK, res)
	})
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodDelete, "/delete-Car/3", nil)
	body, _ := json.Marshal(requestBody)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	router.ServeHTTP(w, ctx.Request)

	assert.Equal(t, http.StatusOK, w.Code)

}
