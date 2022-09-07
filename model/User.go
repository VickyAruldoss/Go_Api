package model

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

var users []User = []User{
	{Name: "Vicky", ID: 1, Email: "vicky@mail.com", Phone: "9834734683", Address: "chennai"},
	{Name: "Aruldoss", ID: 2, Email: "vicky@mail.com", Phone: "9834734683", Address: "chennai"}}

func GetAllUsers() ([]User, error) {
	return users, nil
}

func CreateUser(user User) {
	users = append(users, user)
}

func GetUserById(id int) User {
	var user User
	for _, v := range users {
		if v.ID == id {
			user = v
			break
		}
	}
	return user
}
