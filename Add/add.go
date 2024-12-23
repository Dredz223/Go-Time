package main

import "fmt"

func main() {
	x := 1
	y := 2
	z := x + y
	v := "Word Dawg"

	fmt.Println(z)
	i := 0
	p := &i
	fmt.Println(p)

	for ; i < 10; i++ {
		fmt.Println("Current Loop is", i)
		//fmt.Println("Address of the value is", p)
		//fmt.Println("Address of the value is", *p)
		*p += 2
	}

	p = &x
	fmt.Println("Address of the value is", p)

	fmt.Printf("This is printed using a percent V: %v \n", x)
	fmt.Printf("This is printed using a percent V: %v \n", v)
}
