package memory

import (
	"errors"
	"github.com/google/uuid"
	"github.com/songyue/ddd-go/aggregate"
	"github.com/songyue/ddd-go/domain/product"
	"reflect"
	"testing"
)

func TestProductRepository_Add(t *testing.T) {
	type testCase struct {
		name        string
		productName string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Product",
			productName: "Percy",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := ProductMemoryRepository{
				products: map[uuid.UUID]aggregate.Product{},
			}

			newProduct, err := aggregate.NewProduct(tc.productName)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(newProduct)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(newProduct.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != newProduct.GetID() {
				t.Errorf("Expected %v, got %v", newProduct.GetID(), found.GetID())
			}
		})
	}
}

func TestProductMemoryRepository_Get(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	//创建要添加到存储库中的模拟 Product
	newProduct, err := aggregate.NewProduct("Percy")
	if err != nil {
		t.Fatal(err)
	}
	id := newProduct.GetID()
	// 创建要使用的仓库，并添加一些测试数据进行测试
	// 跳过工厂
	repo := ProductMemoryRepository{
		products: map[uuid.UUID]aggregate.Product{
			id: newProduct,
		},
	}

	testCases := []testCase{
		{
			name:        "No Product By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: product.ErrProductNotFound,
		}, {
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestNewProductRep(t *testing.T) {
	tests := []struct {
		name string
		want *ProductMemoryRepository
	}{
		{
			name: "New MemoryRepository",
			want: &ProductMemoryRepository{
				products: make(map[uuid.UUID]aggregate.Product),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductRep(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
