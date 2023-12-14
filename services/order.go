package services

import (
	"github.com/google/uuid"
	"github.com/songyue/ddd-go/aggregate"
	"github.com/songyue/ddd-go/domain/customer"
	"github.com/songyue/ddd-go/domain/order"
	"github.com/songyue/ddd-go/domain/product"
	"github.com/songyue/ddd-go/memory"
)

// service包，包含将仓库连接到业务流的所有服务

// OrderConfiguration 是一个函数的别名，该函数将接受一个指向OrderService的指针并对其进行修改
type OrderConfiguration func(os *OrderService) error

// OrderService 是OrderService的一个实现
type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
	orders    order.OrderRepository
}

// NewOrderService 接受可变数量的OrderConfiguration函数，并返回一个新的OrderService
// 将按照传入的顺序调用每个OrderConfiguration
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	// 创建orderService
	os := &OrderService{}
	// 应用所有传入的Configurations
	for _, cfg := range cfgs {
		// 将service传递到configuration函数
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository 将给定的客户仓库应用到OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// 返回一个与OrderConfiguration别名匹配的函数，
	// 您需要返回这个，以便父函数可以接受所有需要的参数
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

// WithMemoryCustomerRepository 将内存客户仓库应用到OrderService
func WithMemoryCustomerRepository() OrderConfiguration {
	// 创建内存仓库，如果我们需要参数，如连接字符串，它们可以在这里输入
	cr := memory.New()
	return WithCustomerRepository(cr)
}

// WithProductRepository 将给定的商品仓库应用到OrderService
func WithProductRepository(cr product.ProductRepository) OrderConfiguration {
	// 返回一个与OrderConfiguration别名匹配的函数，
	// 您需要返回这个，以便父函数可以接受所有需要的参数
	return func(os *OrderService) error {
		os.products = cr
		return nil
	}
}

// WithMemoryProductRepository 将内存商品仓库应用到OrderService
func WithMemoryProductRepository() OrderConfiguration {
	// 创建内存仓库，如果我们需要参数，如连接字符串，它们可以在这里输入
	cr := memory.NewProductRep()
	return WithProductRepository(cr)
}

// WithOrderRepository 将给定的订单仓库应用到OrderService
func WithOrderRepository(cr order.OrderRepository) OrderConfiguration {
	// 返回一个与OrderConfiguration别名匹配的函数，
	// 您需要返回这个，以便父函数可以接受所有需要的参数
	return func(os *OrderService) error {
		os.orders = cr
		return nil
	}
}

// WithMemoryOrderRepository 将内存订单仓库应用到OrderService
func WithMemoryOrderRepository() OrderConfiguration {
	// 创建内存仓库，如果我们需要参数，如连接字符串，它们可以在这里输入
	cr := memory.NewOrderRep()
	return WithOrderRepository(cr)
}

// CreateOrder 将所有仓库链接在一起，为客户创建订单
func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) error {
	// 获取customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return err
	}
	// 获取每个产品，我们需要一个Product Repository
	p, err := o.products.Gets(productIDs)
	if err != nil {
		return err
	}
	newOrder, err := aggregate.NewOrder(c, p)
	if err != nil {
		return err
	}
	err = o.orders.Add(newOrder)
	if err != nil {
		return err
	}

	return nil
}
