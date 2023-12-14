package aggregate

import (
	"errors"
	"github.com/google/uuid"
)

var (
	// 当name在NewProduct工厂中无效时返回ErrInvalidProduct
	ErrInvalidProduct = errors.New("a product has to have an valid name")
)

type Product struct {
	ID   uuid.UUID
	Name string `bson:"name"`
}

func NewProduct(name string) (Product, error) {
	// 验证name不是空的
	if name == "" {
		return Product{}, ErrInvalidProduct
	}

	// todo:sy 其他属性
	//product := &entity.Product{
	//
	//}

	return Product{
		ID:   uuid.New(),
		Name: name,
	}, nil

}

// GetID 返回商品的根实体ID
func (c *Product) GetID() uuid.UUID {
	return c.ID
}

// SetName 更改商品的名称
func (c *Product) SetName(name string) {
	c.Name = name
}
