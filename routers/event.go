package routers

import (
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/controllers"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func EventRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", controllers.SeeAllEvent)
	routerGroup.GET("/filter", controllers.ListAllFilterEvents)
	routerGroup.GET("/section/:id", controllers.SeeOneEventByEventId)
	routerGroup.GET("/:id", controllers.SeeOneEventById)
	// routerGroup.POST("/auth/login", controllers.Login)
	routerGroup.PATCH("/:id", controllers.EditEvent)
	routerGroup.DELETE("/:id", controllers.DeleteEventById)
	routerGroup.GET("/payment_method", controllers.ListAllPaymentMethod)
	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.POST("/update", controllers.CreateEvent)
	routerGroup.GET("/see_one_event", controllers.SeeOneEventByUserId)
}
