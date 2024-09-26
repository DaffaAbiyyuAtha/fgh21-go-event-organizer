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

func SeeAllEvent(ctx *gin.Context) {
	search := ctx.Query("search")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	if limit < 1 {
		limit = 20
	}

	if page < 1 {
		page = 1
	}
	// if page > 1 {
	// 	page = (page - 1) * limit
	// }

	result, count := models.FindAllEvents(search, limit, page)

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
		Message:     "List All Event",
		ResultsInfo: totalInfo,
		Results:     result,
	})
}

func SeeOneEventById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	dataEvent := models.FindOneEvent(id)
	fmt.Println(dataEvent)

	if dataEvent.Id != 0 {

		ctx.JSON(http.StatusOK, lib.Server{
			Success: true,
			Message: "Event Found",
			Results: dataEvent,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Server{
			Success: false,
			Message: "Event Not Found",
		})
	}
}

func CreateEvent(ctx *gin.Context) {
	idaaa := ctx.GetInt("userId")
	newEvent := models.Events{}
	err := ctx.Bind(&newEvent)

	// result := models.FindAllEvents()
	fmt.Println(err)
	if err != nil {
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
	// newEvent.Id = ids + 1

	errr := models.CreateEvents(newEvent, idaaa)
	if errr != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Failed to Create Event",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Event created successfully",
		Results: newEvent,
	})
}

func SeeOneEventByUserId(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	fmt.Println(id)
	dataEvent := models.FindEventByUserId(id)

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Event Found",
		Results: dataEvent,
	})
}

func EditEvent(c *gin.Context) {
	// param := c.Param("id")
	// id, _ := strconv.Atoi(param)
	// data := models.FindAllEvents()

	// event := models.Events{}
	// err := c.Bind(&event)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// result := models.Events{}
	// for _, v := range data {
	// 	if v.Id == id {
	// 		result = v
	// 	}
	// }

	// if result.Id == 0 {
	// 	c.JSON(http.StatusNotFound, lib.Server{
	// 		Success: false,
	// 		Message: "event with id " + param + " not found",
	// 	})
	// 	return
	// }

	// ids := 0
	// for _, v := range data {
	// 	ids = v.Id
	// }
	// event.Id = ids

	// models.EditEvent(*event.Image, *event.Title, *event.Date, *event.Description, *event.Location_id, event.Created_by, param)

	// c.JSON(http.StatusOK, lib.Server{
	// 	Success: true,
	// 	Message: "event with id " + param + " Edit Success",
	// 	Results: event,
	// })
}

func DeleteEventById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	dataEvent := models.FindOneEvent(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	err = models.DeleteEvent(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Id Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "User deleted successfully",
		Results: dataEvent,
	})
}

func ListAllFilterEvents(c *gin.Context) {
	event := c.Query("event")

	events, err := models.GetAllEventWithFilters(event)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Id Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Events Has Been Filtered",
		Results: events,
	})
}
