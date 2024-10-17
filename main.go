package main

import (
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/picture", "./picture")
	// r.Use(corsMiddleware())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"Origin", "Content-Type", "Authorization", "Content-Length",
	}
	r.Use(cors.New(corsConfig))

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS", "DELETE", "HEAD"},
	// 	AllowHeaders:     []string{"*"},
	// 	AllowCredentials: true,
	// }))

	routers.RouterCombine(r)

	r.Run("0.0.0.0:8080")
}

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return

// 		}
// 		c.Next()
// 	}

// }
