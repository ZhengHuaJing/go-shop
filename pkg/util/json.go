package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 通过json文件路径打开并转换成map
func JsonFileToMap(jsonPath string) map[string]map[string]map[string]interface{} {
	data := map[string]map[string]map[string]map[string]interface{}{}

	buff, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(buff, &data)

	return data["paths"]
}
