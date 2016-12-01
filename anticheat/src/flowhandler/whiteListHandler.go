package flowhandler

import (
	"strings"
	rpc "thrift/gen-go/filtergen"
)

type WhiteListHandler struct {
	BaseHandler
}

var sdks = map[string]bool{
	"motv":     true,
	"innersdk": true,
	"motvapi":  true,
}

/*
 * 白名单
 */
func (whiteHandler *WhiteListHandler) ProcessRequest(request *rpc.AntiSpamRequest, response *rpc.AntiSpamResponse) *rpc.AntiSpamResponse {
	// start := time.Now().UnixNano()
	if strings.TrimSpace(request.SourceId) != "" {
		if sdks[request.SourceId] {
			return whiteHandler.PackNoneSlotResponse(len(request.SlotIds), "0", true)
		}
	}
	// end := time.Now().UnixNano()
	// util.Log.Info("white time =", (end - start))
	if whiteHandler.GetNext() != nil {
		return whiteHandler.GetNext().ProcessRequest(request, response)
	} else if response != nil {
		return response
	} else {
		return whiteHandler.PackNoneSlotResponse(len(request.SlotIds), "0", true)
	}
}
