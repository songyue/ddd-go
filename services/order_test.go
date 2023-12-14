package services

import (
	"github.com/google/uuid"
	"github.com/songyue/ddd-go/aggregate"
	//"github.com/songyue/ddd-go/domain/customer"
	//"github.com/songyue/ddd-go/domain/order"
	//"github.com/songyue/ddd-go/domain/product"
	"testing"
)

func TestOrderService_CreateOrder(t *testing.T) {
	orderService, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryOrderRepository(),
		WithMemoryProductRepository(),
	)
	if err != nil {
		t.Errorf("NewOrderService() error = %v", err)
	}

	//  创建用户
	newCustomer, err := aggregate.NewCustomer("Jack")
	if err != nil {
		t.Errorf("NewCustomer() error = %v", err)
		return
	}
	err = orderService.customers.Add(newCustomer)
	if err != nil {
		t.Errorf("add customer error = %v", err)
		return
	}

	//  创建商品
	oneProduct, err := aggregate.NewProduct("WaHaHa")
	if err != nil {
		t.Errorf("NewProduct() error = %v", err)
		return
	}
	err = orderService.products.Add(oneProduct)
	if err != nil {
		t.Errorf("add product error = %v", err)
		return
	}

	productIDs := make([]uuid.UUID, 0)
	productIDs = append(productIDs, oneProduct.GetID())

	type fields struct {
		o *OrderService
	}
	type args struct {
		customerID uuid.UUID
		productIDs []uuid.UUID
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "创建订单",
			fields: fields{
				o: orderService,
			},
			args: args{
				customerID: newCustomer.GetID(),
				productIDs: productIDs,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := tt.fields.o.CreateOrder(tt.args.customerID, tt.args.productIDs); (err != nil) != tt.wantErr {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
