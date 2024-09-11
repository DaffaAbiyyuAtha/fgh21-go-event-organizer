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
// type FormTransactions struct {
// 	Event_id          int   `json:"event_id" form:"event_id" db:"event_id"`
// 	Payment_method_id int   `json:"payment_method_id" form:"payment_method_id" db:"payment_method_id"`
// 	Section_id        []int `json:"section_id" form:"section_id" db:"section_id"`
// 	Ticket_qty        []int `json:"ticket_qty" form:"ticket_qty" db:"ticket_qty"`
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

func CreateTransaction(ctx *gin.Context) {
	form := models.Transactions{}
	// id := ctx.GetInt("userId")
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(&form)
	trx, err := models.CreateNewTransactions(models.Transactions{
		UserId:    ctx.GetInt("userId"),
		PaymentId: form.PaymentId,
		EventId:   form.EventId,
	})
	for i := range form.SectionId {
		models.CreateTransactionDetail(models.TransactionDetail{
			TransactionId: trx.Id,
			SectionId:     form.SectionId[i],
			TicketQty:     form.TicketQuantity[i],
		})
		if err != nil {
			fmt.Println("Error creating transaction detail: ", err)
		} else {
			fmt.Println("Transaction detail created for SectionId: ", form.SectionId[i])
		}
	}
	details, err := models.AddDetailsTransaction()
	fmt.Println("Details returned:", details)
	if err != nil || len(details) == 0 {
		ctx.JSON(http.StatusNotFound, lib.Server{
			Success: false,
			Message: "Transaction details not found or empty",
		})
		return
	}
	ctx.JSON(http.StatusOK,
		lib.Server{
			Success: true,
			Message: "Create Transaction success",
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

func ListProductById(c *gin.Context) {
	id := c.GetInt("userId")
	selected, err := models.FindOneTransactionById(id)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Server{
			Success: false,
			Message: "invalid user ID",
		})
		return
	}

	c.JSON(http.StatusOK, lib.Server{
		Success: true,
		Message: "transaction found",
		Results: selected,
	})
}
