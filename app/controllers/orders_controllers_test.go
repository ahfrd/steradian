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

func TestOrdersController_CreateOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controllers.OrdersControllers{}

	router := gin.Default()
	router.POST("/insert-Orders", controller.CreateOrders)

	t.Run("Successful Orders", func(t *testing.T) {
		bodyReq := requests.CreateOrdersRequests{
			CarId:           "2",
			OrderDate:       "2024-23-04",
			DropoffDate:     "2024-06-06",
			PickupLocation:  "l",
			PickupDate:      "2027-02-02",
			DropoffLocation: "k",
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/insert-Orders", bytes.NewReader(requestData))

		controller.OrdersService = &mocks.MockOrdersService{
			CreateOrdersfn: func(ctx *gin.Context, req *requests.CreateOrdersRequests) responses.GenericResponse {
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

	t.Run("Failed Orders - BindJSON Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/insert-Orders", bytes.NewReader([]byte("s-json")))

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("Failed Orders - Service Error", func(t *testing.T) {
		bodyReq := requests.CreateOrdersRequests{
			CarId:           "2",
			OrderDate:       "2024-23-04",
			DropoffDate:     "2024-06-06",
			PickupLocation:  "l",
			PickupDate:      "2027-02-02",
			DropoffLocation: "k",
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/insert-Orders", bytes.NewReader(requestData))

		controller.OrdersService = &mocks.MockOrdersService{
			CreateOrdersfn: func(ctx *gin.Context, req *requests.CreateOrdersRequests) responses.GenericResponse {

				return responses.GenericResponse{}
			},
		}
		// controller.OrdersService = &mock.MockOrdersService{
		// 	AddOrdersfn: func(ctx *gin.Context, req *request.AddOrdersRequest, requestId string) (*response.GeneralResponse, error) {
		// 		return nil, errors.New("service error")
		// 	},
		// }

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

	})
}

func TestOrdersController_DeleteOrdersByOrdersId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controllers.OrdersControllers{}

	router := gin.Default()
	router.DELETE("/delete-Orders/:OrdersId", controller.DeleteOrdersByOrdersId)

	t.Run("Successful delete", func(t *testing.T) {
		OrdersId := "1"
		bodyReq := requests.DeleteOrdersRequest{
			OrdersId: OrdersId,
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodDelete, "/delete-Orders/1", bytes.NewReader(requestData))

		controller.OrdersService = &mocks.MockOrdersService{
			DeleteOrdersByOrdersIdfn: func(ctx *gin.Context, req *requests.DeleteOrdersRequest) responses.GenericResponse {
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

	t.Run("Failed Orders - Service Error", func(t *testing.T) {
		OrdersId := "1"
		bodyReq := requests.DeleteOrdersRequest{
			OrdersId: OrdersId,
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodDelete, "/delete-Orders/1", bytes.NewReader(requestData))

		controller.OrdersService = &mocks.MockOrdersService{
			CreateOrdersfn: func(ctx *gin.Context, req *requests.CreateOrdersRequests) responses.GenericResponse {

				return responses.GenericResponse{}
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

	})
}

func TestOrdersController_GetListOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controllers.OrdersControllers{}

	router := gin.Default()
	router.GET("/get-list-Orders", controller.GetListOrders) // Assuming GetListOrders is a GET method

	t.Run("Successful GetListOrders", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/get-list-Orders", nil)

		// Mock OrdersService method to return a mock response
		controller.OrdersService = &mocks.MockOrdersService{
			GetListOrdersfn: func(ctx *gin.Context) responses.GenericResponse {
				// Mock response
				response := responses.GenericResponse{
					Code:    http.StatusOK,
					Status:  "success",
					Message: "List of Orderss",
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

	t.Run("Failed GetListOrders - Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/get-list-Orders", nil)

		// Mock OrdersService method to return an error response
		controller.OrdersService = &mocks.MockOrdersService{
			GetListOrdersfn: func(ctx *gin.Context) responses.GenericResponse {
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

func TestOrdersController_UpdateOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controllers.OrdersControllers{}

	router := gin.Default()
	router.PUT("/update-Orders", controller.UpdateOrdersByOrdersId)

	t.Run("Successful Orders", func(t *testing.T) {
		bodyReq := requests.UpdateOrdersByOrdersIdRequest{
			CarId:           "2",
			OrderDate:       "2024-23-04",
			DropoffDate:     "2024-06-06",
			PickupLocation:  "l",
			PickupDate:      "2027-02-02",
			DropoffLocation: "k",
			OrdersId:        "1",
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPut, "/update-Orders", bytes.NewReader(requestData))

		controller.OrdersService = &mocks.MockOrdersService{
			UpdateOrdersByOrdersIdfn: func(ctx *gin.Context, req *requests.UpdateOrdersByOrdersIdRequest) responses.GenericResponse {
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

	t.Run("Failed Orders - BindJSON Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPut, "/update-Orders", bytes.NewReader([]byte("s-json")))

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("Failed Orders - Service Error", func(t *testing.T) {
		bodyReq := requests.UpdateOrdersByOrdersIdRequest{
			CarId:           "2",
			OrderDate:       "2024-23-04",
			DropoffDate:     "2024-06-06",
			PickupLocation:  "l",
			PickupDate:      "2027-02-02",
			DropoffLocation: "k",
			OrdersId:        "1",
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPut, "/update-Orders", bytes.NewReader(requestData))

		controller.OrdersService = &mocks.MockOrdersService{
			UpdateOrdersByOrdersIdfn: func(ctx *gin.Context, req *requests.UpdateOrdersByOrdersIdRequest) responses.GenericResponse {

				return responses.GenericResponse{}
			},
		}
		// controller.OrdersService = &mock.MockOrdersService{
		// 	AddOrdersfn: func(ctx *gin.Context, req *request.AddOrdersRequest, requestId string) (*response.GeneralResponse, error) {
		// 		return nil, errors.New("service error")
		// 	},
		// }

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

	})
}

func TestOrdersController_GetDetailOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controllers.OrdersControllers{}

	router := gin.Default()
	router.GET("/get-detail-Orders/:OrdersId", controller.GetDetailOrders)

	t.Run("Successful Orders", func(t *testing.T) {
		bodyReq := requests.GetDetailOrdersRequest{}
		OrdersId := "123"
		bodyReq.OrdersId = OrdersId
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/get-detail-Orders/123", bytes.NewReader(requestData))

		controller.OrdersService = &mocks.MockOrdersService{
			CreateOrdersfn: func(ctx *gin.Context, req *requests.CreateOrdersRequests) responses.GenericResponse {
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

	t.Run("Failed Orders - BindJSON Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/get-detail-Orders/123", bytes.NewReader([]byte("s-json")))

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("Failed Orders - Service Error", func(t *testing.T) {
		bodyReq := requests.GetDetailOrdersRequest{}
		OrdersId := "123"
		bodyReq.OrdersId = OrdersId
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/get-detail-Orders/123", bytes.NewReader(requestData))

		controller.OrdersService = &mocks.MockOrdersService{
			CreateOrdersfn: func(ctx *gin.Context, req *requests.CreateOrdersRequests) responses.GenericResponse {

				return responses.GenericResponse{}
			},
		}
		// controller.OrdersService = &mock.MockOrdersService{
		// 	AddOrdersfn: func(ctx *gin.Context, req *request.AddOrdersRequest, requestId string) (*response.GeneralResponse, error) {
		// 		return nil, errors.New("service error")
		// 	},
		// }

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

	})
}
