package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"min=8"`
}

var dataUser = []User{
	{Id: 1, Name: "Admin", Email: "admin@mail.com", Password: "1234"},
}

func GetAllUsers() []User {
	data := dataUser

	return data
}

func GetOneUser(id int) User {
	data := dataUser

	user := User{}
	for _, item := range data {
		if id == item.Id {
			user = item
		}
	}

	return user
}

func CreateUser(data User) User {
	id := 0
	for _, v := range dataUser {
		id = v.Id
	}

	data.Id = id + 1
	dataUser = append(dataUser, data)

	return data
}

func RemoveData(id int) User {
	index := -1
	userDelete := User{}
	for idx, item := range dataUser {
		if item.Id == id {
			index = idx
			userDelete = item
		}
	}
	if userDelete.Id != 0 {
		dataUser = append(dataUser[:index], dataUser[index+1:]...)
	}

	return userDelete
}

func EditData(data User, id int) User {

	idx := -1

	for index, item := range dataUser {
		if id == item.Id {
			idx = index
		}
	}

	if idx == 0 {
		dataUser[idx].Name = data.Name
		dataUser[idx].Email = data.Email
		dataUser[idx].Password = data.Password
		data.Id = dataUser[idx].Id
	}

	return data
}
