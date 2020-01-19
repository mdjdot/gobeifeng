package main

import "time"

import "fmt"

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		time.Sleep(3 * time.Second)
		ch <- 10
	}(ch)

	select {
	case <-ch:
		fmt.Println("ch")
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}
}
