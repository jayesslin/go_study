package main

import "fmt"

//变量名在前， 类型在后
//变量在一个作用域是不能重复申请的， 但是全局可以被函数覆盖
//批量声明
var (
	name1 string = "da"
	c     bool
	d     int32 = 110444
	e     int64 = 1112241221412
	f           = "aad"
)
var ww int

func main() {
	var name string
	//声明变量必须要使用 上面就报错了，如果不用
	fmt.Print(name)         //输出为空
	fmt.Println(ww)         //out put -> 0
	fmt.Println(name1)      //out put -> da
	fmt.Println(c, d, e, f) //O.P ->  false 110444 1112241221412 aad
	c = true
	//d = 11122212414141214// - > constant 11122212414141214 overflows int32
	d = 1
	e = 0
	fmt.Println(c, d, e)

	//fmt.Printf() 格式化
	fmt.Printf("字符串： %s  整数 ： %d", f, d)

	//类型推导
	var xx = "asd"
	fmt.Println(xx)

	//短变量声明， 只能在函数内部使用
	short := 100
	fmt.Println(short)

	//匿名变量
	//调用foo
	//正常用法
	fo1, fo2 := foo()
	fmt.Println(fo1, fo2) //输出 -> jayess , 9000
	//匿名变量用法, "_"代表不用变量
	fo3, _ := foo()
	fmt.Println(fo3) //输出 -> jayess

	//常量 ： 常量_枚举,  通常用在运行程序时不会改变的值
	const KEY = "thisIsMyKey"
	const PASSWORD = 112223
	const (
		pi = 3.14
		e  = 2.7
	)
}

//匿名变量
//定义两个返回值函数
func foo() (string, int) {
	return "jayess", 9000
}
