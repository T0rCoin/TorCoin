package router

import (
	"TorCoin/app/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	route := gin.Default()

	route.GET("/", service.Root)

	route.POST("/operator/wallets/:id/addresses", service.NewAddress)

	route.GET("/operator/:address/balance",service.GetBalance)

	return route

}

