package main

import "fmt"

func main() {
	n := 0.0000000000000000000000001
	switch {
	case 0 < n && n < 3:
		fmt.Println("0 < n < 3")
	case 3 <= n && n < 6:
		fmt.Println("3 <= n < 6")
	default:
		fmt.Println("others")
	}
}
