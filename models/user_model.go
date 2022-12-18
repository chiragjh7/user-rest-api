package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Schema for the user model
type User struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Dob         string             `json:"dob,omitempty"`
	Address     string             `json:"address,omitempty"`
	Description string             `json:"description,omitempty"`
	CreatedAt   primitive.DateTime          `json:"createdAt,omitempty"`
	UpdatedAt   primitive.DateTime          `json:"updatedAt,omitempty"`
}

