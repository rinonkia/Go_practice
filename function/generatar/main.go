// 1ずつiがインクリメントするやつ

package main

import "fmt"

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()
	otherInts := integers()

	fmt.Println(ints())
	fmt.Println(ints())

	fmt.Println(otherInts())

	fmt.Println(ints())

}
