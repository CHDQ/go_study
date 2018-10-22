package main

import "fmt"

func main() {
	chinese := 90
	english := 80.9
	avg := (chinese + int(english)) / 2
	avg2 := (float64(chinese) + english) / 2 //float 会根据计算机系统的位数决定位数
	fmt.Println(avg)
	fmt.Println(avg2) //bool不能转换成任何其他类型
	fmt.Println(string(chinese))
}
