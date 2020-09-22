package main

import "fmt"

func main() {
	s1 := "hello world!"
	fmt.Printf("%T, %s\n", s1, s1)
	s2 := `test`
	fmt.Printf("%T, %s\n", s2, s2)
	v1 := 'h'
	v2 := "a"
	fmt.Printf("%T, %d, %c, %q\n", v1, v1, v1, v1)
	fmt.Printf("%T, %s\n", v2, v2)
	// 数据类型转换
	// go语言是强类型语言，不同类型之间转换
	// 基本数据运算符 + -*/%><

	n1 := 10.0
	n2 := 4.0
	// n3 := 10 /4
	fmt.Printf("%f / %f = %f", n1, n2, n1/n2)
}
