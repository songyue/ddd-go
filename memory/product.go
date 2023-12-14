package memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/songyue/ddd-go/aggregate"
	"github.com/songyue/ddd-go/domain/product"
	"sync"
)

// memory包是客户仓库的内存中实现

// ProductMemoryRepository 实现了 ProductMemoryRepository 接口
type ProductMemoryRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

// NewProductRep 是一个工厂函数，用于生成新的商品仓库
func NewProductRep() *ProductMemoryRepository {
	return &ProductMemoryRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

// Get 根据ID查找 Product
func (mr *ProductMemoryRepository) Get(id uuid.UUID) (aggregate.Product, error) {
	if c, ok := mr.products[id]; ok {
		return c, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

// Gets 根据IDs查找 Products
func (mr *ProductMemoryRepository) Gets(ids []uuid.UUID) ([]aggregate.Product, error) {

	products := make([]aggregate.Product, 0)
	for _, id := range ids {
		newProduct, err := mr.Get(id)
		if err != nil {
			// todo:sy 打印错误？
			continue
		}
		products = append(products, newProduct)
	}
	if len(products) > 0 {
		return products, nil
	}

	return []aggregate.Product{}, product.ErrProductNotFound
}

// Add 将向存储库添加一个新 Product
func (mr *ProductMemoryRepository) Add(c aggregate.Product) error {
	if mr.products == nil {
		mr.Lock()
		mr.products = make(map[uuid.UUID]aggregate.Product)
		mr.Unlock()
	}
	// 确保 Product 不在仓库中
	if _, ok := mr.products[c.GetID()]; ok {
		return fmt.Errorf("product already exists: %w", product.ErrFailedToAddProduct)
	}
	mr.Lock()
	mr.products[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Update 将用新的商品信息替换现有的 Product 信息
func (mr *ProductMemoryRepository) Update(c aggregate.Product) error {
	// 确保 Product 在存储库中
	if _, ok := mr.products[c.GetID()]; !ok {
		return fmt.Errorf("product does not exist: %w", product.ErrProductNotFound)
	}
	mr.Lock()
	mr.products[c.GetID()] = c
	mr.Unlock()
	return nil
}
