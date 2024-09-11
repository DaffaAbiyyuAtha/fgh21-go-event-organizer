// package models

// import (
// 	"context"
// 	"fmt"

// 	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
// )

// type TransactionDetail struct {
// 	Id            int `json:"id"`
// 	TransactionId int `json:"transactionId" db:"transaction_id"`
// 	SectionId     int `json:"sectionId" db:"section_id"`
// 	TicketQty     int `json:"ticketQty" db:"ticket_qty"`
// }

// func CreateTransactionDetail(data TransactionDetail) TransactionDetail {
// 	db := lib.DB()
// 	fmt.Println(db)
// 	defer db.Close(context.Background())

// 	inputSQL := `insert into "transaction_details" (transaction_id, section_id, ticket_qty) values ($1, $2, $3) returning "id", "transaction_id", "section_id", "ticket_qty"`
// 	row := db.QueryRow(context.Background(), inputSQL, data.TransactionId, data.SectionId, data.TicketQty)

// 	var detail TransactionDetail

// 	row.Scan(
// 		&detail.Id,
// 		&detail.TransactionId,
// 		&detail.SectionId,
// 		&detail.TicketQty,
// 	)
// 	return detail
// }

package models

import (
	"context"
	"fmt"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
)

type TransactionDetail struct {
	Id            int `json:"id"`
	TransactionId int `json:"transactionId" db:"transaction_id"`
	SectionId     int `json:"sectionId" db:"section_id"`
	TicketQty     int `json:"ticketQty" db:"ticket_qty"`
}

func CreateTransactionDetail(data TransactionDetail) (TransactionDetail, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO "transaction_details" (transaction_id, section_id, ticket_qty) 
	        VALUES ($1, $2, $3) RETURNING "id", "transaction_id", "section_id", "ticket_qty"`

	var detail TransactionDetail
	err := db.QueryRow(context.Background(), sql, data.TransactionId, data.SectionId, data.TicketQty).Scan(
		&detail.Id,
		&detail.TransactionId,
		&detail.SectionId,
		&detail.TicketQty,
	)
	if err != nil {
		return TransactionDetail{}, fmt.Errorf("failed to create transaction detail: %v", err)
	}

	return detail, nil
}
