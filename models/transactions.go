package models

import (
	"context"
	"fmt"
	"time"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Transactions struct {
	Id             int   `json:"id"`
	EventId        int   `json:"eventId" form:"eventId" db:"event_id"`
	PaymentId      int   `json:"paymentId" form:"paymentId" db:"payment_method_id"`
	UserId         int   `json:"userId"  db:"user_id"`
	SectionId      []int `json:"sectionId,omitempty" form:"sectionId" db:"section_id"`
	TicketQuantity []int `json:"ticketQuantity,omitempty" form:"ticketQuantity" db:"ticket_qty"`
}

type DetailTransaction struct {
	Id                int       `json:"id"`
	Full_name         string    `json:"full_name" form:"full_name" db:"full_name"`
	Event_title       string    `json:"event_title" form:"event_title" db:"title"`
	Location_id       *int      `json:"location_id" form:"location_id" db:"location"`
	Date              time.Time `json:"date" form:"date" db:"date"`
	Payment_method_id string    `json:"payment_method_id" form:"payment_method_id" db:"payment_method_id"`
	Section_name      []string  `json:"section_name" form:"section_name" db:"section_id"`
	Ticket_qty        []int     `json:"ticket_qty" form:"ticket_qty" db:"ticket_qty"`
}

type FindOneTransaction struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Location string `json:"location"`
	Payment  string `json:"payment"`
}

func CreateNewTransactions(data Transactions) (Transactions, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO "transactions" ("event_id", "payment_method_id", "user_id") 
	        VALUES ($1, $2, $3) RETURNING "id", "event_id", "payment_method_id", "user_id"`

	var results Transactions
	err := db.QueryRow(context.Background(), sql, data.EventId, data.PaymentId, data.UserId).Scan(
		&results.Id,
		&results.EventId,
		&results.PaymentId,
		&results.UserId,
	)
	fmt.Println(err)
	if err != nil {
		return Transactions{}, fmt.Errorf("failed to create transaction: %v", err)
	}

	return results, nil
}

func AddDetailsTransaction() ([]DetailTransaction, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		SELECT t.id, p.full_name, e.title AS "event_title", e.location_id, e.date, pm.name AS "payment_method",
		ARRAY_AGG(es.name) AS "section_name", ARRAY_AGG(td.ticket_qty) AS "ticket_qty"
		FROM "transactions" t
		JOIN "users" u ON u.id = t.user_id
		JOIN "profile" p ON p.user_id = u.id
		JOIN "events" e ON e.id = t.event_id
		JOIN "payment_methods" pm ON pm.id = t.payment_method_id
		JOIN "transaction_details" td ON td.transaction_id = t.id
		JOIN "event_sections" es ON es.id = td.section_id
		GROUP BY t.id, p.full_name, e.title, e.location_id, e.date, pm.name;
	`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("failed to query details: %v", err)
	}
	defer rows.Close()

	details, err := pgx.CollectRows(rows, pgx.RowToStructByPos[DetailTransaction])
	if err != nil {
		return nil, fmt.Errorf("failed to collect rows: %v", err)
	}

	return details, nil
}

func FindAllTransactionByUserId(userId int) ([]Transactions, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM "transactions" WHERE "user_id" = $1`
	rows, err := db.Query(context.Background(), sql, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to find transactions: %v", err)
	}
	defer rows.Close()

	transactions, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Transactions])
	if err != nil {
		return nil, fmt.Errorf("failed to collect transactions: %v", err)
	}

	return transactions, nil
}

func FindOneTransactionById(userId int) ([]FindOneTransaction, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		SELECT "t"."id", "e"."title", "e"."date", "l"."name" AS "location", "pm"."name" AS "payment"
		FROM "transactions" "t"
		JOIN "users" "u"
		ON "u"."id" = "t"."user_id"
		JOIN "events" "e"
		ON "e"."id" = "t"."event_id"
		JOIN "payment_methods" "pm"
		ON "pm"."id" = "t"."payment_method_id"
		JOIN "locations" "l"
		ON "l"."id" = "e"."location_id"
		WHERE "u"."id" = $1
		ORDER BY "t"."id" DESC
	`

	rows, err := db.Query(context.Background(), sql, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to find transactions: %v", err)
	}
	defer rows.Close()

	transactions, err := pgx.CollectRows(rows, pgx.RowToStructByPos[FindOneTransaction])
	if err != nil {
		return nil, fmt.Errorf("failed to collect transactions: %v", err)
	}

	return transactions, nil
}
