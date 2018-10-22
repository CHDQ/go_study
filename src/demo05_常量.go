package main

import "fmt"

func main() {
	const NAME string = "STEVEN" //常量只能是基础类型
	const ( //常量组如果没有初始值，后面的常量会和第一个常量值一致（第一个没有初始值会报错）
		A = 10
		B
		C
	)
	fmt.Println(A, B, C)
	const ( //iota在常量组里面遇见一个常量会自增一次
		L = iota
		M = iota
		N = iota
	)
	const ( //iota在常量组里面遇见一个常量会自增一次
		L1 = iota
		M1
		N1
	)
	const ( //iota在常量组里面遇见一个常量会自增一次
		L2 = NAME
		M2 = iota //1
		N2
	)
	const ( //不定义的时候，值为上一个常量的表达式表示的值
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	const PERSON = iota //单独使用的时候为0,只有在常量组中才会自增
	fmt.Println(L, M, N)
	fmt.Println(L1, M1, N1)
	fmt.Println(L2, M2, N2)
	fmt.Println(i, j, k, l)
}
