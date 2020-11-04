package model

import "time"

// User ...
type User struct {
	ID        string    `json:"id" bson:"-"`
	Name      string    `json:"name"`
	Username  string    `json:"username" bson:"username"`
	Private   bool      `json:"private" bson:"private"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
