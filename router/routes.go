package router

import (
	"github.com/anddreluis2/avenida/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// group all API routes
	v1 := router.Group("/api/v1/")

	v1.POST("/deposit", handler.Deposit)
	v1.POST("/withdraw", handler.Withdraw)
	v1.POST("/scheduleDeposit", handler.ScheduleDeposit)
	v1.POST("/scheduleWithdraw", handler.ScheduleWithdraw)
	v1.GET("/getBalance", handler.GetBalance)
	v1.GET("/getVirtualBalance", handler.GetVirtualBalance)
}
