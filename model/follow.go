package model

import "time"

// Follow ...
type Follow struct {
	ID        string    `json:"id" bson:"-"`
	Username  string    `json:"username" bson:"username"`
	Profile   string    `json:"profile" bson:"profile"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
