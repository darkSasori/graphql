package model

// User model
type User struct {
	Username    string `bson:"_id"`
	Displayname string
	Image       string
}

// NewUser return a pointer to User
func NewUser(username, displayname, image string) *User {
	return &User{
		Username:    username,
		Displayname: displayname,
		Image:       image,
	}
}
