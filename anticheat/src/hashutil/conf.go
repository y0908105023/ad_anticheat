package hashutil

import (
	l4g "code.google.com/p/log4go"
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"runtime"
)

var Log l4g.Logger

var rootPath = flag.String("rootPath", "/opt/zyz/filter", "root path")
var GoMaxProcs int
var WisdomPort int
var DebugFlag bool

type bucketInfo struct {
	Bucket int32
	Conf   string
}

type redisCluster []string
type bucketConfs []bucketInfo

var ServiceConfig struct {
	ZkHosts     string `json:"ZK_HOSTS"`
	GroupName   string `json:"GROUP_NAME"`
	LogConfFile string `json:"LOG_FILE"`
}

func newLogger(log_path string) l4g.Logger {
	Logger := make(l4g.Logger)
	Logger.LoadConfiguration(log_path)
	return Logger
}

func init() {
	flag.IntVar(&GoMaxProcs, "cpu", runtime.NumCPU(), "Number of CPU")
	flag.IntVar(&WisdomPort, "port", 9009, "Wisdom listen port, range [9000, 9100]")
	flag.BoolVar(&DebugFlag, "debugFlag", false, "debug flag ,false:normal,true:debug")
	flag.Parse()
	if GoMaxProcs <= 0 {
		panic("CPU number must be positive")
	}
	if WisdomPort < 9000 || WisdomPort > 9100 {
		panic("Wisdom listen port must be in range [9000, 9100]")
	}

	if !filepath.IsAbs(*rootPath) {
		absPath, err := filepath.Abs(*rootPath)
		if err != nil {
			panic("Convert root path to abs path failed")
		}
		*rootPath = absPath
	}

	LoadConfFile()
	println(ServiceConfig.LogConfFile + ServiceConfig.ZkHosts)
	Log = newLogger(ServiceConfig.LogConfFile)
}
func LoadConfFile() {
	//    configFile, err := os.Open(*rootPath + "/conf/zyz.config")
	filePath := "/opt/zyz/filter"
	configFile, err := os.Open(filePath + "/conf/zyz.config")
	if err != nil {
		panic("Open config file failed" + filePath + "/conf/zyz.config")
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&ServiceConfig); err != nil {
		panic("decode json file failed: " + err.Error())
	}
}
func GetRootPath() string {
	return *rootPath
}

// GetAbsPath returns the absolute path
// if path starts with "/", then returns itself
// else return rootPath + path
func GetAbsPath(path string) string {
	if path != "" && path[0] == '/' {
		return path
	}
	return filepath.Join(*rootPath, path)
}

func ConfigParser(confPath string, i interface{}) error {
	configFile, err := os.Open(confPath)
	if err != nil {
		panic("Open config file failed:" + err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(i); err != nil {
		panic("decode json file failed:" + err.Error())
	}
	return nil
}
