package main

import (
	"fmt"
	"time"
)

func rev(c <-chan int) {
	fmt.Println("ready rev")
	r := <-c
	fmt.Println("rev:", r)
}

func main() {
	c := make(chan int, 2)
	go rev(c)
	fmt.Println("ready send")
	for i := 0; i < 2; i++ {
		c <- i
	}

	fmt.Println("OK")
	time.Sleep(3 * 1e9)
}
