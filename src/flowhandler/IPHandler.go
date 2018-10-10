package flowhandler

import (
	"strings"
	rpc "thrift/gen-go/filtergen"
)

type IpHandler struct {
	BaseHandler
}

func (iphandler *IpHandler) ProcessRequest(request *rpc.AntiSpamRequest, response *rpc.AntiSpamResponse) *rpc.AntiSpamResponse {
	// start := time.Now().UnixNano()
	if strings.TrimSpace(request.IP) != "" {
		ip_value := ip_map[request.IP]
		if ip_value != "" {
			return iphandler.PackNoneSlotResponse(len(request.SlotIds), ip_value, false)
		}
	}
	// end := time.Now().UnixNano()
	// util.Log.Info("ip time =", (end - start))

	//如果前面请求已经处理，代表是ip非法，如果没处理，交给下一个处理
	if iphandler.GetNext() != nil {
		return iphandler.GetNext().ProcessRequest(request, response)
	} else if response != nil {
		return response
	} else {
		//如果没有后继的handler，代表处理完成，并且完全正确
		return iphandler.PackNoneSlotResponse(len(request.SlotIds), "0", true)
	}
}
