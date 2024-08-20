package controllers

import (
	"fmt"
	"net/http"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

// func CreateProfiles(ctx *gin.Context) {
// 	newUser := models.Profile{}
// 	result := models.FindAllUsers()

// 	if err := ctx.ShouldBind(&newUser); err != nil {
// 		ctx.JSON(http.StatusBadRequest, lib.Server{
// 			Success: false,
// 			Message: "Invalid input data",
// 		})
// 		return
// 	}

// 	ids := 0
// 	for _, v := range result {
// 		ids = v.Id
// 	}
// 	newUser.Id = ids + 1

// 	err := models.CreateProfile(newUser)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, lib.Server{
// 			Success: false,
// 			Message: "Failed to create user",
// 		})
// 		return
// 	}

//		ctx.JSON(http.StatusOK, lib.Server{
//			Success: true,
//			Message: "User created successfully",
//			Results: newUser,
//		})
//	}
// func CreateProfiles(ctx *gin.Context) {
// 	var newUser models.Profile

// 	if err := ctx.ShouldBind(&newUser); err != nil {
// 		ctx.JSON(http.StatusBadRequest, lib.Server{
// 			Success: false,
// 			Message: "Invalid input data",
// 		})
// 		return
// 	}

// 	err := models.CreateProfile(newUser)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, lib.Server{
// 			Success: false,
// 			Message: "Failed to create profile: " + err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, lib.Server{
// 		Success: true,
// 		Message: "Profile created successfully",
// 		Results: newUser,
// 	})
// }

func CreateProfiles(ctx *gin.Context) {
	account := models.Regist{}

	if err := ctx.ShouldBind(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := models.CreateProfile(account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Create User success",
		Results: gin.H{
			"id":       profile.Id,
			"fullname": profile.Full_name,
			"email":    account.Email,
		},
	})
}

func SeeOneProfileByUserId(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	fmt.Println(id)
	dataProfile := models.FindProfileByUserId(id)
	dataUser := models.FindOneUser(id)

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Event Found",
		Results: gin.H{
			"profile": dataProfile,
			"user":    dataUser,
		},
	})
}

func ListAllProfile(r *gin.Context) {
	results := models.FindAllProfile()
	r.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "List All Profile",
		Results: results,
	})
}
