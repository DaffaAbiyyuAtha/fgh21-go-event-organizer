package routers

import (
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/controllers"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/", controllers.CreateUser)
	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("/", controllers.SeeAllUsers)
	routerGroup.PATCH("password/", controllers.UpdatePassword)
	routerGroup.GET("/:id", controllers.SeeOneUserById)
	// routerGroup.POST("/auth/login", controllers.Login)
	// routerGroup.PATCH("/:id", controllers.EditUser)
	routerGroup.DELETE("/:id", controllers.DeleteUserById)
	routerGroup.PATCH("/update", controllers.UpdateUser)
}
