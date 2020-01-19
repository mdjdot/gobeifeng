package main

import "time"

import "fmt"

func read(c chan int) {
	fmt.Println(<-c)
}

func write(c chan int) {
	c <- 10
}

func main() {
	{
		// var c chan int
		c := make(chan int) // 不分配内存是不行的
		// a := new(chan int)
		// c := *a
		go read(c)
		time.Sleep(3 * time.Second)
		go write(c)

		time.Sleep(5 * time.Second)
	}

	{
		ch1 := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			ch1 <- i // channel 先进先出
		}
		for i := 1; i <= 5; i++ {
			fmt.Println(<-ch1)
		}
	}

}
