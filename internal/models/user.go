package models

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Mock function (replace with DB logic)
func GetUserByID(id int64) (*User, error) {
	return &User{ID: id, Name: "John Doe", Email: "john@example.com"}, nil
}
