package main

import "fmt"

func main() {
	var x interface{} = 2
	switch v := x.(type) {
	case bool:
		fmt.Println("bool:", v)
	case int:
		fmt.Println(v * v)
	case string:
		fmt.Println(v)
	default:
		fmt.Printf("others: %#v\n", v)
	}
}
