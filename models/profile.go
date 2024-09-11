package models

import (
	"context"
	"fmt"

	"github.com/daffaabiyyuatha/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Profile struct {
	Id             int     `json:"id"`
	Picture        *string `json:"picture" form:"picture" db:"picture"`
	Full_name      string  `json:"full_name" form:"full_name" db:"full_name"`
	Birth_date     *string `json:"birth_date," form:"birth_date" db:"birth_date"`
	Gender         *int    `json:"gender" form:"gender" db:"gender"`
	Phone_number   *string `json:"phone_number" form:"phone_number" db:"phone_number"`
	Profession     *string `json:"profession" form:"profession"`
	Nationality_id *int    `json:"nationality_id" form:"nationality_id" db:"nationality_id"`
	User_id        int     `json:"user_id" form:"user_id" db:"user_id"`
}

type Regist struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" db:"email"`
	Password string `json:"-" form:"password" db:"password"`
	Profile  Profile
}

func CreateProfile(regist Regist) (*Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	regist.Password = lib.Encrypt(regist.Password)

	var userId int
	err := db.QueryRow(
		context.Background(),
		`INSERT INTO "users" ("email", "password") VALUES ($1 ,$2) RETURNING "id"`,
		regist.Email, regist.Password,
	).Scan(&userId)

	if err != nil {
		return nil, fmt.Errorf("failed to insert into users table: %v", err)
	}

	profiles := Profile{
		Full_name: regist.Profile.Full_name,
		User_id:   userId,
	}

	var profile Profile
	err = db.QueryRow(
		context.Background(),
		`INSERT INTO "profile" ("picture", "full_name", "birth_date", "gender", "phone_number", "profession", "nationality_id", "user_id")
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, picture, full_name, birth_date, gender, phone_number, profession, nationality_id, user_id`,
		regist.Profile.Picture, regist.Profile.Full_name, regist.Profile.Birth_date, regist.Profile.Gender,
		regist.Profile.Phone_number, regist.Profile.Profession, regist.Profile.Nationality_id, userId,
	).Scan(
		&profile.Id, &profile.Picture, &profile.Full_name, &profile.Birth_date,
		&profile.Gender, &profile.Phone_number, &profile.Profession, &profile.Nationality_id, &profile.User_id,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert into profile table: %v", err)
	}

	return &profiles, nil
}

func ListAllProfile() []Profile {
	db := lib.DB()
	defer db.Close(context.Background())

	joinSql := `select "u"."id", "u"."email", "p"."full_name", "u"."username", "p"."gender", "p"."phone_number","p"."profession", "p"."nationality_id", "p"."birth_date"  
	from "users" "u" 
	join "profile" "p"
	on "u"."id" = "p"."user_id"`

	rows, _ := db.Query(
		context.Background(),
		joinSql,
	)

	profile, _ := pgx.CollectRows(rows, pgx.RowToStructByPos[Profile])
	return profile
}

func FindProfileByUserId(id int) []Profile {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, err := db.Query(
		context.Background(),
		`select * from "profile" where "user_id" = $1`, id,
	)
	if err != nil {
		fmt.Println(err)
	}
	profiles, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Profile])

	if err != nil {
		fmt.Errorf("Error")
	}

	return profiles
}

func FindAllProfile() []Profile {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(),
		`select * from "profile" order by "id" asc`,
	)
	profile, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Profile])
	if err != nil {
		fmt.Println(err)
	}
	return profile
}

func UpdateProfile(data Profile, id int) error {
	db := lib.DB()
	defer db.Close(context.Background())

	dataSql := `UPDATE "profile" SET "picture" = $1, "full_name" = $2, "birth_date" = $3, "gender" = $4, "phone_number" = $5, "profession" = $6, "nationality_id" = $7 WHERE "user_id" = $8`
	_, err := db.Exec(context.Background(), dataSql, data.Picture, data.Full_name, data.Birth_date, data.Gender, data.Phone_number, data.Profession, data.Nationality_id, id)
	if err != nil {
		return fmt.Errorf("failed to update profile: %v", err)
	}

	return nil
}

func UpdateProfilePicture(data Profile, id int) (Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE "profile" SET "picture" = $1 WHERE "user_id" = $2 RETURNING id, picture, user_id`

	row, err := db.Query(context.Background(), sql, data.Picture, id)
	if err != nil {
		return Profile{}, nil
	}

	profile, err := pgx.CollectOneRow(row, pgx.RowToStructByName[Profile])
	if err != nil {
		return Profile{}, nil
	}

	return profile, nil
}
