package main

import "fmt"

func main() {
	x := 1
	y := 2
	z := x + y
	fmt.Println(z)
	p := &z
	fmt.Println(p)

	for ; z < 10; z++ {
		fmt.Println("Current Loop is", z)
		fmt.Println("Address of the value is", p)
		fmt.Println("Address of the value is", *p)
		*p += 2
	}

	p = &x
	fmt.Println("Address of the value is", p)

}
