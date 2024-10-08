package routers

import (
	"github.com/gin-gonic/gin"
)

func RouterCombine(r *gin.Engine) {
	UserRouter(r.Group("/users"))
	ProfileRouter(r.Group("/profile"))
	AuthRouter(r.Group("/auth"))
	EventRouter(r.Group("/events"))
	CategoryRouter(r.Group("/categories"))
	TransactionsRouter(r.Group("/transactions"))
	LocationRouter(r.Group("/locations"))
	PartnerRouter(r.Group("/partner"))
	UseRoutersWishlist(r.Group("/wishlist"))
	NationalitiesRouter(r.Group("/nationalities"))
	SactionsRouter(r.Group("/saction"))
}
