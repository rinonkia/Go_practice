package main

import "fmt"

// use Slice and Range
func slicefor() {
	s := []string{"Apple", "Banana", "Cherry"}

	for i, v := range s {
		fmt.Printf("[%d] => %s\n", i, v)
	}
}

// 可変長引数(s ...int)とすると、引数の数が可変長となる
func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

// 参照型としてのスライス

func pow(numbers [3]int) {
	for i, v := range numbers {
		numbers[i] = v * v
	}
	return
}

func main() {

	slicefor()

	number := sum(1, 2, 3)
	fmt.Println(number)

	a := [3]int{1, 2, 3}
	pow(a)
	fmt.Println(a)
}
