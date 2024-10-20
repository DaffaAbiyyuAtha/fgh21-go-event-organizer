package controllers

import (
	"net/http"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

// func CreateSaction(ctx *gin.Context) {
// 	var form models.Sactions

// 	if err := ctx.ShouldBind(&form); err != nil {
// 		ctx.JSON(http.StatusBadRequest, lib.Server{
// 			Success: false,
// 			Message: "error to bind form data",
// 		})
// 		return
// 	}

// 	for i := 0; i < len(form.Name); i++ {
// 		saction := models.Sactions{
// 			Name:     []string{form.Name[i]},
// 			Price:    []int{form.Price[i]},
// 			Quantity: []int{form.Quantity[i]},
// 			EventId:  form.EventId,
// 		}
// 		fmt.Println("Saction to be created:", saction)

// 		created, err := models.CreateSaction(saction)
// 		if err != nil {
// 			fmt.Println("Database error:", err)
// 			ctx.JSON(http.StatusBadRequest, lib.Server{
// 				Success: false,
// 				Message: "Failed to create saction",
// 			})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, lib.Server{
// 			Success: true,
// 			Message: "Saction created successfully",
// 			Results: created,
// 		})
// 	}
// }

func CreateSaction(ctx *gin.Context) {
	var form models.Sactions
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sactionIDs, err := models.CreateSaction(models.Sactions{
		Name:     form.Name,
		Price:    form.Price,
		Quantity: form.Quantity,
		EventId:  form.EventId,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create section",
			"error":   err.Error(),
		})
		return
	}

	if len(sactionIDs) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "No section created",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Section created successfully",
		"results": sactionIDs,
	})
}
