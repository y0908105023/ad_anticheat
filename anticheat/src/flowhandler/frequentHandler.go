package flowhandler

import (
	util "hashutil"
	"strings"
	rpc "thrift/gen-go/filtergen"
	"time"
	"tripod/devkit"
)

var Producer chan string = make(chan string, 20000)

func WhriteToChannel(request *rpc.AntiSpamRequest) {
	if strings.TrimSpace(request.ID) != "" && len(Producer) < 20000 {
		if request.IdType == rpc.IDType_origin_id {
			md5_id := devkit.HashFunc(request.ID)
			Producer <- md5_id
		} else {
			Producer <- request.ID
		}
	}
}

var Frequct_num = 100     //大于这个就当做超频
var MAX_MUID_NUM = 300000 //最多放到map 300000w个id

//TODO  把通过4000000判断改为时间的判断
func ReadToChannel() {
	idmap_minute := make(map[string]int)
	for {
		if len(idmap_minute) < MAX_MUID_NUM {
			id := <-Producer
			if strings.TrimSpace(id) != "" {
				cnt := idmap_minute[id]
				if cnt >= Frequct_num {
					OverFrequent.Set(id, cnt)
					delete(idmap_minute, id)
				} else if cnt > 0 && cnt < Frequct_num {
					cnt += 1
					idmap_minute[id] = cnt
				} else if cnt == 0 {
					idmap_minute[id] = 1
				}
			}
		} else {
			util.Log.Info("refresh map,OverFrequent size = %d", OverFrequent.Size())
			util.Log.Info("channel size = %d", len(Producer))
			idmap_minute = make(map[string]int)
			util.Log.Info("clear counter,idmap_minute = %d", len(idmap_minute))
		}
	}
}

func ClearOverFrequect() {
	hour := time.Now().Hour()
	sleep_time := 24 - hour
	time.Sleep(time.Duration(sleep_time) * time.Hour)
	for {
		OverFrequent.Clear()
		util.Log.Info("clear over frequent map,OverFrequent size = ", OverFrequent.Size())
		time.Sleep(24 * time.Hour)
	}
}

func init() {
	go ReadToChannel()
	go ClearOverFrequect()
}
