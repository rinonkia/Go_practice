package main

import "fmt"

func main() {
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	// i：文字列配列のインデックス
	// s：インデックスに対応した値
	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}
}
