package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// group all API routes
	v1 := router.Group("/api/v1/")

	v1.POST("/deposit", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Deposit",
		})
	})

	v1.POST("/withdraw", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Withdraw",
		})
	})

	v1.POST("/scheduleDeposit", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Schedule deposit",
		})
	})

	v1.POST("/scheduleWithdraw", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Schedule withdraw",
		})
	})

	v1.GET("/getBalance", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Show balance",
		})
	})

	v1.GET("/getVirtualBalance", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Show virtual balance",
		})
	})
}
