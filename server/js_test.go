package main

import (
	"chat/common"
	"encoding/json"
	"fmt"
	"testing"
)

func TestJS(t *testing.T) {

	data := common.CommonMsg{01, "02", "03"}
	ret, _ := json.Marshal(data)
	data2 := common.Data{
		Ip:      "127.0.0.1:59741",
		User:    "te22",
		Type:    "CommonMsg",
		Content: string(ret),
	}
	ret2, _ := json.Marshal(data2)
	fmt.Println(string(ret2))
}
