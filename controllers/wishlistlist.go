package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListAllWishlist(r *gin.Context) {
	results := models.FindAllwishlist()
	r.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "List All Wishlist",
		Results: results,
	})
}
func DetailWishlist(ctx *gin.Context) {
	id := ctx.GetInt("userId")

	wishlistItems, err := models.FindOnewishlist(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, lib.Server{
			Success: false,
			Message: "Wishlist Not Found",
		})
		return
	}

	if len(wishlistItems) == 0 {
		ctx.JSON(http.StatusNotFound, lib.Server{
			Success: false,
			Message: "No wishlist items found for this user",
		})
		return
	}

	var results []gin.H

	for _, wishlist := range wishlistItems {
		event, err := models.FindOneeventsbyid(wishlist.Event_id)
		if err != nil {
			log.Printf("Failed to fetch event with id %d: %v", wishlist.Event_id, err)
			continue
		}

		results = append(results, gin.H{
			"wishlist": wishlist,
			"event":    event,
		})
	}

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Wishlist and events found",
		Results: results,
	})
}

func Createwishlist(ctx *gin.Context) {
	var newWish models.Wishlist

	id, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, lib.Server{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	userId, ok := id.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, lib.Server{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	eventid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Invalid event ID",
		})
		return
	}

	log.Printf("userId: %d, eventid: %d", userId, eventid)

	err = models.Createwishlist(eventid, userId)
	if err != nil {

		log.Printf("Createwishlist error: %v", err)

		if err.Error() == "wishlist entry already exists" {
			ctx.JSON(http.StatusConflict, lib.Server{
				Success: false,
				Message: "Event is already in your wishlist",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, lib.Server{
			Success: false,
			Message: "Failed to create Wishlist",
		})
		return
	}

	newWish.User_id = userId
	newWish.Event_id = eventid

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Wishlist created successfully",
		Results: newWish,
	})
}
func DeleteWishlistById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	selectWishlist, err := models.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Invalid event ID",
		})
		return
	}

	err = models.DeleteWishlist(models.Wishlist{}, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "Failed to Delete Wishlist",
		})
		return
	}

	c.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "Wishlist Delete Successfully",
		Results: selectWishlist,
	})
}
