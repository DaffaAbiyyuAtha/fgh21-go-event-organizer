package routers

import (
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/controllers"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func UseRoutersWishlist(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AuthMiddleware())
	// routersGroup.POST("/", controllers.CreateTransaction)
	routerGroup.GET("/", controllers.ListAllWishlist)
	routerGroup.GET("/:id", controllers.DetailWishlist)
	routerGroup.POST("/:id", controllers.Createwishlist)
	routerGroup.DELETE("/:id", controllers.DeleteWishlistById)
}
