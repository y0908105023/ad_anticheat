package flowhandler

import (
	rpc "thrift/gen-go/filtergen"
)

type Handler interface {
	ProcessRequest(request *rpc.AntiSpamRequest, response *rpc.AntiSpamResponse) *rpc.AntiSpamResponse
	SetNext(next Handler)
	GetNext() Handler
}

type BaseHandler struct {
	combine bool // no speical declare,combine is false,it means one handler is enough
	next    Handler
}

func (handler *BaseHandler) PackNoneSlotResponse(length int, reason string, isLegal bool) *rpc.AntiSpamResponse {
	legals := make([]bool, length)
	reasons := make([]string, length)
	for i := 0; i < length; i++ {
		legals[i] = isLegal
		reasons[i] = reason
	}
	r := &rpc.AntiSpamResponse{legals, reasons}
	return r
}

func (handler *BaseHandler) ProcessRequest(request *rpc.AntiSpamRequest, response *rpc.AntiSpamResponse) *rpc.AntiSpamResponse {
	if handler.GetNext() != nil {
		return handler.GetNext().ProcessRequest(request, response)
	} else {
		return handler.PackNoneSlotResponse(len(request.SlotIds), "0", true)
	}
}

func (handler *BaseHandler) SetNext(next Handler) {
	handler.next = next
}

func (handler *BaseHandler) GetNext() Handler {
	return handler.next
}
