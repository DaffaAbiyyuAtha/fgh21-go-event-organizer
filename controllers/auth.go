package controllers

import (
	"fmt"
	"net/http"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

type Token struct {
	JWToken string `json:"token"`
}

func AuthLogin(ctx *gin.Context) {
	var user models.User
	ctx.Bind(&user)

	found := models.FindOneUserByEmail(user.Email)
	fmt.Println(user)

	if found == (models.User{}) {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Wrong Email or Password",
		})
		return
	}

	isVerified := lib.Verify(user.Password, found.Password)

	if isVerified {
		JWToken := lib.GenerateUserIdToken(found.Id)
		ctx.JSON(http.StatusOK, lib.Server{
			Success: true,
			Message: "Login Success",
			Results: Token{
				JWToken,
			},
		})
	} else {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Wrong Email or Password",
		})
	}
}

func AuthProfile(ctx *gin.Context) {
	account := models.Regist{}
	if err := ctx.ShouldBind(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	profile, err := models.CreateProfile(account)
	fmt.Println(profile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,
		lib.Server{
			Success: true,
			Message: "Create User success",
			Results: profile,
		})
}
