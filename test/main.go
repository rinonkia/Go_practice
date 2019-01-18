package main

import "fmt"

// 配列を2乗する関数
func arraySquare(numbers [3]int) {
	for i, v := range numbers {
		numbers[i] = v * v
	}
	return
}

// スライスを2乗する関数
func sliceSquare(numbers []int) {
	for i, v := range numbers {
		numbers[i] = v * v
	}
	return
}

func main() {
	// 配列で宣言をし、関数で2乗、もとのarrayAを出力
	arrayA := [3]int{1, 2, 3}
	arraySquare(arrayA)
	fmt.Println(arrayA)

	// スライスを宣言し、関数で2乗、もとのsliceAを出力
	sliceA := []int{1, 2, 3}
	sliceSquare(sliceA)
	fmt.Println(sliceA)

}
