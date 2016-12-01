package flowhandler

import (
	util "hashutil"
	"strings"
)

const (
	ID_HANDLER_CONST     = "4"   //id过滤
	IP_HANDLER_CONST     = "2"   //ip过滤
	SLOTID_HANDLER_CONST = "1"   //广告位过滤
	LAN_HANDLER_CONST    = "32"  //局域网过滤
	IDC_HANDLER_CONST    = "128" //idc机房过滤
	UA_HANDLER_CONST     = "64"  //idc机房过滤
)

func ReBuildStrategy() {

	util.Log.Info("reload strategy ....")
	var handler_list []Handler
	//目前value是用占位符，存储的是是否启用
	for key, _ := range strategy_map {
		for _, flag := range strings.Split(key, ",") {
			if flag == ID_HANDLER_CONST {
				idhandler := new(IdHandler)
				handler_list = append(handler_list, idhandler)
			} else if flag == IP_HANDLER_CONST {
				iphandler := new(IpHandler)
				handler_list = append(handler_list, iphandler)
			} else if flag == SLOTID_HANDLER_CONST {
				slotidhandler := new(SlotIdHandler)
				handler_list = append(handler_list, slotidhandler)
			} else if flag == LAN_HANDLER_CONST {
				lanhandler := new(LanHandler)
				handler_list = append(handler_list, lanhandler)
			} else if flag == IDC_HANDLER_CONST {
				idcHandler := new(IDCHandler)
				handler_list = append(handler_list, idcHandler)
			} else if flag == UA_HANDLER_CONST {
				uaHandler := new(UaHandler)
				handler_list = append(handler_list, uaHandler)
			}
		}
	}
	stragery_one_temp := new(BaseHandler)
	whilelist := new(WhiteListHandler)
	muidOverHandler := new(MuidOverFrequcentHandler)
	stragery_one_temp.SetNext(whilelist)
	whilelist.SetNext(muidOverHandler)

	for ind, value := range handler_list {
		if ind == 0 {
			muidOverHandler.SetNext(value)
		} else {
			handler_list[ind-1].SetNext(value)
		}
	}
	stragery_one = stragery_one_temp

	util.Log.Info("reload strategy success ....")
}

func setNextHandler(next *BaseHandler, handler *BaseHandler) {
	handler.next = next
}
