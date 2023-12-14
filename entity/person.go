package entity

import (
	"github.com/google/uuid"
)

// Person 在所有领域中代表人
type Person struct {
	// ID是实体的标识符，该ID为所有子领域共享
	ID uuid.UUID `json:"id" bson:"id"`
	//Name就是人的名字
	Name string `json:"name" bson:"name"`
	// 人的年龄
	Age int `json:"age" bson:"age"`
}
