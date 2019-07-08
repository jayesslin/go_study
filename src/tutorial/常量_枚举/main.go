package main

import "fmt"

//常量定义的时候需要赋值
//const 声明如果不屑，默认就和上行一样
// iota
// 1. const 每新增一行变量声明iota就递增1

const pi = 3.14
const (
	n1 = 1000
	n2 //默认跟上面一样也是一千

)

// iota 枚举
const (
	aa = iota //0
	bb
	cc
	dd
	_
	ee
)

//系统容量限制
const (
	_ = iota
	//位运算
	KB = 1 << (10 * iota) // 1 左移 10 位  ， 2的十次方
	MB = 1 << (10 * iota) // 在左移10位， 2的20次方
	GB = 1 << (10 * iota) // 2的30 次方
	// ...
)

//运算2
const (
	a, b = iota + 1, iota + 2 // iota 为 0
	c, d                      // 默认与上面一行一样 ， iota +1, iota + 2 ，但是iota =1
	e, f                      // iota = 2  ， iota +1, iota + 2
)

func main() {

	fmt.Println(pi)
	//pi  = 3.1415 // 编译报错， 常量不许修改
	fmt.Println(n1, n2) // output -> 1000,1000
	//局部修改
	pi := "hello"
	fmt.Println(pi)

	//枚举
	fmt.Println(aa, bb, cc, dd) //output - > 0,1,2,3
	//匿名"_"
	fmt.Println(cc, dd, ee) //output -> 2,3,5

	//系统变量
	fmt.Println(KB, MB, GB) //output - > 1024 1048576 1073741824

	//运算2
	fmt.Println(a, b, c, d, e, f) //output -> 1 2 2 3 3 4

}
