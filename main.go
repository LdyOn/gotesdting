package main

import (
	"encoding/json"
	"fmt"
)

const (
	Y1  = "y1"
	Y2  = "y2"
	Y3  = "y3"
	Y4  = "y4"
	Y5  = "y5"
	ADD = "add"
)

// 定义一个函数从map获取费率
func getCommission(name string, commissionMap map[string]interface{}) (result float64, validity bool) {

	value, ok := commissionMap[name]
	result, validity = 0, false
	var _ok bool
	if ok && value != nil {
		result, _ok = value.(float64)
		if !_ok {
			fmt.Printf("获取费率%v失败", name)
		}
		validity = true
	}

	return result, validity
}

func main() {
	// 待处理的json字符串
	commision := `{
		"y1":null,
		"y2":2000,
		"y4":2000,
		"y5":2000,
		"add":2000
	}`
	// 定义一个map保存json数据
	var c map[string]interface{}
	// 调用库函数解析数据
	err := json.Unmarshal([]byte(commision), &c)
	if err != nil {
		fmt.Printf("json解析失败")
	}

	y1, y1Validity := getCommission(Y1, c)
	y3, y3Validity := getCommission(Y3, c)
	y2, y2Validity := getCommission(Y2, c)
	fmt.Printf("输入的json字符串：%v\n", commision)
	fmt.Printf("y1费率的值为：%v，有效性为：%v\n", y1, y1Validity)
	fmt.Printf("y2费率的值为：%v，有效性为：%v\n", y2, y2Validity)
	fmt.Printf("y3费率的值为：%v，有效性为：%v\n", y3, y3Validity)
}
