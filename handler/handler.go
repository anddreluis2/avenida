package handler

import "github.com/gin-gonic/gin"

func Deposit(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Deposit",
	})
}

func Withdraw(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Withdraw",
	})
}

func ScheduleDeposit(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Schedule deposit",
	})
}

func ScheduleWithdraw(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Schedule withdraw",
	})
}

func GetBalance(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Show balance",
	})
}

func GetVirtualBalance(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Show virtual balance",
	})
}
