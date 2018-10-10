package flowhandler

import (
	util "hashutil"
	"strings"
	rpc "thrift/gen-go/filtergen"
)

type LanHandler struct {
	BaseHandler
}

/*
 * 局域网ip过滤器
 */
func (lanHandler *LanHandler) ProcessRequest(request *rpc.AntiSpamRequest, response *rpc.AntiSpamResponse) *rpc.AntiSpamResponse {
	// start := time.Now().UnixNano()
	if strings.TrimSpace(request.IP) != "" && len(strings.Split(request.IP, ".")) == 4 {
		ip_int := util.Ip_int(request.IP)
		// (127[.]0[.]0[.]1) --------------------------------------- 127.0.0.1(2130706433 --> 2147483647) local machine
		// (localhost) --------------------------------------------- localhost local machine
		// (10[.]d{1,3}[.]d{1,3}[.]d{1,3}) ---------------------- Class A 10.0.0.0-10.255.255.255(167772160->184549375)
		// (172[.]((1[6-9])|(2d)|(3[01]))[.]d{1,3}[.]d{1,3}) ---- Class B 172.16.0.0 (2886729728)-172.31.255.255(2887778303)
		// (192[.]168[.]d{1,3}[.]d{1,3}) ------------------------- Class C 192.168.0.0-192.168.255.255(3232235520->3232301055)
		if (ip_int >= 2130706433 && ip_int <= 2147483647) ||
			(ip_int >= 167772160 && ip_int <= 184549375) ||
			(ip_int >= 2886729728 && ip_int <= 2887778303) ||
			(ip_int >= 3232235520 && ip_int <= 3232301055) {
			return lanHandler.PackNoneSlotResponse(len(request.SlotIds), "20", false)
		} else {

		}
	}

	// end := time.Now().UnixNano()
	// util.Log.Info("lan time =", (end - start))

	if lanHandler.GetNext() != nil {
		return lanHandler.GetNext().ProcessRequest(request, response)
	} else if response != nil {
		return response
	} else {
		return lanHandler.PackNoneSlotResponse(len(request.SlotIds), "0", true)
	}
}
