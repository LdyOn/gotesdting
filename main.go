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
	// ok值判断key=name的值是否存在，若存在ok为true且赋值给value
	value, ok := commissionMap[name]
	result, validity = 0, false
	var _ok bool
	if ok && value != nil {
		// 这里断言保存的值是数值类型
		result, _ok = value.(float64)
		if _ok {
			// 断言成功，有效性设为true
			validity = true
		} else {
			// 断言失败打印错误、记录日志
			fmt.Printf("获取费率%v失败\n", name)
		}

	}
	// 返回结果
	return result, validity
}

func main() {
	// 待处理的json字符串
	commision := `{
		"y1":null,
		"y2":2000,
		"y4":0,
		"y5":"hello",
		"add":2000
	}`
	str := "{\"a\":\"0\", \"b\":\"\", \"c\":\"2000\"}"
	type Str struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
	}
	var s Str
	json.Unmarshal([]byte(str), &s)
	fmt.Printf("s:%#v\n", s)
	fmt.Printf("s.a:%#v\n", s.B)
	type Commission struct {
		Y1  int64
		Y2  int64
		Y3  int64
		Y4  int64
		Y5  int64
		Add int64
	}
	// 定义一个map保存json数据
	var c map[string]interface{}
	// 调用库函数解析数据
	err := json.Unmarshal([]byte(commision), &c)
	if err != nil {
		fmt.Printf("json解析失败")
	}

	y1, y1Validity := getCommission(Y1, c)
	y2, y2Validity := getCommission(Y2, c)
	y3, y3Validity := getCommission(Y3, c)
	y4, y4Validity := getCommission(Y4, c)
	y5, y5Validity := getCommission(Y5, c)
	var cb Commission
	json.Unmarshal([]byte(commision), &cb)
	fmt.Printf("输入的json字符串：\n%v\n", commision)
	fmt.Printf("y1费率的值为：%v，有效性为：%v\n", y1, y1Validity)
	fmt.Printf("y2费率的值为：%v，有效性为：%v\n", y2, y2Validity)
	fmt.Printf("y3费率的值为：%v，有效性为：%v\n", y3, y3Validity)
	fmt.Printf("y4费率的值为：%v，有效性为：%v\n", y4, y4Validity)
	fmt.Printf("y5费率的值为：%v，有效性为：%v\n", y5, y5Validity)
	fmt.Printf("cb费率的值为：%v\n", cb)
}
