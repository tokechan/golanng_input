package main

import "fmt"

func main() {
	//明示的な定義
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello GO"
	fmt.Println(s)

	var (
		i2 int    = 200
		s2 string = "Hello Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 150
	fmt.Println(i3)

	//暗黙的な定義
	// var 変数名 = 値
	i4 := 200
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// i4 := 500
	// fmt.Println(i4)
	i4 := "Hello"
	fmt.Println(i4)
}
