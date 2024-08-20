package routers

import (
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func LocationRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", controllers.SeeAllLocations)
}
