package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	dataUser, err := models.FindOneUser(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "User Not Found",
		})
		return
	}

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

// func UpdateProfile(ctx *gin.Context) {
// 	id := ctx.GetInt("userId")
// 	profiles := models.Profile{}
// 	if err := ctx.ShouldBind(&profiles); err != nil {
// 		ctx.JSON(http.StatusBadRequest, lib.Server{
// 			Success: false,
// 			Message: "Invalid input data",
// 		})
// 		return
// 	}
// 	err := models.UpdateProfile(profiles.Full_name, *profiles.Phone_number, profiles.Profession, id)

// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, lib.Server{
// 			Success: false,
// 			Message: "failed Update Password",
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, lib.Server{
// 		Success: true,
// 		Message: "Password Successfully updated",
// 	})
// }

func UpdateProfile(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	profiles := models.Profile{}
	users := models.User{}
	err := ctx.Bind(&profiles)
	errr := ctx.Bind(&users)
	dataProfile := models.FindProfileByUserId(id)
	dataUser, err := models.FindOneUser(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "User Not Found",
		})
		return
	}

	if err := ctx.ShouldBind(&profiles); err != nil {
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

	if errr != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Failed to update users",
		})
		return
	}

	models.UpdateProfile(profiles, id)
	models.UpdateUser(users, id)
	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Profile successfully updated",
		Results: gin.H{
			"profile": dataProfile,
			"user":    dataUser,
		},
	})
}

func UpdateProfilePicture(c *gin.Context) {
	id := c.GetInt("userId")

	file, err := c.FormFile("picture")
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "No file",
		})
		return
	}

	cek := map[string]bool{".jpg": true, ".png": true, ".jpeg": true}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !cek[ext] {
		c.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Failed to Upload File",
		})
		return
	}

	picture := uuid.New().String() + ext

	savePicture := "./picture/"
	if err := c.SaveUploadedFile(file, savePicture+picture); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Failed to Save File",
		})
		return
	}

	root := "http://localhost:8080/picture/" + picture

	profile, err := models.UpdateProfilePicture(models.Profile{Picture: &root}, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Faileded to Save File",
		})
		return
	}

	c.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Profile picture updated successfully",
		Results: profile,
	})
}
