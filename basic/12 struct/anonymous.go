package main

import (
	"encoding/json"
	"fmt"
)

type Screen01 struct {
	Size       float64 //屏幕尺寸
	ResX, ResY int     //屏幕分辨率 水平 垂直
}

type Battery struct {
	Capacity string
}

func getJson() (result []byte) {
	// 匿名结构体，临时组装数据
	tempData := &struct {
		Screen01
		Battery
		HasTouchId bool //是否有指纹识别
	}{
		Screen01:   Screen01{Size: 12, ResX: 36, ResY: 36},
		Battery:    Battery{Capacity: "3.5"},
		HasTouchId: true,
	}

	result, _ = json.Marshal(tempData)
	return
}

func main() {
	// {"Size":12,"ResX":36,"ResY":36,"Capacity":"3.5","HasTouchId":true}
	fmt.Printf("%s\n", string(getJson()))
}
