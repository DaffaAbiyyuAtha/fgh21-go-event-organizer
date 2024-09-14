package models

import (
	"context"
	"fmt"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int     `json:"id" db:"id"`
	Email    string  `json:"email" form:"email" db:"email"`
	Password string  `json:"-" form:"password" db:"password"`
	Username *string `json:"username" form:"username" db:"username"`
}

type Passwords struct {
	Password string `json:"-"`
}

type StructChangePassword struct {
	OldPassword string `json:"-" form:"oldPassword"`
	Password    string `json:"-" form:"password" db:"password"`
}

func Total(search string) int {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT count(id) as "total" FROM "users" where "email" ilike '%' || $1 || '%'`
	rows := db.QueryRow(context.Background(), sql, search)
	var results int
	rows.Scan(
		&results,
	)
	return results
}

func FindAllUsers(search string, limit int, page int) ([]User, int) {
	db := lib.DB()

	defer db.Close(context.Background())

	offset := (page - 1) * limit

	rows, _ := db.Query(
		context.Background(),
		`select "id", "email", "password", "username" from "users" where "email" ilike '%' || $1 || '%' limit $2 offset $3`, search, limit, offset,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}
	result := Total(search)

	return users, result
}

func FindOneUser(id int) User {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		`select "id", "email", "password", "username" from "users"`,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}

	user := User{}
	for _, v := range users {
		if v.Id == id {
			user = v
		}
	}

	return user
}

func DeleteUser(id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	commandTag, err := db.Exec(
		context.Background(),
		`DELETE FROM "users" WHERE id = $1`,
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

func CreateUser(user User) (*User, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	user.Password = lib.Encrypt(user.Password)

	var users User

	err := db.QueryRow(
		context.Background(),
		`INSERT INTO "users" (email, "password", username) VALUES ($1, $2, $3) returning id, email`,
		user.Email, user.Password, user.Username,
	).Scan(&users.Id, &users.Email)

	if err != nil {
		return nil, fmt.Errorf("failed to execute insert")
	}

	return &users, nil
}

func EditUser(email string, username string, password string, id string) {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `update "users" set (email , username, password) = ($1, $2, $3) where id=$4`

	db.Exec(context.Background(), dataSql, email, username, password, id)

}

func FindOneUserByEmail(email string) User {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(
		context.Background(),
		`select "id", "email", "password","username" from "users"`,
	)

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}

	user := User{}
	for _, v := range users {
		if v.Email == email {
			user = v
		}
	}

	return user
}
func ChangePassword(password string, id int) error {
	db := lib.DB()
	defer db.Close(context.Background())
	editPassword := lib.Encrypt(password)

	dataSql := `UPDATE "users" SET password = $1 WHERE id = $2`
	_, err := db.Exec(context.Background(), dataSql, editPassword, id)
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	return nil
}

// func UpdateProfile(data Profile, id int) error {
// 	db := lib.DB()
// 	defer db.Close(context.Background())

// 	dataSql := `UPDATE "profile" SET "picture" = $1, "full_name" = $2, "birth_date" = $3, "gender" = $4, "phone_number" = $5, "profession" = $6, "nationality_id" = $7 WHERE "user_id" = $8`
// 	_, err := db.Exec(context.Background(), dataSql, data.Picture, data.Full_name, data.Birth_date, data.Gender, data.Phone_number, data.Profession, data.Nationality_id, id)
// 	if err != nil {
// 		return fmt.Errorf("failed to update profile: %v", err)
// 	}

// 	return nil
// }

func UpdateUser(data User, id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `UPDATE "users" SET "username" = $1, "email" = $2 WHERE "id" = $3`
	_, err := db.Exec(context.Background(), dataSql, data.Username, data.Email, id)
	if err != nil {
		return fmt.Errorf("failed to update profile: %v", err)
	}

	return nil
}

func FindPasswordById(id int) (Passwords, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT password FROM users WHERE id=$1`

	row, err := db.Query(context.Background(), sql, id)

	fmt.Println("1", err)
	if err != nil {
		return Passwords{}, err
	}
	user, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[Passwords])

	fmt.Println("2", err)
	if err != nil {
		return Passwords{}, err
	}
	return user, nil
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
