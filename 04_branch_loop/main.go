package main

import "fmt"

func main() {
	x := 2
	switch {
	case x < 2:
		fmt.Println("less than 2")
	case x == 2:
		fmt.Println("equal 2")
		fallthrough
	case x > 3:
		fmt.Println("greater than 2")
	default:
		fmt.Println("default")
	}
	/*
		equal 2
		greater than 2
	*/

	s := "isornot"
	for index, c := range s {
		fmt.Printf("%d %d %c\n", index, c, c)
	}
	/*
		0 105 i
		1 115 s
		2 111 o
		3 114 r
		4 110 n
		5 111 o
		6 116 t
	*/

	s1 := "中文字符串"
	for index, c := range s1 {
		fmt.Printf("%d %d %c\n", index, c, c)
	}
	/*
		0 20013 中
		3 25991 文
		6 23383 字
		9 31526 符
		12 20018 串
	*/

	arr := []int{4, 5, 6}
	for index, i := range arr {
		fmt.Println(index, i)
	}
	/*
		0 4
		1 5
		2 6
	*/

	m := map[string]int{"age": 12, "num": 15}
	for k, v := range m {
		fmt.Println(k, v)
	}
	/*
		age 12
		num 15
	*/
}
