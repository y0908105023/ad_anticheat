package hashutil

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func GetFileLength(path string) int {

	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(f), "\n")
	return len(lines)
}

func ReadFrequentNum(path string) (int, error) {
	if !Exist(path) {
		Log.Info("file %s is not exists", path)
		return 0, nil
	}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
		Log.Error("load strategy file error:", err)
		return 0, err
	}
	defer f.Close()
	num := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := scanner.Text()
		if strings.TrimSpace(value) == "" {
			continue
		}
		if num_temp, error := strconv.Atoi(value); error != nil {
			return 0, error
		} else {
			num = num_temp
		}
	}
	return num, nil
}

func ReadStrategy(path string) (map[string]string, error) {
	if !Exist(path) {
		Log.Info("file %s is not exists", path)
		return nil, nil
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
		Log.Error("load strategy file error:", err)
		return nil, err
	}
	defer f.Close()
	strategy_map_result := map[string]string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()
		if strings.TrimSpace(val) == "" {
			continue
		}
		key_value := strings.Split(val, "	")
		if len(key_value) >= 2 {
			strategy_map_result[key_value[0]] = key_value[1]
		} else {
			Log.Error("val %s is error,val's length less than 2", val)
		}
	}
	return strategy_map_result, nil

}

func ReadFromLocalFile(path string) (map[string]string, error) {
	if !Exist(path) {
		Log.Info("file %s is not exists", path)
		return nil, nil
	}
	// file_length := GetFileLength(path)
	Map_reslut := map[string]string{}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
		Log.Error("load file error:", err)
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()
		if strings.TrimSpace(val) == "" {
			continue
		}
		key_value := strings.Split(val, "	")
		if len(key_value) >= 2 {
			Map_reslut[key_value[0]] = key_value[1]
		} else {
			Log.Error("val %s is error,val's length less than equals 2", val)
		}
	}
	return Map_reslut, nil
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
