package routers

import (
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/controllers"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func TransactionsRouter(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("/:id", controllers.FindTransactionByUserId)
	routerGroup.POST("/", controllers.CreateTransaction)
	routerGroup.GET("/", controllers.ListDetailsTransactions)
	routerGroup.GET("/payment/", controllers.ListProductById)
}
