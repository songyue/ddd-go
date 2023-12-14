package valueObject

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Transaction struct {
	Amount    int       `json:"amount" bson:"amount"`
	From      uuid.UUID `json:"from" bson:"from"`
	To        uuid.UUID `json:"to" bson:"to"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
