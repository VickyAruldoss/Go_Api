package model

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	// UN         sql.NullString `json:"user_num" swaggertype:"string"`
}

func GetAllUsers() (User, error) {
	return User{
		ID:      1,
		Name:    "vicky",
		Email:   "Myemail",
		Phone:   "98349834",
		Address: "",
	}, nil
}
