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
	UserRole int     `json:"userRole" form:"userRole" db:"user_role"`
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

func FindOneUser(id int) (User, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `select "id", "email", "password", "username", "user_role" from "users" where id = $1`

	query, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return User{}, err
	}

	users, err := pgx.CollectOneRow(query, pgx.RowToStructByPos[User])

	if err != nil {
		return User{}, err
	}

	return users, err
}

func DeleteUserById(id int) (string, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `DELETE FROM users WHERE id=$1 RETURNING id`

	row, err := db.Query(context.Background(), sql, id)

	if err != nil {
		return "", err
	}

	result, err := pgx.CollectOneRow(row, pgx.RowTo[string])

	if err != nil {
		return "", err
	}

	return result, nil
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
		`select "id", "email", "password", "username", "user_role" from "users"`,
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

	if err != nil {
		return Passwords{}, err
	}
	user, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[Passwords])

	if err != nil {
		return Passwords{}, err
	}
	return user, nil
}

func FindAllUsersWithPagination(search string, page int, limit int) ([]User, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	var offset int = (page - 1) * limit

	sql := `select "id", "email", "password", "username", "user_role" 
		from "users" 
		where "email" ilike $1
		ORDER BY "id" ASC
		limit $2 offset $3`

	rows, err := db.Query(context.Background(), sql, "%"+search+"%", limit, offset)

	if err != nil {
		return []User{}, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		return []User{}, err
	}

	return users, err
}

// func FindUserById(id int) (User, error) {
// 	db := lib.DB()
// 	defer db.Close(context.Background())

// 	sql := `SELECT id, username, email, password, user_role
// 			FROM users
// 			WHERE id = $1`

// 	// Eksekusi query dan kumpulkan hasil menggunakan CollectOneRow
// 	query, err := db.Query(context.Background(), sql, id)
// 	if err != nil {
// 		return User{}, err
// 	}

// 	user, err := pgx.CollectOneRow(query, pgx.RowToStructByPos[User])
// 	if err != nil {
// 		if err == pgx.ErrNoRows {
// 			return User{}, errors.New("user not found")
// 		}
// 		return User{}, err
// 	}

// 	return user, nil
// }

// func EditRoleUser(adminId int, tagerUserId int, newRole int) (User, error) {
// 	db := lib.DB()
// 	defer db.Close(context.Background())

// 	sql := `UPDATE users
// 		SET user_role = $1
// 		WHERE id = $2
// 		AND EXISTS (
// 			SELECT 1
// 			FROM users
// 			WHERE id = $3
// 			AND user_role = 1
// 		);`

// 	query, err := db.Query(context.Background(), sql, newRole, tagerUserId, adminId)

// 	if err != nil {
// 		return User{}, err
// 	}

// 	users, err := pgx.CollectOneRow(query, pgx.RowToStructByPos[User])

// 	if err != nil {
// 		return User{}, err
// 	}

// 	return users, err
// }
