package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Team represents team model
type Team struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Location    string             `json:"location"`
	Players     []Player           `json:"players"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}
