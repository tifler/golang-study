package main

import (
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		println("input routine");
        time.Sleep(2 * time.Second)
		println("will send");
		ch <- 123
		println("sent");
	}()

	println("output routine");
	var i int
	println("will recv");
	i = <-ch
	println("recved");
	println(i)
}
