package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Person Model
type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
	City  string             `json:"city"`
	Age   int                `json:"age"`
}
