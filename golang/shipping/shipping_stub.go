package shipping

import (
	"context"

	"google.golang.org/grpc"
)

// Messages
type Item struct {
	ProductCode string
	Quantity    int32
}

type EstimateRequest struct {
	OrderId int64
	Items   []*Item
}

type EstimateResponse struct {
	DeliveryDays int32
}

func (r *EstimateResponse) GetDeliveryDays() int32 { return r.DeliveryDays }

// Client interface
type ShippingClient interface {
	EstimateDelivery(ctx context.Context, in *EstimateRequest, opts ...grpc.CallOption) (*EstimateResponse, error)
}


type shippingClient struct{ cc grpc.ClientConnInterface }

func NewShippingClient(cc grpc.ClientConnInterface) ShippingClient { return &shippingClient{cc} }

func (c *shippingClient) EstimateDelivery(ctx context.Context, in *EstimateRequest, opts ...grpc.CallOption) (*EstimateResponse, error) {
	return nil, nil
}
