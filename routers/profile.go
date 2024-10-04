package routers

import (
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/controllers"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func ProfileRouter(routerGroup *gin.RouterGroup) {
	routerGroup.DELETE("/:id", controllers.DeleteProfileById)
	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.POST("/", controllers.CreateProfiles)
	// routerGroup.GET("/", controllers.ListAllProfile)
	routerGroup.GET("/", controllers.SeeOneProfileByUserId)
	routerGroup.PATCH("/update", controllers.UpdateProfile)
	routerGroup.PATCH("/picture", controllers.UpdateProfilePicture)
}
