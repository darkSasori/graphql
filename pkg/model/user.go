package model

// User model
type User struct {
	ID                    interface{} `bson:"_id,omitempty"`
	Username, Displayname string
}

// NewUser return a pointer to User
func NewUser(username, displayname string) *User {
	return &User{
		Username:    username,
		Displayname: displayname,
	}
}
