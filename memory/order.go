package memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/songyue/ddd-go/aggregate"
	"github.com/songyue/ddd-go/domain/order"
	"sync"
)

// memory包是客户仓库的内存中实现

// OrderMemoryRepository 实现了 OrderMemoryRepository 接口
type OrderMemoryRepository struct {
	orders map[uuid.UUID]aggregate.Order
	sync.Mutex
}

// NewOrderRep 是一个工厂函数，用于生成新的商品仓库
func NewOrderRep() *OrderMemoryRepository {
	return &OrderMemoryRepository{
		orders: make(map[uuid.UUID]aggregate.Order),
	}
}

// Get 根据ID查找 Order
func (mr *OrderMemoryRepository) Get(id uuid.UUID) (aggregate.Order, error) {
	if c, ok := mr.orders[id]; ok {
		return c, nil
	}
	return aggregate.Order{}, order.ErrOrderNotFound
}

// Gets 根据IDs查找 Orders
func (mr *OrderMemoryRepository) Gets(ids []uuid.UUID) ([]aggregate.Order, error) {

	orders := make([]aggregate.Order, 0)
	for _, id := range ids {
		newOrder, err := mr.Get(id)
		if err != nil {
			// todo:sy 打印错误？
			continue
		}
		orders = append(orders, newOrder)
	}
	if len(orders) > 0 {
		return orders, nil
	}

	return []aggregate.Order{}, order.ErrOrderNotFound
}

// Add 将向存储库添加一个新 Order
func (mr *OrderMemoryRepository) Add(c aggregate.Order) error {
	if mr.orders == nil {
		mr.Lock()
		mr.orders = make(map[uuid.UUID]aggregate.Order)
		mr.Unlock()
	}
	// 确保 Order 不在仓库中
	if _, ok := mr.orders[c.GetID()]; ok {
		return fmt.Errorf("order already exists: %w", order.ErrFailedToAddOrder)
	}
	mr.Lock()
	mr.orders[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Update 将用新的商品信息替换现有的 Order 信息
func (mr *OrderMemoryRepository) Update(c aggregate.Order) error {
	// 确保 Order 在存储库中
	if _, ok := mr.orders[c.GetID()]; !ok {
		return fmt.Errorf("order does not exist: %w", order.ErrOrderNotFound)
	}
	mr.Lock()
	mr.orders[c.GetID()] = c
	mr.Unlock()
	return nil
}
