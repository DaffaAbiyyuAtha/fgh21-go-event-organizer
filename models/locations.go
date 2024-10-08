package models

import (
	"context"
	"fmt"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Locations struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" form:"name" db:"name"`
	Lat   string `json:"lat" form:"lat" db:"lat"`
	Long  string `json:"long" form:"long" db:"long"`
	Image string `json:"image" form:"image" db:"image"`
}

type JLocations struct {
	Id            int    `json:"id"`
	Image         string `json:"image" form:"image"`
	Title         string `json:"title" form:"title"`
	Name          string `json:"name" form:"name"`
	LocationImage string `json:"location_mage" form:"location_mage"`
}

func FindAllLocations(search string, limit int, page int) []Locations {
	db := lib.DB()

	defer db.Close(context.Background())

	offset := (page - 1) * limit

	rows, _ := db.Query(
		context.Background(),
		`select * from "locations" where "name" ilike '%' || $1 || '%' limit $2 offset $3`, search, limit, offset,
	)

	locations, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Locations])

	if err != nil {
		fmt.Println(err)
	}
	return locations
}

func GetAllEventWithFilter(location string) ([]JLocations, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
		SELECT "e"."id", "e"."image", "e"."title", "l"."name", "l"."image" AS "location_image"
		FROM "events" "e"
		JOIN "locations" "l"
		ON "l"."id" = "e"."location_id"
		WHERE "name" ILIKE $1
	`

	rows, err := db.Query(context.Background(), sql, "%"+location+"%")

	if err != nil {
		return []JLocations{}, err
	}

	locations, err := pgx.CollectRows(rows, pgx.RowToStructByPos[JLocations])

	if err != nil {
		return []JLocations{}, err
	}

	return locations, err
}
