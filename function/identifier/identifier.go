package main

import "fmt"

const (
	morning   = "おはよう"
	afternoon = "こんにちは"
	evening   = "こんばんは"
)

func あいさつ(m string) {
	fmt.Println(m)
}

func main() {
	あいさつ(evening)
}
