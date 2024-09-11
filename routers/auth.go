package routers

import (
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(RouterGroup *gin.RouterGroup) {
	RouterGroup.POST("/login", controllers.AuthLogin)
	RouterGroup.POST("/register", controllers.AuthProfile)
}
