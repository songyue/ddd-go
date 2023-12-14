package memory

import (
	"errors"
	"github.com/google/uuid"
	"github.com/songyue/ddd-go/aggregate"
	"github.com/songyue/ddd-go/domain/order"
	"reflect"
	"testing"
)

func TestOrderRepository_Add(t *testing.T) {
	type testCase struct {
		name        string
		orderName   string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Order",
			orderName:   "Percy",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := OrderMemoryRepository{
				orders: map[uuid.UUID]aggregate.Order{},
			}

			oneCustomer, err := aggregate.NewCustomer("jack")
			if err != nil {
				t.Errorf("Failed to create customer %v", err)
			}

			oneProduct, err := aggregate.NewProduct("WanHaHa")
			if err != nil {
				t.Errorf("Failed to create product %v", err)
			}

			products := make([]aggregate.Product, 0)
			products = append(products, oneProduct)
			newOrder, err := aggregate.NewOrder(oneCustomer, products)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(newOrder)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(newOrder.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != newOrder.GetID() {
				t.Errorf("Expected %v, got %v", newOrder.GetID(), found.GetID())
			}
		})
	}
}

func TestOrderMemoryRepository_Get(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	oneCustomer, err := aggregate.NewCustomer("jack")
	if err != nil {
		t.Errorf("Failed to create customer %v", err)
	}

	oneProduct, err := aggregate.NewProduct("WanHaHa")
	if err != nil {
		t.Errorf("Failed to create product %v", err)
	}

	products := make([]aggregate.Product, 0)
	products = append(products, oneProduct)
	newOrder, err := aggregate.NewOrder(oneCustomer, products)
	if err != nil {
		t.Fatal(err)
	}
	id := newOrder.GetID()
	// 创建要使用的仓库，并添加一些测试数据进行测试
	// 跳过工厂
	repo := OrderMemoryRepository{
		orders: map[uuid.UUID]aggregate.Order{
			id: newOrder,
		},
	}

	testCases := []testCase{
		{
			name:        "No Order By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: order.ErrOrderNotFound,
		}, {
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestNewOrderRep(t *testing.T) {
	tests := []struct {
		name string
		want *OrderMemoryRepository
	}{
		{
			name: "New MemoryRepository",
			want: &OrderMemoryRepository{
				orders: make(map[uuid.UUID]aggregate.Order),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderRep(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
