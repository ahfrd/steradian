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

// MockOrdersRepository is a mock implementation of repository.OrdersRepository for testing purposes

func TestCreateOrders_Success(t *testing.T) {
	// Prepare a mock request
	services := &services.OrdersServicesImpl{
		OrdersRepository: &mocks.MockOrdersRepository{
			CreateOrdersfn: func(ctx *gin.Context, request *requests.CreateOrdersRequests) error {
				return nil
			},
		},
	}
	requestBody := &requests.CreateOrdersRequests{
		CarId:           "3",
		OrderDate:       "2024-03-05",
		PickupDate:      "2024-03-07",
		DropoffDate:     "2024-03-09",
		PickupLocation:  "Pool",
		DropoffLocation: "Home",
	}
	router := gin.Default()
	router.POST("/insert-orders", func(ctx *gin.Context) {
		var req requests.CreateOrdersRequests
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res := services.CreateOrders(ctx, &req)

		ctx.JSON(http.StatusOK, res)
	})
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/insert-orders", nil)
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

func TestDetailOrders_Success(t *testing.T) {
	// Prepare a mock request
	services := &services.OrdersServicesImpl{
		OrdersRepository: &mocks.MockOrdersRepository{
			SelectDataOrdersByIdfn: func(ctx *gin.Context, ordersId string) (*responses.SelectListDataOrdersResponses, error) {
				return nil, nil
			},
		},
	}
	requestBody := &requests.GetDetailOrdersRequest{
		OrdersId: "3",
	}
	router := gin.Default()
	router.POST("/get-detail-orders/2", func(ctx *gin.Context) {
		var req requests.GetDetailOrdersRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res := services.GetDetailOrders(ctx, &req)

		ctx.JSON(http.StatusOK, res)
	})
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/get-detail-orders/2", nil)
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
func TestListOrders_Success(t *testing.T) {
	// Prepare a mock request
	services := &services.OrdersServicesImpl{
		OrdersRepository: &mocks.MockOrdersRepository{
			SelectListDataOrdersfn: func(ctx *gin.Context) ([]responses.SelectListDataOrdersResponses, error) {
				return nil, nil
			},
		},
	}
	requestBody := &requests.GetDetailOrdersRequest{
		OrdersId: "3",
	}
	router := gin.Default()
	router.POST("/get-list-orders", func(ctx *gin.Context) {

		res := services.GetListOrders(ctx)

		ctx.JSON(http.StatusOK, res)
	})
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/get-list-orders", nil)
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
func TestUpdateOrders_Success(t *testing.T) {
	// Prepare a mock request
	services := &services.OrdersServicesImpl{
		OrdersRepository: &mocks.MockOrdersRepository{
			UpdateOrdersByOrdersIdfn: func(ctx *gin.Context, request *requests.UpdateOrdersByOrdersIdRequest) error {
				return nil
			},
		},
	}
	requestBody := requests.UpdateOrdersByOrdersIdRequest{
		CarId:           "3",
		OrderDate:       "2024-03-05",
		PickupDate:      "2024-03-07",
		DropoffDate:     "2024-03-09",
		PickupLocation:  "Pool",
		DropoffLocation: "Home",
		OrdersId:        "2",
	}
	router := gin.Default()
	router.PUT("/update-orders", func(ctx *gin.Context) {

		res := services.UpdateOrdersByOrdersId(ctx, &requestBody)

		ctx.JSON(http.StatusOK, res)
	})
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodPut, "/update-orders", nil)
	body, _ := json.Marshal(requestBody)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	router.ServeHTTP(w, ctx.Request)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestDeleteOrders_Success(t *testing.T) {
	// Prepare a mock request
	services := &services.OrdersServicesImpl{
		OrdersRepository: &mocks.MockOrdersRepository{
			DeleteOrdersByOrdersIdfn: func(ctx *gin.Context, ordersId string) error {
				return nil
			},
		},
	}
	requestBody := requests.DeleteOrdersRequest{
		OrdersId: "3",
	}
	router := gin.Default()
	router.DELETE("/delete-orders/3", func(ctx *gin.Context) {

		res := services.DeleteOrdersByOrdersId(ctx, &requestBody)

		ctx.JSON(http.StatusOK, res)
	})
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest(http.MethodDelete, "/delete-orders/3", nil)
	body, _ := json.Marshal(requestBody)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	router.ServeHTTP(w, ctx.Request)

	assert.Equal(t, http.StatusOK, w.Code)

}
