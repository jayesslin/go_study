package main

import "fmt"

func main() {
	hw1("google")

	//编写代码打印9*9乘法表。
	hw2()
}

func hw1(test string) {
	// "hello" --> "olleh"
	s1 := test
	lenS1 := len(s1)
	byteArr := []byte(s1)
	//fmt.Println(lenS1,lenS1>>1)
	for k, _ := range byteArr {
		if k <= (lenS1 >> 1) {
			tmp := byteArr[k]
			byteArr[k] = byteArr[lenS1-1]
			//fmt.Println("k"+string(tmp)+"-1:"+string(byteArr[lenS1-1]))
			byteArr[lenS1-1] = tmp
			lenS1--
		} else {
			//fmt.Println(":end")
			break
		}
	}

	fmt.Println(string(byteArr))
}

func hw2() {
	//编写代码打印9*9乘法表。
	const mul = 9
	for i := 0; i <= mul; i++ {
		for j := 0; j <= mul; j++ {
			fmt.Printf("%d ", i*j)
		}
		fmt.Println()
	}
}
