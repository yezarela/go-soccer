package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Player represents player model
type Player struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Nickname  string             `json:"nickname"`
	Position  string             `json:"position"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
