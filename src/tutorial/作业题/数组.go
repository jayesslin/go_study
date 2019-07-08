package main

import "fmt"

func main() {

	//求数组[1, 3, 5, 7, 8]所有元素的和
	hw1_arr()
	//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	hw2_arr()
}

func hw1_arr() {
	//	求数组[1, 3, 5, 7, 8]所有元素的和
	a := [...]int{1, 3, 5, 7, 8}
	var res int
	for _, v := range a {
		res = res + v
	}
	fmt.Println(res)
}
func hw2_arr() {
	const z = 0
	//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	a := [...]int{1, 3, 5, 7, 8}
	//切片
	res := [][]int{}
	for k, v1 := range a {
		for i := k + 1; i < len(a)-1; i++ {
			if v1+a[i] == 8 {
				tmp := []int{k, i}
				res = append(res, tmp)
			}
		}
	}
	fmt.Println(res)
}
