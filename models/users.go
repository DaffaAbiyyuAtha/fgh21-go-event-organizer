package models

type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"-" form:"password" binding:"required,min=8"`
}

var dataUser = []Users{
	{Id: 1, Name: "Admin", Email: "admin@mail.com", Password: "1234"},
}

func GetAllUsers() []Users {
	data := dataUser

	return data
}

func GetOneUserById(id int) Users {
	data := dataUser

	user := Users{}
	for _, item := range data {
		if id == item.Id {
			user = item
		}
	}

	return user
}

func CreateUser(data Users) Users {
	id := 0
	for _, ids := range dataUser {
		id = ids.Id
	}

	data.Id = id + 1
	dataUser = append(dataUser, data)

	return data
}

func DeleteDataById(id int) Users {
	index := -1
	userDelete := Users{}
	for ids, item := range dataUser {
		if item.Id == id {
			index = ids
			userDelete = item
		}
	}
	if userDelete.Id != 0 {
		dataUser = append(dataUser[:index], dataUser[index+1:]...)
	}

	return userDelete
}

func UpdateDataById(data Users, id int) Users {

	ids := -1

	for index, item := range dataUser {
		if id == item.Id {
			ids = index
		}
	}

	if ids == 0 {
		dataUser[ids].Name = data.Name
		dataUser[ids].Email = data.Email
		dataUser[ids].Password = data.Password
		data.Id = dataUser[ids].Id
	}

	return data
}
