package middlewares

import (
	"net/http"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/gin-gonic/gin"
)

func tokenFailed(cc *gin.Context) {
	if e := recover(); e != nil {
		cc.JSON(http.StatusUnauthorized, lib.Server{
			Success: false,
			Message: "Unauthorized",
		})
		cc.Abort()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer tokenFailed(ctx)
		token := ctx.GetHeader("Authorization")[7:]
		isValidated, userId := lib.ValidateToken(token)
		if isValidated {
			ctx.Set("userId", userId)
			ctx.Next()
		} else {
			panic("Error: token invalid")
		}
	}
}
