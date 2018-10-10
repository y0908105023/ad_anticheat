package flowhandler

import (
	// util "hashutil"
	"strings"
	rpc "thrift/gen-go/filtergen"
	// "time"
	"tripod/devkit"
)

type IdHandler struct {
	BaseHandler
}

func (idhandler *IdHandler) ProcessRequest(request *rpc.AntiSpamRequest, response *rpc.AntiSpamResponse) *rpc.AntiSpamResponse {
	// start := time.Now().UnixNano()
	if strings.TrimSpace(request.ID) != "" {
		isover := OverFrequent.Get(request.ID)
		if isover != 0 {
			return idhandler.PackNoneSlotResponse(len(request.SlotIds), "100", false)
		}
		original_id_value := muid_map[request.ID]
		if original_id_value != "" {
			return idhandler.PackNoneSlotResponse(len(request.SlotIds), original_id_value, false)
		}
		//如果id是原始类型
		if request.IdType == rpc.IDType_origin_id {
			md5id := devkit.HashFunc(request.ID)
			md5_id_value := muid_map[md5id]
			over_md5_id_value := OverFrequent.Get(md5id)

			if md5_id_value != "" {
				return idhandler.PackNoneSlotResponse(len(request.SlotIds), md5_id_value, false)
			}
			if over_md5_id_value != 0 {
				return idhandler.PackNoneSlotResponse(len(request.SlotIds), "100", false)
			}
		}
	}
	// end := time.Now().UnixNano()
	// util.Log.Info("id time =", (end - start))
	//如果前面请求已经处理，代表是id非法，如果没处理，交给下一个处理
	if idhandler.GetNext() != nil {
		return idhandler.GetNext().ProcessRequest(request, response)
	} else if response != nil {
		return response
	} else {
		//如果没有后继的handler，代表处理完成，并且完全正确
		return idhandler.PackNoneSlotResponse(len(request.SlotIds), "0", true)
	}

}
