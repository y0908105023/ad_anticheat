package flowhandler

import (
	util "hashutil"
	"os"
	"time"
)

type _loadFunc func(file string) (map[string]string, error)
type _loadFunc2 func(file string) (int, error)

// ReloadFile trys to reload file every particular interval, for ever
// always call this as go routine
// arguments:
//  1. load function
//  2. file name
//  3. reload interval
//  4. address of the stored result
func ReloadFileInterVal(f _loadFunc, file string, t time.Duration, h *map[string]string, filename string) {
	util.Log.Info("load file %s begin", filename)
	mt := 0
	for {
		if _mt := GetUnixFileMtime(file); _mt > mt || _mt == 0 {
			mt = _mt // reset mtime no matter success or fail
			if temp_map, err := f(file); err == nil {
				if temp_map != nil && len(temp_map) != 0 {
					util.Log.Info("reload file %s ok : filesize : %d", filename, len(temp_map))
					*h = temp_map
					if filename == "afs_rule.conf" {
						ReBuildStrategy()
					} else if filename == "ip_blacklist.model" {
						GetIDCFromIp(temp_map)
					}
				} else {
					*h = map[string]string{}
					util.Log.Info("file %s maybe is empty", filename)
				}
			}
			util.Log.Info("load file %s successful", filename)
		}
		time.Sleep(t)
	}
}

func ReloadFrequect(f _loadFunc2, file string, t time.Duration, filename string) {
	util.Log.Info("load file %s begin", filename)
	mt := 0
	for {
		if _mt := GetUnixFileMtime(file); _mt > mt {
			mt = _mt // reset mtime no matter success or fail
			if frequent, err := f(file); err == nil {
				if frequent != 0 {
					Frequct_num = frequent
					util.Log.Info("reload file %s ok : new frequent = : %d", filename, Frequct_num)
				} else {
					util.Log.Info("file %s maybe is empty", filename)
				}
			}
			util.Log.Info("load file %s successful", filename)
		}
		time.Sleep(t)
	}
}

func GetUnixFileMtime(file string) int {
	if info, err := os.Stat(file); err == nil {
		return int(info.ModTime().Unix())
	}
	util.Log.Info("file %s not exists", file)
	return 0
}
