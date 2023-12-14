package memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/songyue/ddd-go/aggregate"
	"github.com/songyue/ddd-go/domain/customer"
	"sync"
)

// memory包是客户仓库的内存中实现

// MemoryRepository 实现了CustomerRepository接口
type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

// New 是一个工厂函数，用于生成新的客户仓库
func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

// Get 根据ID查找Customer
func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if c, ok := mr.customers[id]; ok {
		return c, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

// Add 将向存储库添加一个新Customer
func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	// 确保Customer不在仓库中
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Update 将用新的客户信息替换现有的Customer信息
func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	// 确保Customer在存储库中
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrCustomerNotFound)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
