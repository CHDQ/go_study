package main

import "fmt"

func main() {
	a := 100
	var (
		b byte = 100
		c rune = 200
		e byte = 'a' //uint8
		f rune = 'a' //int32
	)
	fmt.Printf("%T , %v \n", a, a)
	fmt.Printf("%T , %v \n", b, b)
	fmt.Printf("%T , %v \n", c, c)
	fmt.Printf("%T , %v \n", e, e)
	fmt.Printf("%T , %v \n", f, f)
	temp := `//原样输出字符串
	x := 10
	y := 20
	fmt.Println(x, "\t", y)
	x, y = y, x
	fmt.Println(x, "\t", y)
`
	fmt.Println(temp)
}
