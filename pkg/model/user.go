package model

// User model
type User struct {
	Username    string `bson:"_id"`
	Displayname string
}

// NewUser return a pointer to User
func NewUser(username, displayname string) *User {
	return &User{
		Username:    username,
		Displayname: displayname,
	}
}
