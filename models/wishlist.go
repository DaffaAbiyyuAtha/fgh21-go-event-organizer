package models

import (
	"context"
	"fmt"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Wishlist struct {
	Id       int `json:"id"`
	User_id  int `json:"user_id" form:"user_id"`
	Event_id int `json:"event_id" form:"event_id"`
}
type JoinEvents struct {
	Id          int     `json:"id"`
	Image       *string `json:"image" form:"image" db:"image"`
	Title       *string `json:"title" form:"title" db:"title"`
	Date        *string `json:"date" form:"date" db:"date"`
	Description *string `json:"description" form:"description" db:"description"`
	Location    *int    `json:"location_id" form:"location_id" db:"location_id"`
	Created_by  *int    `json:"created_by" form:"created_by" db:"created_by"`
}

func FindAllwishlist() []Wishlist {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "wishlist" order by "id" asc`,
	)

	wishlists, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Wishlist])
	if err != nil {
		fmt.Println(err)
	}
	return wishlists
}
func FindOnewishlist(id int) ([]Wishlist, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	// Query the wishlist items for the given user ID
	rows, err := db.Query(context.Background(),
		`SELECT * FROM "wishlist" WHERE "user_id" = $1 ORDER BY "id" ASC`, id,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query wishlist: %w", err)
	}
	defer rows.Close() // Ensure rows are closed after the operation

	// Collect the results into a slice of Wishlist structs
	wishlists, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Wishlist])
	if err != nil {
		return nil, fmt.Errorf("failed to collect wishlist rows: %w", err)
	}

	// Return the collected wishlists and nil error if successful
	return wishlists, nil
}
func Createwishlist(event_id int, id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	var exists bool
	err := db.QueryRow(
		context.Background(),
		`SELECT EXISTS (SELECT 1 FROM "wishlist" WHERE user_id = $1 AND event_id = $2)`,
		id, event_id,
	).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check existing wishlist entry: %w", err)
	}

	if exists {
		return fmt.Errorf("wishlist entry already exists")
	}

	_, err = db.Exec(
		context.Background(),
		`INSERT INTO "wishlist" (user_id, event_id) VALUES ($1, $2)`,
		id, event_id,
	)

	if err != nil {
		return fmt.Errorf("failed to insert wishlist entry: %w", err)
	}

	return nil
}
func FindOneeventsbyid(event_id int) (JoinEvents, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	var event JoinEvents
	err := db.QueryRow(context.Background(),
		`SELECT id, image, title, date, description, location_id, created_by 
         FROM "events" WHERE id = $1`, event_id,
	).Scan(&event.Id, &event.Image, &event.Title, &event.Date, &event.Description, &event.Location, &event.Created_by)

	if err != nil {
		return JoinEvents{}, fmt.Errorf("failed to find event: %w", err)
	}

	return event, nil
}
func DeleteWishlist(data Wishlist, id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `DELETE FROM wishlist WHERE id=$1;`

	query, _ := db.Exec(context.Background(), sql, id)

	if query.RowsAffected() == 0 {
		return fmt.Errorf("data not found")
	}

	return nil
}
func GetProductById(id int) (Wishlist, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * from wishlist WHERE id=$1;`

	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return Wishlist{}, err
	}

	selectedRow, err := pgx.CollectOneRow(query, pgx.RowToStructByName[Wishlist])

	if err != nil {
		return Wishlist{}, err
	}

	return selectedRow, err
}
