package flowhandler

import (
	util "hashutil"
	"strings"
	rpc "thrift/gen-go/filtergen"
)

type IDCHandler struct {
	BaseHandler
}

/*
 * IDC机房过滤
 */
func (idcHandler *IDCHandler) ProcessRequest(request *rpc.AntiSpamRequest, response *rpc.AntiSpamResponse) *rpc.AntiSpamResponse {
	// start := time.Now().UnixNano()
	if strings.TrimSpace(request.IP) != "" && len(strings.Split(request.IP, ".")) == 4 {
		ip := util.Ip_int(request.IP)
		idc_value := Idc_map[ip]

		if idc_value != "" {
			return idcHandler.PackNoneSlotResponse(len(request.SlotIds), idc_value, false)
		}
	}
	// end := time.Now().UnixNano()
	// util.Log.Info("idc time =", (end - start))
	if idcHandler.GetNext() != nil {
		return idcHandler.GetNext().ProcessRequest(request, response)
	} else if response != nil {
		return response
	} else {
		return idcHandler.PackNoneSlotResponse(len(request.SlotIds), "0", true)
	}
}

func GetIDCFromIp(all_ip map[string]string) {
	idc_map_temp := map[int64]string{}
	for key, value := range all_ip {
		if strings.Contains(key, "->") {
			keys := strings.Split(key, "->")
			key_start := util.Ip_int(keys[0])
			key_end := util.Ip_int(keys[1])

			for ; key_start <= key_end; key_start++ {
				idc_map_temp[key_start] = value
			}
		}
	}
	Idc_map = idc_map_temp
}
