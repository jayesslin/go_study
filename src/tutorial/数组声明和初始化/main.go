package main

import "fmt"

func main() {
	//数组大小不能变
	var a [3]int
	fmt.Println(a) // 0 0 0
	a = [3]int{1, 2}
	fmt.Println(a) // 1,2,0

	b := [2][2]string{
		{"sd", "aa"},
		{"ds", "d"},
	}
	fmt.Println(b)

	//省略数组大小
	c := [...]int{1, 2, 3, 45, 5}
	fmt.Println(c)

	//
	d := [20]int{}
	d = [20]int{18: 1, 2}
	fmt.Println(d)
	fmt.Println(d[19])
}
