package main

import (
	"fmt"
)

var S string = ""

func init() {
	S = S + "A"
}

func init() {
	S = S + "B"
}

func init() {
	S = S + "C"
}

func main() {
	fmt.Println(S)
}
