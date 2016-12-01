package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"runtime"
	rpc "thrift/gen-go/filtergen"
	"time"
)

var quit chan int = make(chan int, 10)

var transportFactory thrift.TTransportFactory
var protocolFactory thrift.TProtocolFactory

func main() {
	runtime.GOMAXPROCS(16) // 最多使用2个核
	startTime := currentTimeMillis()

	transportFactory = thrift.NewTBufferedTransportFactory(1024)
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	for i := 0; i < 1; i++ {
		go call(i)
	}

	t1 := currentTimeMillis()
	fmt.Println("t1. time->", (t1 - startTime))

	for i := 0; i < 1; i++ {
		<-quit
	}

	endTime := currentTimeMillis()
	fmt.Println("Program exit. time->", endTime, startTime, (endTime - startTime))
}

func call(i int) {

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
	}()

	time.Sleep(1 * time.Millisecond)

	transport, err := thrift.NewTSocketTimeout(net.JoinHostPort("10.200.95.55", "10000"), 30*time.Second)

	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	useTransport := transportFactory.GetTransport(transport)
	client := rpc.NewIllegalServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 10.200.95.55:10000", " ", err)
		os.Exit(1)
	}
	defer transport.Close()
	request := &rpc.AntiSpamRequest{"fafb1727044ab8855e9b8d9a55b6b14b", 2, "101.90.253.84", []string{"com.cleanmaster.security_cn"}, "GDTADNetClient-%5BDalvik%2F1.6.0+%28Linux%3B+U%3B+Android+4.4.4%3B+C8817D+Build%2FHuaweiC8817D%29%5D", "gdt", "CN310000", 12, 3}
	// client.GetIllegalReason(request)
	responses, _ := client.GetIllegalReason(request)
	legal := responses.Legals
	reasons := responses.Reasons
	fmt.Println(reasons)
	for index, reason := range reasons {
		fmt.Println(reason[index])
		fmt.Println(legal[index])
	}

	quit <- i
}

// transfer time to num
func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
