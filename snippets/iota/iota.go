package main

import "fmt"

const (
	InitTimer uint32 = (1 << iota)
	InitVideo
	InitAudio
	InitEvent
)

func main() {
	fmt.Println("InitTimer: ", InitTimer)
	fmt.Println("InitVideo: ", InitVideo)
	fmt.Println("InitAudio: ", InitAudio)
	fmt.Println("InitEvent: ", InitEvent)
}
