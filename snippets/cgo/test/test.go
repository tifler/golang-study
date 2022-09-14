package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>

/*
#include <stdio.h>
#include <errno.h>
void hello() {
	printf("Hello\n");
}
*/

import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("hello, from go")
	cs := C.CString("Charles")
	fmt.Println(cs)
	C.hello()
	C.hello(cs)
	C.free(unsafe.Pointer(cs))
}
