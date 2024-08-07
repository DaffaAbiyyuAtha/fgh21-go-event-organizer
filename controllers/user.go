package controllers

import (
	"net/http"
	"strconv"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func SeeAllUsers(c *gin.Context) {
	data := models.GetAllUsers()
	c.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Ok",
		Result:  data,
	})
}

func SeeOneUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.GetOneUserById(id)

	if data.Id != 0 {
		c.JSON(http.StatusOK, lib.Server{
			Success: true,
			Message: "User Found",
			Result:  data,
		})
	} else {
		c.JSON(http.StatusNotFound, lib.Server{
			Success: false,
			Message: "User Not Found",
		})
	}
}

func CreateUser(c *gin.Context) {
	user := models.Users{}
	c.Bind(&user)

	data := models.CreateUser(user)
	c.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Create User Success!",
		Result:  data,
	})
}

func Login(c *gin.Context) {
	user := models.Users{}
	c.Bind(&user)
	email := user.Email
	password := user.Password
	dataResults := true
	if dataResults {
		for dataResults {
			for i := 0; i < len(data); i++ {
				resultsEmail := data[i].Email
				resultsPassword := data[i].Password
				if email == resultsEmail && password == resultsPassword {
					c.JSON(http.StatusOK, lib.Server{
						Success: true,
						Message: "Login succes",
					})
				}
			}

			dataResults = false
		}
		c.JSON(http.StatusUnauthorized, lib.Server{
			Success: false,
			Message: "Not Unauthorized",
		})
	}
}

func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	updatedData := models.Users{}
	c.Bind(&updatedData)

	data := models.UpdateDataById(updatedData, id)

	if data.Id != 0 {
		c.JSON(http.StatusOK, lib.Server{
			Success: true,
			Message: "Update Data Success!",
			Result:  data,
		})
	} else {
		c.JSON(http.StatusNotFound, lib.Server{
			Success: false,
			Message: "User Not Found",
		})
	}
}

func DeleteUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data := models.DeleteDataById(id)

	if data.Id != 0 {
		c.JSON(http.StatusOK, lib.Server{
			Success: true,
			Message: "Delete Data Success!",
			Result:  data,
		})
	} else {
		c.JSON(http.StatusNotFound, lib.Server{
			Success: false,
			Message: "Data Not Found",
		})
	}

}
