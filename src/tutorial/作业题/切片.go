package main

import (
	"fmt"
	"sort"
)

//用内置sort包排序

func main() {
	a := []int{111, 2, 3, 5, 6}
	sort.Ints(a)
	fmt.Println(a)
}
