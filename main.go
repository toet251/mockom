package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"log"
	"net"
	orders "toet.io/mockom/pb"
)


type OMV3MockServer struct {
	orders.OrderCapturingServiceServer
}

func (O OMV3MockServer) InternalScheduleOrderCommand(ctx context.Context, request *orders.InternalScheduleOrderCommandRequest) (*orders.InternalScheduleOrderCommandResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) CaptureOrder(ctx context.Context, request *orders.CaptureOrderRequest) (*orders.CaptureOrderResponse, error) {
	panic("implement me")
}

// Server will return result according to parameter `amount`
func (O OMV3MockServer) RecordPaymentIPN(ctx context.Context, request *orders.RecordPaymentIPNRequest) (*orders.RecordPaymentIPNResponse, error) {
	mapErr := map[int64]map[string]string{
		200: {"code": "200", "msg": "Thành công"},
		410: {"code": "410", "msg": "Sai checksum"},
		411: {"code": "411", "msg": "Dữ liệu không đúng định dạng"},
		421: {"code": "421", "msg": "Xử lý Đơn hàng không thành công <kèm mô tả chi tiết>"},
		500: {"code": "500", "msg": "Lỗi hệ thống"},
		503: {"code": "503", "msg": "Hệ thống đang bảo trì"},
		600: {"code": "600", "msg": "Lỗi ngoài danh mục mô tả"},
	}

	code := request.Amount % 1000

	if err, ok := mapErr[code]; ok {
		resp := &orders.RecordPaymentIPNResponse{
			Code:    err["code"],
			Message: err["msg"],
			TraceId: "trace-id",
		}
		log.Printf("return %#v\n", resp)
		return resp, nil
	}

	log.Fatal("error")

	return nil, errors.New("merchant internal err")
}

func (O OMV3MockServer) ChangeCODPayment(ctx context.Context, request *orders.ChangeCODPaymentRequest) (*orders.ChangeCODPaymentResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) GetOrderById(ctx context.Context, request *orders.GetOrderByIdRequest) (*orders.GetOrderByIdResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) ListOrders(ctx context.Context, request *orders.ListOrdersRequest) (*orders.ListOrdersResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) CancelOrder(ctx context.Context, request *orders.CancelOrderRequest) (*orders.CancelOrderResponse, error) {
	panic("implement me")
}

func main() {
	lis, err := net.Listen("tcp", ":11111")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	omv3Server := &OMV3MockServer{}
	orders.RegisterOrderCapturingServiceServer(s, omv3Server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		log.Println("serve on 11111")
	}
}