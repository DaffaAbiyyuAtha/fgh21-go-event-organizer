package main

import (
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())

	routers.RouterCombine(r)
	// data := []Users{
	// 	{
	// 		Id:       1,
	// 		Name:     "Fazz",
	// 		Email:    "fazz@mail.com",
	// 		Password: "admin1234",
	// 	},
	// }
	// r.GET("/users", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, Server{
	// 		Success: true,
	// 		Message: "OK",
	// 		Result:  data,
	// 	})
	// })
	// r.GET("/users/:id", func(c *gin.Context) {
	// 	id, _ := strconv.Atoi(c.Param("id"))

	// 	selected := -1

	// 	for index, item := range data {
	// 		if item.Id == id {
	// 			selected = index
	// 		}
	// 	}

	// 	if selected != -1 {
	// 		c.JSON(http.StatusOK, Server{
	// 			Success: true,
	// 			Message: "User Found",
	// 			Result:  data[selected],
	// 		})
	// 	} else {
	// 		c.JSON(http.StatusNotFound, Server{
	// 			Success: false,
	// 			Message: "User Not Found",
	// 		})
	// 	}

	// })
	// r.POST("/users", func(c *gin.Context) {
	// 	user := Users{}
	// 	datas := c.Bind(&user)

	// 	condition := true

	// 	result := 0
	// 	for _, i := range data {
	// 		result = i.Id
	// 	}
	// 	user.Id = result + 1
	// 	for _, i := range data {
	// 		if i.Email == user.Email && i.Password == user.Password {
	// 			condition = false
	// 		}
	// 	}
	// 	if datas != nil {
	// 		c.JSON(http.StatusBadRequest, Server{
	// 			Success: false,
	// 			Message: "Please Fill Form",
	// 		})
	// 	} else {
	// 		if condition {
	// 			data = append(data, user)
	// 			c.JSON(http.StatusOK, Server{
	// 				Success: true,
	// 				Message: "Create User Success",
	// 				Result:  user,
	// 			})
	// 		} else {
	// 			c.JSON(http.StatusBadRequest, Server{
	// 				Success: false,
	// 				Message: "Email Already Use",
	// 			})
	// 		}
	// 	}
	// })
	// r.POST("/auth/login", func(c *gin.Context) {
	// 	user := Users{}
	// 	c.Bind(&user)
	// 	email := user.Email
	// 	password := user.Password
	// 	dataResults := true
	// 	if dataResults {
	// 		for dataResults {
	// 			for i := 0; i < len(data); i++ {
	// 				resultsEmail := data[i].Email
	// 				resultsPassword := data[i].Password
	// 				if email == resultsEmail && password == resultsPassword {
	// 					c.JSON(http.StatusOK, Server{
	// 						Success: true,
	// 						Message: "Login succes",
	// 					})
	// 				}
	// 			}

	// 			dataResults = false
	// 		}
	// 		c.JSON(http.StatusUnauthorized, Server{
	// 			Success: false,
	// 			Message: "Not Unauthorized",
	// 		})
	// 	}
	// })
	// r.PATCH("/users/:id", func(c *gin.Context) {
	// 	id, _ := strconv.Atoi(c.Param("id"))

	// 	updatedData := Users{}
	// 	err := c.Bind(&updatedData)
	// 	conditional := true
	// 	for _, i := range data {
	// 		if i.Email == updatedData.Email {
	// 			conditional = false
	// 		}
	// 	}
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, Server{
	// 			Success: false,
	// 			Message: "Data Not Found",
	// 		})
	// 	} else {
	// 		if conditional {
	// 			for i, updated := range data {
	// 				if updated.Id == id {
	// 					data[i].Name = updatedData.Name
	// 					data[i].Email = updatedData.Email
	// 					data[i].Password = updatedData.Password
	// 					c.JSON(http.StatusOK, Server{
	// 						Success: true,
	// 						Message: "User Update",
	// 						Result:  data,
	// 					})
	// 					return
	// 				}
	// 			}

	// 		} else {
	// 			c.JSON(http.StatusNotFound, Server{
	// 				Success: false,
	// 				Message: "User Not Found",
	// 			})
	// 		}
	// 	}

	// 	c.JSON(http.StatusNotFound, Server{
	// 		Success: false,
	// 		Message: "Users Not Found",
	// 	})
	// })
	// r.DELETE("/users/:id", func(c *gin.Context) {
	// 	id, _ := strconv.Atoi(c.Param("id"))

	// 	selected := -1

	// 	for index, item := range data {
	// 		if item.Id == id {
	// 			selected = index
	// 		}
	// 	}

	// 	if selected != -1 {
	// 		form := Users{}
	// 		c.Bind(&form)
	// 		showData := data[selected]
	// 		data = append(data[:selected], data[selected+1:]...)
	// 		c.JSON(http.StatusOK, Server{
	// 			Success: true,
	// 			Message: "Update Success",
	// 			Result:  showData,
	// 		})
	// 	} else {
	// 		c.JSON(http.StatusNotFound, Server{
	// 			Success: false,
	// 			Message: "User Not Found",
	// 		})
	// 	}

	// })
	r.Run("localhost:8080")
}
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
