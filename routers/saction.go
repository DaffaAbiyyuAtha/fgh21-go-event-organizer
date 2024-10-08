package routers

import (
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func SactionsRouter(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/", controllers.CreateSaction)
}
