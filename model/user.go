package model

type User struct {
	ID                    interface{} `bson:"_id,omitempty"`
	Username, Displayname string
}

func NewUser(username, displayname string) *User {
	return &User{
		Username:    username,
		Displayname: displayname,
	}
}
