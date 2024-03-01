package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"steradian/app/controllers"
	"steradian/app/models/mocks"
	"steradian/app/models/requests"
	"steradian/app/models/responses"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCarController_CreateCar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controllers.CarController{}

	router := gin.Default()
	router.POST("/insert-car", controller.CreateCar)

	t.Run("Successful Car", func(t *testing.T) {
		bodyReq := requests.CreateCarRequests{
			CarName:   "camry",
			DayRate:   200,
			MonthRate: 100,
			ImageCar:  "img",
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/insert-car", bytes.NewReader(requestData))

		controller.CarService = &mocks.MockCarService{
			CreateCarfn: func(ctx *gin.Context, req *requests.CreateCarRequests) responses.GenericResponse {
				response := &responses.GenericResponse{
					Code:    http.StatusOK,
					Status:  "succses",
					Message: "sucses",
				}
				return *response
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

		var response responses.GenericResponse
		_ = json.Unmarshal(w.Body.Bytes(), &response)

	})

	t.Run("Failed Car - BindJSON Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/insert-car", bytes.NewReader([]byte("s-json")))

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("Failed Car - Service Error", func(t *testing.T) {
		bodyReq := requests.CreateCarRequests{
			CarName:   "camry",
			DayRate:   200,
			MonthRate: 100,
			ImageCar:  "img",
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/insert-car", bytes.NewReader(requestData))

		controller.CarService = &mocks.MockCarService{
			CreateCarfn: func(ctx *gin.Context, req *requests.CreateCarRequests) responses.GenericResponse {

				return responses.GenericResponse{}
			},
		}
		// controller.CarService = &mock.MockCarService{
		// 	AddCarfn: func(ctx *gin.Context, req *request.AddCarRequest, requestId string) (*response.GeneralResponse, error) {
		// 		return nil, errors.New("service error")
		// 	},
		// }

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

	})
}

func TestCarController_DeleteCarByCarId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controllers.CarController{}

	router := gin.Default()
	router.DELETE("/delete-car/:carId", controller.DeleteCarByCarId)

	t.Run("Successful delete", func(t *testing.T) {
		carId := "1"
		bodyReq := requests.DeleteCarRequest{
			CarId: carId,
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodDelete, "/delete-car/1", bytes.NewReader(requestData))

		controller.CarService = &mocks.MockCarService{
			DeleteCarByCarIdfn: func(ctx *gin.Context, req *requests.DeleteCarRequest) responses.GenericResponse {
				response := &responses.GenericResponse{
					Code:    http.StatusOK,
					Status:  "succses",
					Message: "sucses",
				}
				return *response
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

		var response responses.GenericResponse
		_ = json.Unmarshal(w.Body.Bytes(), &response)

	})

	t.Run("Failed Car - Service Error", func(t *testing.T) {
		carId := "1"
		bodyReq := requests.DeleteCarRequest{
			CarId: carId,
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodDelete, "/delete-car/1", bytes.NewReader(requestData))

		controller.CarService = &mocks.MockCarService{
			CreateCarfn: func(ctx *gin.Context, req *requests.CreateCarRequests) responses.GenericResponse {

				return responses.GenericResponse{}
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

	})
}

func TestCarController_GetListCar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controllers.CarController{}

	router := gin.Default()
	router.GET("/get-list-car", controller.GetListCar) // Assuming GetListCar is a GET method

	t.Run("Successful GetListCar", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/get-list-car", nil)

		// Mock CarService method to return a mock response
		controller.CarService = &mocks.MockCarService{
			GetListCarfn: func(ctx *gin.Context) responses.GenericResponse {
				// Mock response
				response := responses.GenericResponse{
					Code:    http.StatusOK,
					Status:  "success",
					Message: "List of cars",
				}
				return response
			},
		}

		// Serve HTTP request to the mock recorder
		router.ServeHTTP(w, ctx.Request)

		// Check response status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Parse response body and perform assertions
		var response responses.GenericResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
		// Add more assertions as per your response structure
	})

	t.Run("Failed GetListCar - Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/get-list-car", nil)

		// Mock CarService method to return an error response
		controller.CarService = &mocks.MockCarService{
			GetListCarfn: func(ctx *gin.Context) responses.GenericResponse {
				// Mock error response
				response := responses.GenericResponse{
					Code:    http.StatusInternalServerError,
					Status:  "error",
					Message: "Internal server error",
				}
				return response
			},
		}

		// Serve HTTP request to the mock recorder
		router.ServeHTTP(w, ctx.Request)

		// Check response status code
		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response responses.GenericResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestCarController_UpdateCar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controllers.CarController{}

	router := gin.Default()
	router.PUT("/update-car", controller.UpdateCarByCarId)

	t.Run("Successful Car", func(t *testing.T) {
		bodyReq := requests.UpdateCarByCarIdRequest{
			CarName:   "camry",
			DayRate:   200,
			MonthRate: 100,
			ImageCar:  "img",
			CarId:     "1",
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPut, "/update-car", bytes.NewReader(requestData))

		controller.CarService = &mocks.MockCarService{
			UpdateCarByCarIdfn: func(ctx *gin.Context, req *requests.UpdateCarByCarIdRequest) responses.GenericResponse {
				response := &responses.GenericResponse{
					Code:    http.StatusOK,
					Status:  "succses",
					Message: "sucses",
				}
				return *response
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

		var response responses.GenericResponse
		_ = json.Unmarshal(w.Body.Bytes(), &response)

	})

	t.Run("Failed Car - BindJSON Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPut, "/update-car", bytes.NewReader([]byte("s-json")))

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("Failed Car - Service Error", func(t *testing.T) {
		bodyReq := requests.CreateCarRequests{
			CarName:   "camry",
			DayRate:   200,
			MonthRate: 100,
			ImageCar:  "img",
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPut, "/update-car", bytes.NewReader(requestData))

		controller.CarService = &mocks.MockCarService{
			UpdateCarByCarIdfn: func(ctx *gin.Context, req *requests.UpdateCarByCarIdRequest) responses.GenericResponse {

				return responses.GenericResponse{}
			},
		}
		// controller.CarService = &mock.MockCarService{
		// 	AddCarfn: func(ctx *gin.Context, req *request.AddCarRequest, requestId string) (*response.GeneralResponse, error) {
		// 		return nil, errors.New("service error")
		// 	},
		// }

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

	})
}

func TestCarController_GetDetailCar(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controllers.CarController{}

	router := gin.Default()
	router.GET("/get-detail-car/:carId", controller.GetDetailCar)

	t.Run("Successful Car", func(t *testing.T) {
		bodyReq := requests.GetDetailCarRequest{}
		carId := "123"
		bodyReq.CarId = carId
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/get-detail-car/123", bytes.NewReader(requestData))

		controller.CarService = &mocks.MockCarService{
			CreateCarfn: func(ctx *gin.Context, req *requests.CreateCarRequests) responses.GenericResponse {
				response := &responses.GenericResponse{
					Code:    http.StatusOK,
					Status:  "succses",
					Message: "sucses",
				}
				return *response
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

		var response responses.GenericResponse
		_ = json.Unmarshal(w.Body.Bytes(), &response)

	})

	t.Run("Failed Car - BindJSON Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/get-detail-car/123", bytes.NewReader([]byte("s-json")))

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("Failed Car - Service Error", func(t *testing.T) {
		bodyReq := requests.GetDetailCarRequest{}
		carId := "123"
		bodyReq.CarId = carId
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/get-detail-car/123", bytes.NewReader(requestData))

		controller.CarService = &mocks.MockCarService{
			CreateCarfn: func(ctx *gin.Context, req *requests.CreateCarRequests) responses.GenericResponse {

				return responses.GenericResponse{}
			},
		}
		// controller.CarService = &mock.MockCarService{
		// 	AddCarfn: func(ctx *gin.Context, req *request.AddCarRequest, requestId string) (*response.GeneralResponse, error) {
		// 		return nil, errors.New("service error")
		// 	},
		// }

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

	})
}
