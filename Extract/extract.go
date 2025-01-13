package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "abc1234def5678ghi9, \n231ads"

	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(str, -1)

	fmt.Println(nums)

	// //////////////
	// numz := make([]int, 0)
	// num := ""
	// for _, char := range str {
	// 	if char >= '0' && char <= '9' {
	// 		num += string(char)
	// 	} else if num != "" {
	// 		n, _ := strconv.Atoi(num)
	// 		numz = append(numz, n)
	// 		num = ""
	// 	}
	// }
	// fmt.Println(numz)

	// ////////////////
	// numb := make([]int, 0)
	// numr := 0
	// for _, char := range str {
	// 	if char >= '0' && char <= '9' {
	// 		numr = numr*10 + int(char-'0')
	// 	} else if numr != 0 {
	// 		numb = append(numb, numr)
	// 		numr = 0
	// 	}
	// }
	// fmt.Println(numb)
}
