package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Id int `json:"id"`
	Name string	`json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"-" form:"password"`
}
type Server struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Result interface{} `json:"result,omitempty"`
}


func main() {
	r := gin.Default()
	r.Use(corsMiddleware())
	data := []Users{
		{
			Id : 1, 
			Name : "Fazz",
			Email : "fazz@mail.com",
			Password: "admin1234",
		},
	}
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, Server{
			Success: true,
			Message: "OK",
			Result: data,
		})
	})
	r.POST("/users", func(c *gin.Context) {
		user := Users{}
		c.Bind(&user)
		user.Id = len(data) + 1
		data = append(data, user)
		c.JSON(http.StatusOK, Server{
			Success: true,
			Message: "Create User Success",
			Result: user,
		})
	})
	r.POST("/auth/login", func(c *gin.Context) {
		user := Users{}
		c.Bind(&user)
		email := user.Email
		password := user.Password
		dataResults := true
        if dataResults{
            for dataResults {
                for i := 0; i <len(data); i++ {
                    resultsEmail := data[i].Email
                    resultsPassword := data[i].Password
                    if email == resultsEmail && password == resultsPassword {
                        c.JSON(http.StatusOK, Server{
                            Success : true,
                            Message: "Login succes",
                        })
                    }
                }

                dataResults = false
            }
			c.JSON(http.StatusUnauthorized, Server{
			Success : false,
			Message: "salah",
			})
        } 
	})
	r.PATCH("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		selected := -1

		for index, item := range data {
			if item.Id == id {
				selected = index
			}
		}

		if selected != -1 {
			form := Users{}
			c.Bind(&form)
			data[selected].Name = form.Name
			data[selected].Email = form.Email
			data[selected].Password = form.Password
			c.JSON(http.StatusOK, Server{
				Success: true,
				Message: "Update Success",
				Result: data[selected],
			})
		}else{
			c.JSON(http.StatusNotFound, Server{
				Success: false,
				Message: "User Not Found",
			})
		}

	})
	r.DELETE("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		selected := -1

		for index, item := range data {
			if item.Id == id {
				selected = index
			}
		}

		if selected != -1 {
			form := Users{}
			c.Bind(&form)
			showData := data[selected]
			data = append(data[:selected],data[selected+1:]...)
			c.JSON(http.StatusOK, Server{
				Success: true,
				Message: "Update Success",
				Result: showData,
			})
		}else{
			c.JSON(http.StatusNotFound, Server{
				Success: false,
				Message: "User Not Found",
			})
		}

	})
	// router.Use(cors.Default())
  	// router.Run()
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