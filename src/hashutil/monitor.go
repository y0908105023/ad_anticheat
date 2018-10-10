package hashutil

import (
	"math"
	"sync/atomic"
	rpc "thrift/gen-go/filtergen"
	"time"
)

var Count_Channel chan float64 = make(chan float64, 20000)

type monitourFunc func(request *rpc.AntiSpamRequest, response *rpc.AntiSpamResponse) *rpc.AntiSpamResponse

/*
 *  监控函数执行时间，并写入channel，此处是监控主执行时间
 */
func MonitorProcess(tag monitourFunc, request *rpc.AntiSpamRequest, start time.Time) *rpc.AntiSpamResponse {
	r := tag(request, nil)
	//执行时间---微秒单位
	process_time := (time.Now().Sub(start).Seconds() * 1000 * 1000)
	Count_Channel <- process_time
	return r
}

func WriteMonitor() {
	ticker := time.NewTicker(time.Minute * 2) // 定时2分钟
	var sum_process_time AtomicFloat64 = 0.0
	var sum_process_count AtomicFloat64 = 0.0
	go func() {
		for {
			select {
			case process_time := <-Count_Channel:
				sum_process_time.Add(process_time)
				sum_process_count.Add(1)
			}
		}
	}()
	for {
		select {
		case <-ticker.C:
			Log.Info("process_time = %f ,process_count = %f", sum_process_time.Value(), sum_process_count.Value())
			if sum_process_count.Value() == 0.0 {
				Log.Info("fanzuobimonitor\t%f", 0.0)
			} else {
				avg_time := sum_process_time.Value() / sum_process_count.Value()
				Log.Info("fanzuobimonitor\t%f", avg_time)
				sum_process_time = 0.0
				sum_process_count = 0.0
			}
		}
	}
}

type AtomicFloat64 uint64

func (f *AtomicFloat64) Value() float64 {
	return math.Float64frombits(atomic.LoadUint64((*uint64)(f)))
}

func (f *AtomicFloat64) Add(n float64) {
	for {
		a := atomic.LoadUint64((*uint64)(f))
		b := math.Float64bits(math.Float64frombits(a) + n)
		if atomic.CompareAndSwapUint64((*uint64)(f), a, b) {
			return
		}
	}
}

func init() {
	go WriteMonitor()
}
