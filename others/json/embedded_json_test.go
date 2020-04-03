package jsontest

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}`

func TestEmbeddedJson(t *testing.T) {
	// 内置json解析是利用反射完成的，性能不是很高
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e) // json对应的字段赋值给结构体e
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e) // {{Mike 30} {[Java Go C]}}

	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
		// {"basic_info":{"name":"Mike","age":30},"job_info":{"skills":["Java","Go","C"]}}
	} else {
		t.Error(err)
	}
}
