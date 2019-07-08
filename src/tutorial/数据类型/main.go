package main

//整型 ，  int8 ,int16, int32,int64， 对应无符号整型， uint8, uint16 ,uint32
//int16 对应c语言short
// int对应系统位， 32位一般是3个g多一点
// 对象长度内建 len（）

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	inttype()

	float()

	string1()

	char()

	rune1()

	changestring()

}

func inttype() {
	//十进制
	var a int = 10
	fmt.Printf("%d \n", a) //10
	fmt.Printf("%b \n", a) //1010 占位符%b表示二进制

	//八进制 以0开头
	var b int = 077
	fmt.Printf("%o \n", b) //77

	//十六, 0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) //ff
	fmt.Printf("%X \n", c) //FF

	//变量的内存地址
	fmt.Printf("%p \n", &a) //站位符%p表示十六进制的内存地址

}

func float() {

	//float32  , float64
	//float 常量
	max := math.MaxFloat64
	fmt.Println(max)

}

func string1() {
	//字符串转义
	fmt.Println("c:\\\"go\"") //  output - > c:\"go"

	//单行文本
	var s1 = "sksk"
	//	多行
	var mString = `
daasdasda
apple
pear
da
da"不需要转义"
`
	fmt.Println(s1, mString)
	fmt.Println(len(mString), len(s1))

	//拼接加字符串
	fmt.Println(s1 + mString)
	s3 := fmt.Sprintf("%s%s", s1, mString)
	fmt.Println(s3)

	//分割
	s4 := strings.Split(mString, "d")
	fmt.Println(s4)

	//判断是否包含
	s5 := strings.Contains(mString, "da")
	fmt.Println(s5) // op-> true

	//前缀后缀判断
	s6 := strings.HasPrefix(mString, "aw")
	s7 := strings.HasSuffix(mString, "da")
	fmt.Println(s6, s7)

	//字串位置
	fmt.Println(strings.Index(mString, "apple"))    //18
	fmt.Println(strings.LastIndex(mString, "pear")) //24

	//join
	a1 := []int{1, 2, 3, 4}
	fmt.Println(a1)

	//遍历字符串
	for i := 0; i < len(mString); i++ {
		fmt.Printf("%c", mString[i]) //出现乱码 ，因为拆分了8位中文， 需要使用rune类型
	}

	//rune1（）

}

func char() {
	c1 := "golang"
	c2 := 'c' //ASCII码站一个字节 （8bit）
	c3 := '中'
	fmt.Println(c1, c2) //golang , 99 是 c的编码
	fmt.Println(c3)     //3位 -》 20013
	c4 := "hello中国"
	fmt.Println(len(c4)) // hello -> 5  ,  中->3  国-> 3  一共是 ： 11
}

func rune1() {
	c4 := "hello中国"
	//遍历字符串
	fmt.Println("rune类型")
	for i := 0; i < len(c4); i++ {
		fmt.Printf("%c", c4[i]) // 输出乱码helloä¸­å½
	}
	fmt.Println()
	//rune
	for k, v := range c4 {
		fmt.Printf("%d  %c", k, v) //  h  e  l  l  o  中  国      ,0  h1  e2  l3  l4  o5  中8  国
	}

	//如果想要修改字符串， 需要修改成rune，或者byte类型， 在转换回去
	//见changestring()
}

//强制类型转换
func changestring() {
	s1 := "big"
	var byteArr = []byte(s1)
	fmt.Println(byteArr)
	fmt.Println("开始修改")
	byteArr[2] = 'a'
	fmt.Println(string(byteArr))
}
