package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 1; i < 100; i++ {
		go func(num int) {
			time.Sleep(5 * time.Second)
			fmt.Println(num)
		}(i)
	}
	time.Sleep(3*time.Second)
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(10*time.Second)
}
