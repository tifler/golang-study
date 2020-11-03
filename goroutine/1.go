package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	say("Sync")

	go say("Asyny1");
	go say("Asyny2");
	go say("Asyny3");

	time.Sleep(time.Second * 3)
}
