// package controllers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
// 	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/models"
// 	"github.com/gin-gonic/gin"
// )

// type FormTransactions struct {
// 	Event_id          int   `json:"event_id" form:"event_id" db:"event_id"`
// 	Payment_method_id int   `json:"payment_method_id" form:"payment_method_id" db:"payment_method_id"`
// 	Section_id        []int `json:"section_id" form:"section_id" db:"section_id"`
// 	Ticket_qty        []int `json:"ticket_qty" form:"ticket_qty" db:"ticket_qty"`
// }

// func CreateTransaction(ctx *gin.Context) {
// 	form := FormTransactions{}

// 	if err := ctx.ShouldBind(&form); err != nil {
// 		ctx.JSON(http.StatusBadRequest, lib.Server{
// 			Success: false,
// 			Message: "invalid input data",
// 		})
// 		return
// 	}

// 	trx := models.CreateNewTransactions(models.Transaction{
// 		User_id:           ctx.GetInt("userId"),
// 		Payment_method_id: form.Payment_method_id,
// 		Event_id:          form.Event_id,
// 	})

// 	for i := range form.Section_id {
// 		models.CreateTransactionDetail(models.TransactionDetail{
// 			SectionId:     form.Section_id[i],
// 			TicketQty:     form.Ticket_qty[i],
// 			TransactionId: trx.Id,
// 		})

// 	}
// 	details, err := models.AddDetailsTransaction()
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, lib.Server{
// 			Success: false,
// 			Message: "transaction Not Found",
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, lib.Server{
// 		Success: true,
// 		Message: "transaction Found",
// 		Results: details,
// 	})
// }

// func ListDetailsTransactions(ctx *gin.Context) {
// 	details, err := models.AddDetailsTransaction()
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, lib.Server{
// 			Success: false,
// 			Message: "Transaction Not Found",
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, lib.Server{
// 		Success: true,
// 		Message: "Transaction Found",
// 		Results: details,
// 	})
// }

// func FindTransactionByUserId(ctx *gin.Context) {
// 	id, _ := strconv.Atoi(ctx.Param("id"))
// 	data, err := models.FindAllTransactionByUserId(id)

// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, lib.Server{
// 			Success: false,
// 			Message: "Transaction Not Found",
// 		})
// 	}

// 	ctx.JSON(http.StatusOK, lib.Server{
// 		Success: true,
// 		Message: "Transaction Found",
// 		Results: data,
// 	})

// }

package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

type FormTransactions struct {
	Event_id          int   `json:"event_id" form:"event_id" db:"event_id"`
	Payment_method_id int   `json:"payment_method_id" form:"payment_method_id" db:"payment_method_id"`
	Section_id        []int `json:"section_id" form:"section_id" db:"section_id"`
	Ticket_qty        []int `json:"ticket_qty" form:"ticket_qty" db:"ticket_qty"`
}

func CreateTransaction(ctx *gin.Context) {
	form := FormTransactions{}
	id, _ := strconv.Atoi(ctx.Param("id"))
	userId := ctx.GetInt("userId")
	fmt.Println(userId)
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "invalid input data",
		})
		return
	}

	trx, err := models.CreateNewTransactions(models.Transaction{
		User_id:           userId,
		Payment_method_id: form.Payment_method_id,
		Event_id:          id,
	})
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Server{
			Success: false,
			Message: "failed to create transaction",
		})
		return
	}

	for i := range form.Section_id {
		_, err := models.CreateTransactionDetail(models.TransactionDetail{
			SectionId:     form.Section_id[i],
			TicketQty:     form.Ticket_qty[i],
			TransactionId: trx.Id,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, lib.Server{
				Success: false,
				Message: "failed to create transaction detail",
			})
			return
		}
	}

	details, err := models.AddDetailsTransaction()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, lib.Server{
			Success: false,
			Message: "transaction details not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "transaction created successfully",
		Results: details,
	})
}

func ListDetailsTransactions(ctx *gin.Context) {
	details, err := models.AddDetailsTransaction()
	if err != nil {
		ctx.JSON(http.StatusNotFound, lib.Server{
			Success: false,
			Message: "transaction details not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "transaction details found",
		Results: details,
	})
}

func FindTransactionByUserId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "invalid user ID",
		})
		return
	}

	data, err := models.FindAllTransactionByUserId(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, lib.Server{
			Success: false,
			Message: "transaction not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "transaction found",
		Results: data,
	})
}
