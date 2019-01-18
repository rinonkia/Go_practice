package main

import "fmt"

// スライスs0, s1, s2 を簡易宣言、それぞれを出力する関数
func app() {
	s0 := []int{1, 2, 3}
	s1 := []int{4, 5, 6}
	s2 := append(s0, s1...)

	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(s2)
	// 追加しても元の配列は変わらない
	return
}

// スライスの容量の増え方確認
func space() {
	// (A) まずは要素、容量共に0のスライスをmakeで作成
	sl := make([]int, 0, 0)
	fmt.Printf("(A) len=%d, cap=%d\n", len(sl), cap(sl))

	// (B) 要素を１つ追加
	sl = append(sl, 1)
	fmt.Printf("(B) len=%d, cap=%d\n", len(sl), cap(sl))

	// (C) 簡易スライスを追加、...って何やろ
	sl = append(sl, []int{2, 3, 4}...)
	fmt.Printf("(C) len=%d, cap=%d\n", len(sl), cap(sl))

	// (D) 1つ追加
	sl = append(sl, 5)
	fmt.Printf("(D) len=%d, cap=%d\n", len(sl), cap(sl))

	// (E) 4つ追加
	sl = append(sl, 6, 7, 8, 9)
	fmt.Printf("(E) len=%d, cap=%d\n", len(sl), cap(sl))

	return
}

// copy(a, b) aのコピー先として（容量そのまま）、b先頭(インデックス:0)から上書きする
// コピー先が変わっちゃう
func copyNum() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{10, 11}
	n := copy(slice1, slice2) // 10, 11, 3, 4, 5
	// p := copy(slice2, slice1) にした場合-> 1, 2
	fmt.Println(n)
	fmt.Println(slice1)
	fmt.Println(slice2)
}

// 完全スライス式容量を設定できる

func finalSlice() {
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ss := numbers[2:4]
	fmt.Println(ss)

	sss := numbers[2:4:4]
	fmt.Println(sss)
}

func main() {

	app()

	space()

	copyNum()

	finalSlice()
}
