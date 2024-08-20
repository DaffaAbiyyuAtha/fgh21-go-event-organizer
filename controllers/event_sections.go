package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func SeeOneEventByEventId(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	dataEvent, err := models.FindSectionsByEventId(id)
	fmt.Println(dataEvent)

	if err != nil {
		ctx.JSON(http.StatusNotFound, lib.Server{
			Success: false,
			Message: "Event Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Event Found",
		Results: dataEvent,
	})
}
