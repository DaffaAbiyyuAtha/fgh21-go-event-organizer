package routers

import (
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func NationalitiesRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", controllers.ListAllNationalities)
}
