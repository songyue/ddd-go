package aggregate

import (
	"errors"
	"github.com/google/uuid"
)

var (
	// 当customer在NewOrder工厂中无效时返回ErrInvalidCustomer
	ErrInvalidCustomer = errors.New("a order has to have an valid customer")
)

type Order struct {
	ID       uuid.UUID
	Customer Customer
	Products []Product
	Address  string // todo:sy 值对象
	// todo:sy 收货地址 等
}

func NewOrder(customer Customer, products []Product) (Order, error) {

	// todo 基础验证
	return Order{
		ID:       uuid.New(),
		Customer: customer,
		Products: products,
	}, nil
}

func (o *Order) GetID() uuid.UUID {
	return o.ID
}

func (o *Order) SetAddress(address string) {
	o.Address = address // todo:sy 值对象
}
