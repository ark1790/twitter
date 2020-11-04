package model

import "time"

// Tweet ...
type Tweet struct {
	ID        string    `json:"id" bson:"-"`
	Body      string    `json:"body" bson:"bson"`
	Username  string    `json:"username" bson:"username"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
