package routers

import (
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func CategoryRouter(routerGroup *gin.RouterGroup) {
	// routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.POST("/", controllers.CreateCategory)
	routerGroup.GET("/", controllers.SeeAllCategories)
	routerGroup.GET("/:id", controllers.SeeOneCategoryById)
	routerGroup.PATCH("/:id", controllers.EditCategory)
	routerGroup.DELETE("/:id", controllers.DeleteCategoryById)
}
