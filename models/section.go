package models

import (
	"context"
	"fmt"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Sactions struct {
	Name     []string `json:"name" form:"name" db:"name"`
	Price    []int    `json:"price" form:"price" db:"price"`
	Quantity []int    `json:"quantity" form:"quantity" db:"quantity"`
	EventId  int      `json:"eventId" form:"eventId" db:"event_id"`
}

// func CreateSaction(saction Sactions) ([]Sactions, error) {
// 	db := lib.DB()
// 	defer db.Close(context.Background())

// 	sql := `INSERT INTO "event_sections"
// 		("name", "price", "quantity", "event_id")
// 		VALUES
// 		($1, $2, $3, $4) RETURNING id`

// 	rows, err := db.Query(context.Background(), sql, saction.Name, saction.Price, saction.Quantity, saction.EventId)

// 	if err != nil {
// 		return []Sactions{}, err
// 	}

// 	// sactions, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Sactions])
// 	sactions, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Sactions])

// 	if err != nil {
// 		return []Sactions{}, err
// 	}

// 	return sactions, err
// }

// func CreateSaction(saction Sactions) (Sactions, error) {
// 	db, err := pgx.Connect(context.Background(), "your_database_connection_string")
// 	if err != nil {
// 		return Sactions{}, err
// 	}
// 	defer db.Close(context.Background())

// 	query := `INSERT INTO sactions (name, price, quantity, event_id) VALUES ($1, $2, $3, $4) RETURNING id`
// 	_, err = db.Exec(context.Background(), query, saction.Name, saction.Price, saction.Quantity, saction.EventId)
// 	if err != nil {
// 		return Sactions{}, err
// 	}

// 	return saction, nil
// }

// func CreateSaction(saction Sactions) ([]int, error) {
// 	db := lib.DB()
// 	defer db.Close(context.Background())

// 	var sactionIDs []int // Menyimpan id yang berhasil diinsert

// 	for i := 0; i < len(saction.Name); i++ {
// 		sql := `INSERT INTO "event_sections" ("name", "price", "quantity", "event_id")
//                 VALUES ($1, $2, $3, $4) RETURNING id`

// 		fmt.Printf("Inserting: name=%s, price=%d, quantity=%d, event_id=%d\n",
// 			saction.Name[i], saction.Price[i], saction.Quantity[i], saction.EventId)

// 		var id int
// 		err := db.QueryRow(context.Background(), sql, saction.Name[i], saction.Price[i], saction.Quantity[i], saction.EventId).Scan(&id)
// 		if err != nil {
// 			fmt.Println("QueryRow error:", err)
// 			return nil, err
// 		}

// 		sactionIDs = append(sactionIDs, id)
// 	}

// 	return sactionIDs, nil
// }

// func CreateSaction(saction Sactions) ([]Sactions, error) {
// 	db := lib.DB()
// 	defer db.Close(context.Background())

// 	sql := `INSERT INTO "event_sections"
// 		("name", "price", "quantity", "event_id")
// 		VALUES
// 		($1, $2, $3, $4) RETURNING "id", "name", "price", "quantity", "event_id"`

// 	row := db.QueryRow(context.Background(), sql, saction.Name, saction.Price, saction.Quantity, saction.EventId)

// 	var createdSaction Sactions
// 	err := row.Scan(&createdSaction.Name, &createdSaction.Price, &createdSaction.Quantity, &createdSaction.EventId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return []Sactions{createdSaction}, nil
// }

func CreateSaction(saction Sactions) ([]int, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	var sactionIDs []int // Menyimpan id yang berhasil diinsert

	for i := 0; i < len(saction.Name); i++ {
		sql := `INSERT INTO "event_sections" ("name", "price", "quantity", "event_id")
                VALUES ($1, $2, $3, $4) RETURNING id`

		fmt.Printf("Inserting: name=%s, price=%d, quantity=%d, event_id=%d\n",
			saction.Name[i], saction.Price[i], saction.Quantity[i], saction.EventId)

		// Menjalankan query dan mengumpulkan hasilnya
		row, err := db.Query(context.Background(), sql, saction.Name[i], saction.Price[i], saction.Quantity[i], saction.EventId)
		if err != nil {
			fmt.Println("Query error:", err)
			return nil, err
		}

		// Mengambil ID yang diinsert
		id, err := pgx.CollectOneRow(row, pgx.RowTo[int])
		if err != nil {
			fmt.Println("CollectOneRow error:", err)
			return nil, err
		}

		sactionIDs = append(sactionIDs, id)
	}

	return sactionIDs, nil
}
