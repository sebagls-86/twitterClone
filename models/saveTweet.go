package models

import "time"

type SaveTweet struct {
	UserID  string    `bson:"userid" json:"id,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
