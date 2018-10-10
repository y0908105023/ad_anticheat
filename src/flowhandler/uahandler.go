package flowhandler

import (
	"strings"
	rpc "thrift/gen-go/filtergen"
)

type UaHandler struct {
	BaseHandler
}

var ua_illegals = [7]string{"spider", "robot", "spyder", "crawler", "superbot", "download", "nutch"}

/*
 * 1 : ua白名单
 * 2 ：ua长度
 * 3 ：关键词过滤
 * 4 ：黑名单
 */
func (uaHandler *UaHandler) ProcessRequest(request *rpc.AntiSpamRequest,
	response *rpc.AntiSpamResponse) *rpc.AntiSpamResponse {
	// start := time.Now().UnixNano()
	if strings.TrimSpace(request.Ua) != "" {

		//mob ua长度
		if request.DevType != 0 && request.DevType != 1 && request.DevType != 7 && len(request.Ua) < 2 && len(request.Ua) > 0 {
			return uaHandler.PackNoneSlotResponse(len(request.SlotIds), "40", false)
		}

		//pc ua长度
		if request.DevType == 0 && len(request.Ua) < 60 && len(request.Ua) > 0 {
			return uaHandler.PackNoneSlotResponse(len(request.SlotIds), "40", false)
		}

		//关键词过滤
		for _, ua_illegal := range ua_illegals {
			if strings.Contains(request.Ua, ua_illegal) {
				if ua_illegal == "spider" && strings.Contains(ua_illegal, "com.") {
					break
				}
				return uaHandler.PackNoneSlotResponse(len(request.SlotIds), "40", false)
			}
		}
		//黑名单
		ua_value := ua_map[request.Ua]
		if ua_value != "" {
			return uaHandler.PackNoneSlotResponse(len(request.SlotIds), ua_value, false)
		}
	}

	// end := time.Now().UnixNano()
	// util.Log.Info("id time =", (end - start))
	//如果前面请求已经处理，代表是id非法，如果没处理，交给下一个处理
	if uaHandler.GetNext() != nil {
		return uaHandler.GetNext().ProcessRequest(request, response)
	} else if response != nil {
		return response
	} else {
		//如果没有后继的handler，代表处理完成，并且完全正确
		return uaHandler.PackNoneSlotResponse(len(request.SlotIds), "0", true)
	}
}
