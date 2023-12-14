package customer

import (
	"errors"
	"github.com/google/uuid"
	"github.com/songyue/ddd-go/aggregate"
)

// Customer包保存了客户领域的所有域逻辑

var (
	// 当没有找到客户时返回ErrCustomerNotFound。
	ErrCustomerNotFound = errors.New("the customer was not found in the repository")
	// ErrFailedToAddCustomer在无法将客户添加到存储库时返回。
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	// ErrFailedToAddCustomer在无法将客户添加到存储库时返回。
	ErrUpdateCustomer = errors.New("failed to update the customer in the repository")
)

//	CustomerRepository是一个接口，它定义了围绕客户仓库的规则
//
// 必须实现的函数
type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(customer aggregate.Customer) error
}
