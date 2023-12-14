package order

import (
	"errors"
	"github.com/google/uuid"
	"github.com/songyue/ddd-go/aggregate"
)

// Order 包保存了订单领域的所有域逻辑

var (
	// 当没有找到商品时返回ErrCustomerNotFound。
	ErrOrderNotFound = errors.New("the Order was not found in the repository")
	// ErrFailedToAddOrder在无法将商品添加到存储库时返回。
	ErrFailedToAddOrder = errors.New("failed to add the Order to the repository")
	// ErrFailedToAddOrder在无法将商品添加到存储库时返回。
	ErrUpdateOrder = errors.New("failed to update the Order in the repository")
)

//	OrderRepository是一个接口，它定义了围绕商品仓库的规则
//
// 必须实现的函数
type OrderRepository interface {
	Get(uuid.UUID) (aggregate.Order, error)
	Gets([]uuid.UUID) ([]aggregate.Order, error)
	Add(aggregate.Order) error
	Update(Order aggregate.Order) error
}
