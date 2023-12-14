package aggregate_test

import (
	"testing"

	"github.com/songyue/ddd-go/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	// 构建我们需要的测试用例数据结构
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	//创建新的测试用例
	testCases := []testCase{
		{
			test:        "Empty Name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		}, {
			test:        "Valid Name",
			name:        "Percy Bolmer",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		// Run Tests
		t.Run(tc.test, func(t *testing.T) {
			//创建新的customer
			_, err := aggregate.NewCustomer(tc.name)
			//检查错误是否与预期的错误匹配
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})

	}

}
