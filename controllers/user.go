// package controllers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
// 	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/models"
// 	"github.com/gin-gonic/gin"
// )

// func SeeAllUsers(c *gin.Context) {
// 	data := models.GetAllUsers()
// 	c.JSON(http.StatusOK, lib.Server{
// 		Success: true,
// 		Message: "Ok",
// 		Result:  data,
// 	})
// }

// func SeeOneUserById(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	data := models.GetOneUserById(id)

// 	if data.Id != 0 {
// 		c.JSON(http.StatusOK, lib.Server{
// 			Success: true,
// 			Message: "User Found",
// 			Result:  data,
// 		})
// 	} else {
// 		c.JSON(http.StatusNotFound, lib.Server{
// 			Success: false,
// 			Message: "User Not Found",
// 		})
// 	}
// }

// func CreateUser(c *gin.Context) {
// 	user := models.Users{}
// 	c.Bind(&user)

// 	data := models.CreateUser(user)
// 	c.JSON(http.StatusOK, lib.Server{
// 		Success: true,
// 		Message: "Create User Success!",
// 		Result:  data,
// 	})
// }

// func Login(c *gin.Context) {
// 	user := models.Users{}
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
// 					c.JSON(http.StatusOK, lib.Server{
// 						Success: true,
// 						Message: "Login succes",
// 					})
// 				}
// 			}

// 			dataResults = false
// 		}
// 		c.JSON(http.StatusUnauthorized, lib.Server{
// 			Success: false,
// 			Message: "Not Unauthorized",
// 		})
// 	}
// }

// func EditUser(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	updatedData := models.Users{}
// 	c.Bind(&updatedData)

// 	data := models.UpdateDataById(updatedData, id)

// 	if data.Id != 0 {
// 		c.JSON(http.StatusOK, lib.Server{
// 			Success: true,
// 			Message: "Update Data Success!",
// 			Result:  data,
// 		})
// 	} else {
// 		c.JSON(http.StatusNotFound, lib.Server{
// 			Success: false,
// 			Message: "User Not Found",
// 		})
// 	}
// }

// func DeleteUserById(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	data := models.DeleteDataById(id)

// 	if data.Id != 0 {
// 		c.JSON(http.StatusOK, lib.Server{
// 			Success: true,
// 			Message: "Delete Data Success!",
// 			Result:  data,
// 		})
// 	} else {
// 		c.JSON(http.StatusNotFound, lib.Server{
// 			Success: false,
// 			Message: "Data Not Found",
// 		})
// 	}

// }

package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func SeeAllUsers(ctx *gin.Context) {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	if limit < 1 {
		limit = 5
	}

	if page < 1 {
		page = 1
	}
	// if page > 1 {
	// 	page = (page - 1) * limit
	// }

	result, count := models.FindAllUsers(search, limit, page)

	totalPage := math.Ceil(float64(count) / float64(limit))
	next := 0
	prev := 0

	if int(totalPage) > 1 {
		next = int(totalPage) - page
	}
	if int(totalPage) > 1 {
		prev = int(totalPage) - 1
	}
	totalInfo := lib.PageInfo{
		TotalData: count,
		TotalPage: int(totalPage),
		Page:      page,
		Limit:     limit,
		Next:      next,
		Prev:      prev,
	}

	ctx.JSON(http.StatusOK, lib.Server{
		Success:     true,
		Message:     "List All Users",
		ResultsInfo: totalInfo,
		Results:     result,
	})
}

func SeeOneUserById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	dataUser, err := models.FindOneUser(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "User Not Found",
		})
		return
	}

	if dataUser.Id != 0 {
		ctx.JSON(http.StatusOK, lib.Server{
			Success: true,
			Message: "User Found",
			Results: dataUser,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Server{
			Success: false,
			Message: "User Not Found",
		})
	}
}

func CreateUser(ctx *gin.Context) {
	// search := ctx.Query("search")
	newUser := models.User{}
	// limit, _ := strconv.Atoi(ctx.Query("limit"))
	// offset, _ := strconv.Atoi(ctx.Query("offset"))

	// result, _ := models.FindAllUsers(search, limit, offset)

	if err := ctx.ShouldBind(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	// ids := 0
	// for _, v := range result {
	// 	ids = v.Id
	// }
	// newUser.Id = ids + 1

	user, err := models.CreateUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Server{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "User created successfully",
		Results: user,
	})
}

// func EditUser(c *gin.Context) {
// 	param := c.Param("id")
// 	id, _ := strconv.Atoi(param)
// 	search := c.Query("search")
// 	offset, _ := strconv.Atoi(c.Query("offset"))
// 	limit, _ := strconv.Atoi(c.Query("limit"))
// 	data := models.FindAllUsers(search, limit, offset)

// 	user := models.User{}
// 	err := c.Bind(&user)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	result := models.User{}
// 	for _, v := range data {
// 		if v.Id == id {
// 			result = v
// 		}
// 	}

// 	if result.Id == 0 {
// 		c.JSON(http.StatusNotFound, lib.Server{
// 			Success: false,
// 			Message: "user with id " + param + " not found",
// 		})
// 		return
// 	}

// 	ids := 0
// 	for _, v := range data {
// 		ids = v.Id
// 	}
// 	user.Id = ids

// 	models.EditUser(user.Email, *user.Username, user.Password, param)

// 	c.JSON(http.StatusOK, lib.Server{
// 		Success: true,
// 		Message: "user with id " + param + " Edit Success",
// 		Results: user,
// 	})
// }

func DeleteUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Invalid user ID format",
		})
		return
	}

	dataUser, err := models.FindOneUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, lib.Server{
			Success: false,
			Message: "User Not Found",
		})
		return
	}

	_, err = models.DeleteUserById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Server{
			Success: false,
			Message: "Failed to delete user",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "User deleted successfully",
		Results: dataUser,
	})
}

func UpdatePassword(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	pass := models.StructChangePassword{}
	found, err := models.FindPasswordById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	fmt.Println(found)
	if err := ctx.ShouldBind(&pass); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}
	fmt.Println(pass.OldPassword)
	isVerified := lib.Verify(pass.OldPassword, found.Password)
	fmt.Println(isVerified)
	if isVerified {
		err := models.ChangePassword(pass.Password, id)
		if err != nil {
			ctx.JSON(http.StatusOK, lib.Server{
				Success: true,
				Message: "Change password is Successfully",
			})
			return
		}
		ctx.JSON(http.StatusOK, lib.Server{
			Success: true,
			Message: " Update Successfully",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Wrong Password",
		})
	}
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	users := models.User{}
	err := ctx.Bind(&users)
	if err := ctx.ShouldBind(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Invalid input data",
		})

		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Failed to update profile",
		})
		return
	}
	models.UpdateUser(users, id)
	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Profile successfully updated",
	})
}

func ListAllFilterUsersWithPagination(c *gin.Context) {
	search := c.Query("search")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 5
	}
	products, err := models.FindAllUsersWithPagination(search, page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Failed to find profile",
		})
		return
	}
	c.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "List Users",
		Results: products,
	})
}

// func UpdateRoleUser(ctx *gin.Context) {
// 	id := ctx.GetInt("userId")

// 	var user models.User
// 	if err := ctx.ShouldBindJSON(&user); err != nil {
// 		ctx.JSON(http.StatusBadRequest, lib.Server{
// 			Success: false,
// 			Message: "Invalid input data",
// 		})
// 		return
// 	}

// 	adminId := ctx.GetInt("adminId")

// 	var adminUser models.User
// 	if err := models.FindUserById(adminId, &adminUser); err != nil || adminUser.UserRole != 1 {
// 		ctx.JSON(http.StatusForbidden, lib.Server{
// 			Success: false,
// 			Message: "Only admin can update the user role",
// 		})
// 		return
// 	}

// 	if err := models.UpdateUser(user, id); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, lib.Server{
// 			Success: false,
// 			Message: "Failed to update user profile",
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, lib.Server{
// 		Success: true,
// 		Message: "Profile successfully updated",
// 	})
// }
