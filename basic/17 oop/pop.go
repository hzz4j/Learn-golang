package main

import "fmt"

//准备洗衣服
//输入参数：
//powder 洗衣机里放多少洗衣粉
//closes 洗衣机里放多少衣服
//clean 衣服是否是干净的
//返回值：
//洗衣机是否开启
//准备洗多少衣服
func prepare(powder, closes int, clean bool) (bool, int) {
	if powder <= 0 || closes <= 0 || clean {
		return false, 0
	}
	return true, closes
}

//开始洗衣服
//输入参数：
//washer_state 洗衣机是否开启
//closes 准备洗多少衣服
//返回值：
//衣服是否是干净的
//洗了多少衣服
//洗衣机是否开启
func wash(washer_state bool, closes int) (bool, int, bool) {
	if !washer_state {
		return false, 0, false
	} else {
		fmt.Println("注水")
		fmt.Println("滚动")
		fmt.Println("关机")
		return true, closes, false
	}
}

//检查最终状态
//输入参数：
//clean 衣服是否是干净的
//closes 洗了多少衣服
//washer_state 洗衣机是否开启
func check(clean bool, closes int, washer_state bool) {
	if clean && closes > 0 {
		fmt.Printf("洗干净了%d件衣服\n", closes)
		if washer_state {
			fmt.Println("你忘关洗衣机了")
		}
	} else {
		fmt.Println("洗衣失败")
	}
}

//整个洗衣服的过程
func WashProcedure(powder, closes int) {
	washer_state := false
	clean := false

	washer_state, closes = prepare(powder, closes, clean)
	clean, closes, washer_state = wash(washer_state, closes)
	check(clean, closes, washer_state)
}
