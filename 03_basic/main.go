package main

import "fmt"

func main() {
	var x int
	x = 2
	y := "beifeng"
	fmt.Printf("%d %s\n", x, y) // 2 beifeng

	// 复数
	// 加法 (a+bi)+(c+di)=(a+c)+(b+d)i
	// 减法 (a+bi)-(c+di)=(a-c)+(b-d)i
	// 乘法 (a+bi)(c+di)=(ac-bd)+(bc+ad)i
	// 除法
	var z complex64 = 2 + 4i
	fmt.Println(z * z) // (-12+16i)

	var aa [10]int
	fmt.Println(aa) // [0 0 0 0 0 0 0 0 0 0]
	aa[0] = 1
	aa[1] = 2
	fmt.Println(aa) // [1 2 0 0 0 0 0 0 0 0]

	bb := [10]int{1, 2, 3, 4, 5}
	fmt.Println(bb)        // [1 2 3 4 5 0 0 0 0 0]
	fmt.Printf("%v\n", bb) // [1 2 3 4 5 0 0 0 0 0]

	ss := [10]string{"a", "b"}
	fmt.Println(ss)        // [a b        ]
	fmt.Printf("%v\n", ss) // [a b        ]

	var s1 []int
	fmt.Println(s1) // []
	s2 := make([]int, 10)
	fmt.Println(s2) // [0 0 0 0 0 0 0 0 0 0]

	s3 := new([]int)
	fmt.Println(s3) // &[]

	arr := [10]int{1, 2, 3, 4, 5}
	fmt.Println(arr) // [1 2 3 4 5 0 0 0 0 0]
	s4 := arr[3:7]
	fmt.Println(s4) // [4 5 0 0]
	s4[3] = 6
	fmt.Println(s4)               // [4 5 0 6]
	fmt.Println(len(s4), cap(s4)) // 4 7
	fmt.Println(arr)              // [1 2 3 4 5 0 6 0 0 0]

}
