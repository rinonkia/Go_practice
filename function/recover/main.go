package main

import "fmt"

func main() {
	defer func() {
		if x := recover(); x != nil {
			/* 変数xはpanicに渡されたinterface{} */
			fmt.Println(x)
		}
	}()
	panic("panic!")
	/* これは実行されない */
	fmt.Println("Hello, World!")
}
