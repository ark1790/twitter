package model

import "time"

// Feed ...
type Feed struct {
	ID        string    `json:"id" bson:"-"`
	Body      string    `json:"body" bson:"bson"`
	For       string    `json:"for" bson:"for"`
	Username  string    `json:"username" bson:"username"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
