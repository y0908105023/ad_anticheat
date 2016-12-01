package flowhandler

import (
	"strings"
	rpc "thrift/gen-go/filtergen"
	"tripod/devkit"
)

/*
 * 超频设备id，多天任务分析得到
 */
type MuidOverFrequcentHandler struct {
	BaseHandler
}

func (muidOverHandler *MuidOverFrequcentHandler) ProcessRequest(request *rpc.AntiSpamRequest,
	response *rpc.AntiSpamResponse) *rpc.AntiSpamResponse {

	if strings.TrimSpace(request.ID) != "" {
		isover := muid_frequcent_map[request.ID]
		if isover != "" {
			return muidOverHandler.PackNoneSlotResponse(len(request.SlotIds), isover, false)
		}
		//如果id是原始类型
		if request.IdType == rpc.IDType_origin_id {
			md5id := devkit.HashFunc(request.ID)
			md5_id_value := muid_frequcent_map[md5id]

			if md5_id_value != "" {
				return muidOverHandler.PackNoneSlotResponse(len(request.SlotIds), md5_id_value, false)
			}
		}
	}

	//如果前面请求已经处理，代表是id超频非法，如果没处理，交给下一个处理
	if muidOverHandler.GetNext() != nil {
		return muidOverHandler.GetNext().ProcessRequest(request, response)
	} else if response != nil {
		return response
	} else {
		//如果没有后继的handler，代表处理完成，并且完全正确
		return muidOverHandler.PackNoneSlotResponse(len(request.SlotIds), "0", true)
	}

}
