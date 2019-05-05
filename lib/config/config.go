// 配置文件相关
package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// 保存配置文件的参数
var g_config map[string]interface{}
var g_filename string

// 加载配置文件
func Load(filename string) error {
	g_filename = filename
	f, err := os.Open(g_filename)
	if err != nil {
		return err
	}
	defer f.Close()
	fileBuffer, err := ioutil.ReadAll(f)
	if err != nil {
		return nil
	}
	return json.Unmarshal(fileBuffer, &g_config)
}

// 获取字符串型配置文件参数
func GetString(key string, args ...string) string {
	defaultValue := ""
	if len(args) > 0 {
		defaultValue = args[0]
	}
	value := GetParam(key)
	if value == nil {
		return defaultValue
	}
	return value.(string)
}

// 获取[]string型参数
func GetStrings(key string) []string {
	value := GetParam(key)
	if value == nil {
		return nil
	}
	values := value.([]interface{})
	var r []string
	for _, v := range values {
		r = append(r, v.(string))
	}
	return r
}

func SetString(key string, value string) {
	g_config[key] = value
}

// 获取整数型配置文件参数
func GetInt(key string, args ...int) int {
	defaultValue := 0
	if len(args) > 0 {
		defaultValue = args[0]
	}
	value := GetParam(key)
	if value == nil {
		return defaultValue
	}
	return int(value.(float64))
}

// 设置整数型参数
func SetInt(key string, value int) {
	g_config[key] = value
}

// 获取参数
func GetParam(key string) interface{} {
	if _, ok := g_config[key]; !ok {
		return nil
	}
	return g_config[key]
}

// 设置参数
func SetParam(key string, value interface{}) {
	g_config[key] = value
}

// 创建全局map保存参数
func init() {
	g_config = make(map[string]interface{})
}
