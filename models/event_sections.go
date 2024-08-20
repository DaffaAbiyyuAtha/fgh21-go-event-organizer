package models

import (
	"context"
	"fmt"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type EventSections struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" form:"name" db:"name"`
	Price    int    `json:"price" form:"price" db:"price"`
	Quantity int    `json:"quantity" form:"quantity" db:"quantity"`
	EventId  int    `json:"eventId"`
}

func FindSectionsByEventId(eventId int) ([]EventSections, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, err := db.Query(
		context.Background(),
		`select * from "event_sections" where "event_id" = $1`, eventId,
	)
	if err != nil {
		fmt.Println(err)
	}
	es, err := pgx.CollectRows(rows, pgx.RowToStructByPos[EventSections])

	if err != nil {
		return nil, fmt.Errorf("Error")
	}

	return es, nil
}
