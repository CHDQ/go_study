package main

import "fmt"

/**
http://docscn.studygolang.com/pkg/fmt/
 */
type point struct {
	x, y int
}

func main() {
	//通用格式
	var str = "steven"
	fmt.Printf("%T , %v \n", str, str)
	p := point{1, 2}
	fmt.Printf("%T , %+v \n", p, p)
	//bool格式
	fmt.Printf("%T , %t \n", true, true)
	//整型
	fmt.Printf("%T , %d \n", 123, 123)
	fmt.Printf("%T , %0d \n", 123, 123)
	fmt.Printf("%T , %05d \n", 123, 123)
	fmt.Printf("%T , %b \n", 123, 123) //二进制
	str1 := fmt.Sprintf("%b", 1234)    //会返回一个字符串
	fmt.Println(str1)
	fmt.Printf("%X\n", 123)
	fmt.Printf("%U\n", 'a')
	//浮点型
	fmt.Printf("%.2f\n", 123.1) //会四舍五入
	fmt.Printf("%.2e\n", 123.1) //
	//字符串
	arr := [3]byte{
		97, 98, 99,
	}
	arr2 := [3]rune{
		97, 98, 99,
	}
	fmt.Printf("%s\n", "测试")            //科学计数法
	fmt.Printf("%s\n", "测试")            //科学计数法
	fmt.Printf("%T , %s\n", arr, arr)   //科学计数法
	fmt.Printf("%T , %s\n", arr2, arr2) //科学计数法

}
