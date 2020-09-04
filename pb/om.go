package orders

import (
	"context"
	"errors"
	"log"
)

type OMV3MockServer struct {
	RecordInternalPaymentIPNResponse
}

func (O OMV3MockServer) CaptureInternalOrder(ctx context.Context, request *CaptureInternalOrderRequest) (*CaptureInternalOrderResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) CreateInternalOrder(ctx context.Context, request *CreateInternalOrderRequest) (*CreateInternalOrderResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) ChangeInternalCODPayment(ctx context.Context, request *ChangeInternalCODPaymentRequest) (*ChangeInternalCODPaymentResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) GetInternalOrderByCode(ctx context.Context, request *GetInternalOrderByCodeRequest) (*GetInternalOrderByCodeResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) ListInternalOrders(ctx context.Context, request *ListInternalOrdersRequest) (*ListInternalOrdersResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) CancelInternalOrder(ctx context.Context, request *CancelInternalOrderRequest) (*CancelInternalOrderResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) RecordInternalPaymentIPN(ctx context.Context, request *RecordInternalPaymentIPNRequest) (*RecordInternalPaymentIPNResponse, error) {
	mapErr := map[int64]map[string]string{
				200: {"code": "200", "msg": "Thành công"},
		410:         {"code": "410", "msg": "Sai checksum"},
		411:         {"code": "411", "msg": "Dữ liệu không đúng định dạng"},
		421:         {"code": "421", "msg": "Xử lý Đơn hàng không thành công ...kèm mô tả chi tiết..."},
		500:         {"code": "500", "msg": "Lỗi hệ thống"},
		503:         {"code": "503", "msg": "Hệ thống đang bảo trì"},
		600:         {"code": "600", "msg": "Lỗi ngoài danh mục mô tả"},

		412: {"code": "412", "msg": "ABC XYZ"},
	}

	code := request.Amount % 1000

	log.Printf("Receive: %#v\n", request)
	if err, ok := mapErr[code]; ok {
		resp := &RecordInternalPaymentIPNResponse{
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

func (O OMV3MockServer) ScheduleInternalOrderCommand(ctx context.Context, request *ScheduleInternalOrderCommandRequest) (*ScheduleInternalOrderCommandResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) CaptureInternalOrderFromCart(ctx context.Context, request *CaptureInternalOrderFromCartRequest) (*CaptureInternalOrderFromCartResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) UpdateInternalShippingInfo(ctx context.Context, request *UpdateInternalShippingInfoRequest) (*UpdateInternalShippingInfoResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) UpdateInternalBillingInfo(ctx context.Context, request *UpdateInternalBillingInfoRequest) (*UpdateInternalBillingInfoResponse, error) {
	panic("implement me")
}

func (O OMV3MockServer) mustEmbedUnimplementedOrderCapturingServiceServer() {
	panic("implement me")
}

func NewServerImpl() *OMV3MockServer {
	return &OMV3MockServer{}
}
