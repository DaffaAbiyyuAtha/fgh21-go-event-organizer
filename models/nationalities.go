package models

import (
	"context"
	"fmt"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Nationality struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FindAllNationalities() []Nationality {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "nationalities" order by "id" asc`,
	)
	nationalities, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Nationality])
	if err != nil {
		fmt.Println(err)
	}
	return nationalities
}
