package orders

import (
	"context"
	"errors"
	"log"
)

type OMV3MockServer struct {
	OrderCapturingServiceServer
}

func NewServerImpl() *OMV3MockServer {
	return &OMV3MockServer{}
}

func (O OMV3MockServer) PrivateCaptureOrder(ctx context.Context, request *PrivateCaptureOrderRequest) (*PrivateCaptureOrderResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) PrivateChangeCODPayment(ctx context.Context, request *PrivateChangeCODPaymentRequest) (*PrivateChangeCODPaymentResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) PrivateGetOrderByCode(ctx context.Context, request *PrivateGetOrderByCodeRequest) (*PrivateGetOrderByCodeResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) PrivateListOrders(ctx context.Context, request *PrivateListOrdersRequest) (*PrivateListOrdersResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) PrivateCancelOrder(ctx context.Context, request *PrivateCancelOrderRequest) (*PrivateCancelOrderResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) InternalGetOrderByCode(ctx context.Context, request *InternalGetOrderByCodeRequest) (*InternalGetOrderByCodeResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) InternalListOrders(ctx context.Context, request *InternalListOrdersRequest) (*InternalListOrdersResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) InternalScheduleOrderCommand(ctx context.Context, request *InternalScheduleOrderCommandRequest) (*InternalScheduleOrderCommandResponse, error) {
	panic("implement me")
}

// Server will return result according to parameter `amount`
func (O OMV3MockServer) InternalRecordPaymentIPN(ctx context.Context, request *InternalRecordPaymentIPNRequest) (*InternalRecordPaymentIPNResponse, error) {
	mapErr := map[int64]map[string]string{
		200: {"code": "200", "msg": "Thành công"},
		410: {"code": "410", "msg": "Sai checksum"},
		411: {"code": "411", "msg": "Dữ liệu không đúng định dạng"},
		421: {"code": "421", "msg": "Xử lý Đơn hàng không thành công <kèm mô tả chi tiết>"},
		500: {"code": "500", "msg": "Lỗi hệ thống"},
		503: {"code": "503", "msg": "Hệ thống đang bảo trì"},
		600: {"code": "600", "msg": "Lỗi ngoài danh mục mô tả"},

		412: {"code": "412", "msg": "ABC XYZ"},
	}

	code := request.Amount % 1000

	log.Printf("Receive: %#v\n", request)
	if err, ok := mapErr[code]; ok {
		resp := &InternalRecordPaymentIPNResponse{
			Code:    err["code"],
			Message: err["msg"],
			TraceId: "trace-id",
		}
		log.Printf("Return %#v\n", resp)
		return resp, nil
	}

	log.Printf("Return error")

	return nil, errors.New("merchant internal err")
}

func (o OMV3MockServer) mustEmbedUnimplementedOrderCapturingServiceServer() {}
