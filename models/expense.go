package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Expense struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Title    string             `bson:"title" json:"title"`
	Amount   float64            `bson:"amount" json:"amount"`
	Category string             `bson:"category" json:"category"`
	Date     time.Time          `bson:"date" json:"date"`
	Note     string             `bson:"note" json:"note"`
}
