package controllers

import (
	"net/http"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllNationalities(r *gin.Context) {
	results := models.FindAllNationalities()
	r.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "List All Wishlist",
		Results: results,
	})
}
