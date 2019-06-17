package model

import "time"

// Tweet model
type Tweet struct {
	ID       interface{} `bson:"_id,omitempty"`
	UserID   interface{}
	Likes    int
	Body     string
	DateTime time.Time
}

// NewTweet return a pointer to Tweet
func NewTweet(body string, user *User) *Tweet {
	return &Tweet{
		UserID:   user.ID,
		Body:     body,
		DateTime: time.Now(),
	}
}
