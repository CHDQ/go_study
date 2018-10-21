package main

import "fmt"

func main() {
	var (
		a int
		b string
		c float32
		d []int
		e int32
		f [3]int
		g bool
		h func() string
		l = 200.100
	)
	m := 19
	fmt.Printf("%T , %q \n", a, a)
	fmt.Printf("%T , %q \n", b, b)
	fmt.Printf("%T , %v \n", c, c)
	fmt.Printf("%T , %v \n", d, d)
	fmt.Printf("%T , %v \n", e, e)
	fmt.Printf("%T , %v \n", f, f)
	fmt.Printf("%T , %v \n", g, g)
	fmt.Printf("%T , %v \n", h, h)
	fmt.Printf("%T , %v \n", l, l)
	fmt.Printf("%T , %v \n", m, m)
	x := 10
	y := 20
	fmt.Println(x, "\t", y)
	x, y = y, x
	fmt.Println(x, "\t", y)
}
