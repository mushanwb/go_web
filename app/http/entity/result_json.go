package entity

import "encoding/json"

// 返回的 json 数据格式
type jsonResult struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ReturnJson(message string, data interface{}) []byte {
	jsonData, _ := json.Marshal(jsonResult{message, data})
	return jsonData
}
