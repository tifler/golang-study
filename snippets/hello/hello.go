package main

import (
	"fmt"
	"example/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
    //fmt.Println("Hello, world.")
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	sum := func(a, b int) int { return a * b } (3, 4)
	fmt.Println("sum: ", sum)
}
