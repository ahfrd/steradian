package entity

import (
	"steradian/entity/request"
	"steradian/entity/response"

	"github.com/gin-gonic/gin"
)

type ExampleService interface {
	ExampleService(ctx *gin.Context, request *request.ExampleRequest, uid string) (*response.LoginResponse, error)
}

type ExampleRepository interface {
	ExampleRepository(request *request.ExampleRequest) (*response.LoginResponse, error)
}
