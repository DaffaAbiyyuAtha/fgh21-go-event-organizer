package routers

import (
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", controllers.SeeAllUsers)
	routerGroup.GET("/:id", controllers.SeeOneUserById)
	routerGroup.POST("/", controllers.CreateUser)
	routerGroup.POST("/auth/login", controllers.Login)
	routerGroup.PATCH("/:id", controllers.EditUser)
	routerGroup.DELETE("/:id", controllers.DeleteUserById)
}
