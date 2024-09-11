package models

import (
	"context"
	"fmt"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Partner struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" form:"name" db:"name"`
	Image string `json:"image" form:"image" db:"image"`
}

func FindAllPartners(search string, limit int, page int) []Partner {
	db := lib.DB()

	defer db.Close(context.Background())

	offset := (page - 1) * limit

	rows, _ := db.Query(
		context.Background(),
		`select * from "partners" where "name" ilike '%' || $1 || '%' limit $2 offset $3`, search, limit, offset,
	)

	partners, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Partner])

	if err != nil {
		fmt.Println(err)
	}
	return partners
}
