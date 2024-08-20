package controllers

import (
	"net/http"
	"strconv"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func SeeAllLocations(ctx *gin.Context) {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	if limit < 1 {
		limit = 7
	}

	if page < 1 {
		page = 1
	}

	result := models.FindAllLocations(search, limit, page)

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "List All Locations",
		Results: result,
	})
}
