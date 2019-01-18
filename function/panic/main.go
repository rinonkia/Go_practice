package main

import "fmt"

func main() {

	defer fmt.Println("first defer")

	panic("runtime error!")
	fmt.Println("Hello Word!")
	defer fmt.Println("second defer")
}
