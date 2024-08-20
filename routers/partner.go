package routers

import (
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func PartnerRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", controllers.SeeAllPartner)
}
