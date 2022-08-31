package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Age       int
	Gender    string
	Born      string
	CreatedAt time.Time `bson:"createdAt"`
}
