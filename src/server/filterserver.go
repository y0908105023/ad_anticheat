package main

import (
	"flowhandler"
	"fmt"
	util "hashutil"
	rpc "thrift/gen-go/filtergen"
	"tripod/define"
	"tripod/devkit"
	"tripod/zkutil"
)

func main() {

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
			util.Log.Error(err)
		}
	}()
	handler := flowhandler.NewFlowFilterService()
	processor := rpc.NewIllegalServiceProcessor(handler)
	rpc.NewIllegalServiceProcessor(handler)
	zkInstance := zkutil.GetZkInstance(util.ServiceConfig.ZkHosts)
	serverGroupName :=
		func() string {
			if util.DebugFlag {
				return define.GroupServingTest
			}
			return define.GroupFilter
		}()

	statPath := "/opt/zyz/" + define.FilterServerName + ".stop"
	server := zkutil.NewZServer(zkInstance, processor, serverGroupName,
		devkit.GetServiceName(define.FilterServerName), statPath)

	util.Log.Info("CREATE FILTER SERVICE SUCCESSFULLY.")
	server.Start()

}
