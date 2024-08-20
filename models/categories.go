package models

import (
	"context"
	"fmt"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Category struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" form:"name" db:"name"`
}

func FindAllCategories(search string, limit int, page int) []Category {
	db := lib.DB()

	defer db.Close(context.Background())

	offset := (page - 1) * limit

	rows, _ := db.Query(
		context.Background(),
		`select "id", "name" from "categories" where "name" ilike '%' || $1 || '%' limit $2 offset $3`, search, limit, offset,
	)

	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Category])

	if err != nil {
		fmt.Println(err)
	}
	return categories
}

func FindOneCategory(id int) Category {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		`select "id", "name" from "categories"`,
	)

	categories, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Category])

	if err != nil {
		fmt.Println(err)
	}

	category := Category{}
	for _, v := range categories {
		if v.Id == id {
			category = v
		}
	}

	return category
}

func DeleteCategory(id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	commandTag, err := db.Exec(
		context.Background(),
		`DELETE FROM "categories" WHERE id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to execute delete")
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no category found")
	}

	return nil
}

func CreateCategory(category Category) error {
	db := lib.DB()
	defer db.Close(context.Background())

	_, err := db.Exec(
		context.Background(),
		`INSERT INTO "categories" (name) VALUES ($1)`,
		category.Name,
	)

	if err != nil {
		return fmt.Errorf("failed to execute insert")
	}

	return nil
}

func EditCategory(name string, id string) {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `update "categories" set (name) = ($1) where id=$2`

	db.Exec(context.Background(), dataSql, name, id)

}
