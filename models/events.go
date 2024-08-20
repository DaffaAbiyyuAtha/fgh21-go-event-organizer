package models

import (
	"context"
	"fmt"
	"time"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Events struct {
	Id          int       `json:"id" db:"id"`
	Image       *string   `json:"image" form:"image" db:"image"`
	Title       *string   `json:"title" form:"title" db:"title"`
	Date        time.Time `json:"date" form:"date" db:"date"`
	Description *string   `json:"description" form:"description" db:"description"`
	Location_id *int      `json:"location_id" form:"location_id" db:"location_id"`
	Created_by  *int      `json:"created_by" form:"created_by" db:"created_by"`
}

func FindAllEvents(search string, limit int, page int) ([]Events, int) {
	db := lib.DB()

	defer db.Close(context.Background())

	offset := (page - 1) * limit

	rows, _ := db.Query(
		context.Background(),
		`select * from "events" where "title" ilike '%' || $1 || '%' limit $2 offset $3`, search, limit, offset,
	)

	event, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Events])

	if err != nil {
		fmt.Println(err)
	}
	result := Total(search)

	return event, result
}

func FindOneEvent(id int) Events {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		`select "id", "image", "title", "date", "description", "location_id", "created_by" from "events"`,
	)

	events, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Events])

	if err != nil {
		fmt.Println(err)
	}

	event := Events{}
	for _, v := range events {
		if v.Id == id {
			event = v
		}
	}
	fmt.Println(event)

	return event
}

// func FindOneUserByEmail(email string) User {
// 	db := lib.DB()
// 	defer db.Close(context.Background())

// 	rows, _ := db.Query(
// 		context.Background(),
// 		`select "id", "email", "password","username" from "users"`,
// 	)

// 	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	user := User{}
// 	for _, v := range users {
// 		if v.Email == email {
// 			user = v
// 		}
// 	}

// 	return user
// }

func DeleteEvent(id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	commandTag, err := db.Exec(
		context.Background(),
		`DELETE FROM "events" WHERE id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("failed to execute delete")
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no user found")
	}

	return nil
}

func CreateEvents(event Events, id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	_, err := db.Exec(
		context.Background(),
		`INSERT INTO "events" ("image", "title", "date", "description", "location_id", "created_by") VALUES ($1, $2, $3, $4, $5, $6)`,
		event.Image, event.Title, event.Date, event.Description, event.Location_id, id,
	)
	fmt.Println(err)

	if err != nil {
		return fmt.Errorf("failed to execute insert")
	}

	return nil
}

func EditEvent(image string, title string, date string, description string, location_id int, created_by int, id string) {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `update "events" set ("image", "title", "date", "description", "location_id", "created_by") = ($1, $2, $3, $4, $5, $6) where id=$7`

	db.Exec(context.Background(), dataSql, image, title, date, description, location_id, created_by, id)

}

// package models

// import (
// 	"context"
// 	"fmt"

// 	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
// 	"github.com/jackc/pgx/v5"
// )

// type Users struct {
// 	Id       int    `json:"id"`
// 	Username string `json:"username" form:"username" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required"`
// 	Password string `json:"-" form:"password" binding:"required,min=8"`
// }

// var dataUser = []Users{
// 	{Id: 1, Username: "Admin", Email: "admin@mail.com", Password: "1234"},
// }

// func GetAllUsers() []Users {
// 	// data := dataUser

// 	// return data
// 	db := lib.DB()
// 	defer db.Close(context.Background())

// 	rows, _ := db.Query(
// 		context.Background(),
// 		`select * from "users"`,
// 	)

// 	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Users])

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return users
// }

// func GetOneUserById(id int) Users {
// 	data := dataUser

// 	user := Users{}
// 	for _, item := range data {
// 		if id == item.Id {
// 			user = item
// 		}
// 	}
// 	return user
// }

// func CreateUser(data Users) Users {
// 	id := 0
// 	for _, ids := range dataUser {
// 		id = ids.Id
// 	}

// 	data.Id = id + 1
// 	dataUser = append(dataUser, data)

// 	return data
// }

// func DeleteDataById(id int) Users {
// 	index := -1
// 	userDelete := Users{}
// 	for ids, item := range dataUser {
// 		if item.Id == id {
// 			index = ids
// 			userDelete = item
// 		}
// 	}
// 	if userDelete.Id != 0 {
// 		dataUser = append(dataUser[:index], dataUser[index+1:]...)
// 	}

// 	return userDelete
// }

// func UpdateDataById(data Users, id int) Users {

// 	ids := -1

// 	for index, item := range dataUser {
// 		if id == item.Id {
// 			ids = index
// 		}
// 	}

// 	if ids == 0 {
// 		dataUser[ids].Username = data.Username
// 		dataUser[ids].Email = data.Email
// 		dataUser[ids].Password = data.Password
// 		data.Id = dataUser[ids].Id
// 	}

// 	return data
// }
