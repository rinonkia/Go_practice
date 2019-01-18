// 関数から戻り値として関数を返す
package main

import (
	"fmt"
)

// returnFunc() and return value typus decleared
func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func main() {
	f := returnFunc()
	f()
}
