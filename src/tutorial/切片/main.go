package main

import (
	"fmt"
	"sort"
)

//切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
//
//切片是一个*引用类型*，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。

//切片的长度和容量
//切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
//
//基于数组定义切片
//由于切片的底层就是一个数组，所以我们可以基于数组定义切片。

func main() {
	//定义切片
	slc := []int{1, 2, 3, 4, 5}
	part := slc[2:]
	fmt.Println(part)     //3,4,5   包含2号位
	fmt.Println(slc[1:4]) //[2 3 4]

	//切片在切片
	fmt.Println(part[1:]) // 3,4,5 - > [4 5]

	//切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
	// len 求长度， cap 求切片容量
	s1 := slc
	s2 := slc[1:4]
	s3 := s2[:2]
	s4 := s3[2:]
	fmt.Println(s1, len(s1), cap(s1)) //[1 2 3 4 5] 5 5
	fmt.Println(s2, len(s2), cap(s2)) //[2 3 4] 3 4
	fmt.Println(s3, len(s3), cap(s3)) //[2] 1 4
	fmt.Println(s4, len(s4), cap(s4)) //[] 0 2

	//make() 函数构造切片
	//我们上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置的make()函数，格式如下：
	//make([]T, size, cap)
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10

	//切片排序--》 快排
	sorttest := []int{123, 432, 24, 2, 6, 332, 6, 43}
	sort.Ints(sorttest)
	fmt.Println(sorttest) //[2 6 6 24 43 123 332 432]

}
