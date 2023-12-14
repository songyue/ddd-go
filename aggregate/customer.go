package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/songyue/ddd-go/entity"
	"github.com/songyue/ddd-go/valueObject"
)

var (
	// 当person在newcustom工厂中无效时返回ErrInvalidPerson
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

type Customer struct {
	// Person是客户的根实体
	// person.ID是聚合的主标识符
	Person *entity.Person `bson:"person"`
	//一个客户可以持有许多产品
	Products []*entity.Item `bson:"products"`
	// 一个客户可以执行许多事务
	Transactions []valueObject.Transaction `bson:"transactions"`
}

// NewCustomer 是创建新的Customer聚合的工厂
// 它将验证名称是否为空
func NewCustomer(name string) (Customer, error) {
	// 验证name不是空的
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	// 创建一个新person并生成ID
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	// 创建一个customer对象并初始化所有的值以避免空指针异常
	return Customer{
		Person:       person,
		Products:     make([]*entity.Item, 0),
		Transactions: make([]valueObject.Transaction, 0),
	}, nil
}

// GetID 返回客户的根实体ID
func (c *Customer) GetID() uuid.UUID {
	return c.Person.ID
}

// SetName 更改客户的名称
func (c *Customer) SetName(name string) {
	c.Person.Name = name
}
