package entity

import "github.com/google/uuid"

// Item表示所有子领域的Item
type Item struct {
	ID          uuid.UUID `json:"id" bson:"id"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
}
