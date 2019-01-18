// 関数の引数として関数を渡す
package main

import "fmt"

func callFunction(f func()) {
	f()
}

func main() {
	callFunction(func() {
		fmt.Println("I'm a function")
	})
}
