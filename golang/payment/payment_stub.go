package payment

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Messages
type CreatePaymentRequest struct {
	UserId     int64
	OrderId    int64
	TotalPrice float32
}

type CreatePaymentResponse struct {
	PaymentId int64
	BillId    int64
}

// Client interface
type PaymentClient interface {
	Create(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
}


type paymentClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentClient(cc grpc.ClientConnInterface) PaymentClient {
	return &paymentClient{cc}
}

func (c *paymentClient) Create(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	out := new(CreatePaymentResponse)
	if err := c.cc.Invoke(ctx, "/payment.Payment/Create", in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// Server API
type PaymentServer interface {
	Create(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
}

type UnimplementedPaymentServer struct{}

func (*UnimplementedPaymentServer) Create(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterPaymentServer(s *grpc.Server, srv PaymentServer) {
	s.RegisterService(&_Payment_serviceDesc, srv)
}

func _Payment_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.Payment/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).Create(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Payment_serviceDesc = grpc.ServiceDesc{
	ServiceName: "payment.Payment",
	HandlerType: (*PaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Payment_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment.proto",
}
