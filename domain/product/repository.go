package product

import (
	"errors"
	"github.com/google/uuid"
	"github.com/songyue/ddd-go/aggregate"
)

// Customer包保存了客户领域的所有域逻辑

var (
	// 当没有找到商品时返回ErrCustomerNotFound。
	ErrProductNotFound = errors.New("the Product was not found in the repository")
	// ErrFailedToAddProduct在无法将商品添加到存储库时返回。
	ErrFailedToAddProduct = errors.New("failed to add the Product to the repository")
	// ErrFailedToAddProduct在无法将商品添加到存储库时返回。
	ErrUpdateProduct = errors.New("failed to update the Product in the repository")
)

//	ProductRepository是一个接口，它定义了围绕商品仓库的规则
//
// 必须实现的函数 todo:sy 待实现
type ProductRepository interface {
	Get(uuid.UUID) (aggregate.Product, error)
	Gets([]uuid.UUID) ([]aggregate.Product, error)
	Add(aggregate.Product) error
	Update(Product aggregate.Product) error
}
