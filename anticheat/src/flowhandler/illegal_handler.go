package flowhandler

import (
	util "hashutil"
	rpc "thrift/gen-go/filtergen"
	"time"
)

const (
	IP_FILE_PATH        = "/opt/zyz/filter/data/model/ip_blacklist.model"
	MUID_FILE_PATH      = "/opt/zyz/filter/data/model/muid_blacklist.model"
	SLOTID_FILE_PATH    = "/opt/zyz/filter/data/model/slotid_blacklist.model"
	STRATEGY_PATH       = "/opt/zyz/filter/data/model/afs_rule.conf"
	UA_PATH             = "/opt/zyz/filter/data/model/ua_blacklist.model"
	OVERFREQUCENT_PATH  = "/opt/zyz/filter/data/model/frequent.model" //频次文件，每分钟出现多少次当做超时
	MUID_FREQUCENT_PATH = "/opt/zyz/filter/data/model/muid_blacklist_active.model"
)

type FlowFilterService struct{}

var ip_map = map[string]string{}
var muid_map = map[string]string{}
var slotid_map = map[string]string{}
var strategy_map = map[string]string{}
var Idc_map = map[int64]string{}
var ua_map = map[string]string{}
var muid_frequcent_map = map[string]string{}

var OverFrequent = util.NewBeeMap()

func (ff *FlowFilterService) GetIllegalReason(request *rpc.AntiSpamRequest) (response *rpc.AntiSpamResponse, err error) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			util.Log.Error(err) // 这里的err其实就是panic传入的内容
		}
	}()
	startTime := time.Now()
	WhriteToChannel(request)
	return util.MonitorProcess(stragery_one.ProcessRequest, request, startTime), nil
}

func NewFlowFilterService() *FlowFilterService {
	return &FlowFilterService{}
}

func init() {
	go ReloadFileInterVal(util.ReadFromLocalFile, IP_FILE_PATH, 2*time.Minute, &ip_map, "ip_blacklist.model")
	go ReloadFileInterVal(util.ReadFromLocalFile, MUID_FILE_PATH, 2*time.Minute, &muid_map, "muid_blacklist.model")
	go ReloadFileInterVal(util.ReadFromLocalFile, SLOTID_FILE_PATH, 2*time.Minute, &slotid_map, "slotid_blacklist.model")
	go ReloadFileInterVal(util.ReadStrategy, STRATEGY_PATH, 1*time.Minute, &strategy_map, "afs_rule.conf")
	go ReloadFileInterVal(util.ReadStrategy, UA_PATH, 1*time.Minute, &ua_map, "ua_blacklist.model")
	go ReloadFileInterVal(util.ReadStrategy, MUID_FREQUCENT_PATH, 2*time.Minute, &muid_frequcent_map, "muid_blacklist_active.model")
	//每分钟出现多少次当做超时，默认100
	go ReloadFrequect(util.ReadFrequentNum, OVERFREQUCENT_PATH, 5*time.Minute, "frequent.model")
	util.Log.Info("INIT SUCCESSFULLY.")
}
