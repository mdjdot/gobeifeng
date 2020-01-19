package main

import (
	"errors"
	"fmt"
)

func add(num int) int {
	num++
	return num
}

func add1(nump *int) {
	*nump++
}

func checkZero(num int) {
	if num == 0 {
		panic(errors.New("num is 0"))
	}
}
func main() {
	{
		a := 1
		b := add(a)
		fmt.Println(a, b) // 1 2

		add1(&a)
		fmt.Println(&a, a) // 0xc000016088 2

		f := func(num int) int {
			num += 2
			return num
		}

		c := f(a)
		fmt.Println(c) // 4
	}

	{
		defer fmt.Println("最后输出")
		defer fmt.Println("倒数第二输出")
		func() {
			fmt.Println("something")
			// 退出应用
			return
			// os.Exit(0) // 使用os.Exit() defer将不执行
		}()
		/*
			something
			倒数第二输出
			最后输出
		*/
	}

	{
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("捕获到panic")
				fmt.Println(err)
			}
		}()
		checkZero(0)
		/*
			捕获到panic
			num is 0
		*/
	}

}
