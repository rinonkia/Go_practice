package main

import "fmt"

// 戻り値としてfunc(next string)を返し、そのfunc(next string)の結果string型の戻り値 newLine を返す関数
func later() func(string) string {
	// 明示的なstring型定義のみ。 newLine = ""
	var store string
	// 引数に仮変数next(string型)を取り、戻り値(string型)を返す
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := later()

	// f("1")はlater("1")に継承され、func(next string)に代入される
	fmt.Println(f("1"))
	fmt.Println(f("2"))
	fmt.Println(f("sunshine！!"))
}
