package main

import "fmt"

func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			// panicによるinterpace{}型の値に応じて処理を分岐
			switch v := x.(type) {
			case int:
				fmt.Printf("panic: int=%v\n", v)
			case string:
				fmt.Printf("panic: string=%v\n", v)
			default:
				fmt.Println("panic: unknown")
			}
		}
	}()
	panic(src)
	return
}

func main() {
	testRecover(128)
	testRecover("hogehoge")
	testRecover([...]int{1, 2, 3})
}
