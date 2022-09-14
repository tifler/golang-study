package main

import (
    "fmt"
)

type Phone struct {
    Model string
}

func (p Phone) Call(num string) {
    fmt.Println("Ring Ring....", num)
}

type Camera struct {
    Model string
}

func (p Camera) TakePicture() {
    fmt.Println("Click ....")
}

type Recorder struct {
	Model string
}

func (p Recorder) Record(dur int) {
	fmt.Println("Recording", dur, "...")
}

type SmartPhone struct {
    Phone
    Camera
	Recorder
}

func main() {
    myPhone := SmartPhone{}
    myPhone.TakePicture()
    myPhone.Call("101-1111-2222")
    myPhone.Record(10)
    fmt.Println("=================")
    yourPhone := SmartPhone{}
    yourPhone.Call("201-2222-1111")

}

