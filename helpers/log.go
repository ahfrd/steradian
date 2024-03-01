package helpers

import (
	"fmt"
	"log"
	"steradian/helpers/constant"
	"time"

	"github.com/gin-gonic/gin"
)

func LogRequest(ctx *gin.Context, requestString string) {
	currentTime := time.Now()
	logMessage := fmt.Sprint(constant.LogStartPrefix +
		constant.LogMethodPrefix + ctx.Request.Method +
		constant.LogPathPrefix + ctx.FullPath() +
		constant.LogIPPrefix + ctx.ClientIP() +
		constant.LogTimePrefix + currentTime.Format(constant.DateYYYYMMDDHHIISS) +
		constant.LogRequestPrefix + requestString)
	log.Println(logMessage)
}

func LogResponse(ctx *gin.Context, responseString string) {
	currentTime := time.Now()
	logMessage := fmt.Sprint(constant.LogStopPrefix +
		constant.LogMethodPrefix + ctx.Request.Method +
		constant.LogPathPrefix + ctx.FullPath() +
		constant.LogIPPrefix + ctx.ClientIP() +
		constant.LogTimePrefix + currentTime.Format(constant.DateYYYYMMDDHHIISS) +
		constant.LogResponsePrefix + responseString)
	log.Println(logMessage)
}

func LogError(ctx *gin.Context, responseString string) {
	currentTime := time.Now()
	logMessage := fmt.Sprint(constant.LogStopPrefix +
		constant.LogMethodPrefix + ctx.Request.Method +
		constant.LogPathPrefix + ctx.FullPath() +
		constant.LogIPPrefix + ctx.ClientIP() +
		constant.LogTimePrefix + currentTime.Format(constant.DateYYYYMMDDHHIISS) +
		constant.LogErrorPrefix + responseString)
	log.Println(logMessage)

}
