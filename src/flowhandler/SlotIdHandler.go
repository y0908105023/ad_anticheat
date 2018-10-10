package flowhandler

import (
	"strings"
	rpc "thrift/gen-go/filtergen"
)

type SlotIdHandler struct {
	BaseHandler
}

func (slotidhandler *SlotIdHandler) ProcessRequest(request *rpc.AntiSpamRequest, response *rpc.AntiSpamResponse) *rpc.AntiSpamResponse {
	// start := time.Now().UnixNano()
	legals := make([]bool, len(request.SlotIds))
	reasons := make([]string, len(request.SlotIds))

	for index, slotId := range request.SlotIds {
		if strings.TrimSpace(slotId) != "" {
			slotid_value := slotid_map[slotId]
			if slotid_value != "" {
				legals[index] = false
				reasons[index] = slotid_value
			} else {
				legals[index] = true
				reasons[index] = "0"
			}
		}
	}
	// end := time.Now().UnixNano()
	// util.Log.Info("slot time =", (end - start))
	r := &rpc.AntiSpamResponse{legals, reasons}
	/*
	 *  slotid处理的是一个半成品，也就是有非法的和不非法的，还有继续判断其中合法的一部分
	 *  判断这一部分，还需要用另外一个半成品判断
	 */
	if slotidhandler.GetNext() != nil {
		return slotidhandler.GetNext().ProcessRequest(request, r)
	} else {
		//如果没有后继的handler，代表处理完成，并且完全正确
		return r
	}
}
