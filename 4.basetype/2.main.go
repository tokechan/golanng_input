package main

import "fmt"

func main() {
	var i int = 100

	var i2 int64 = 200

	fmt.Println(i + 50)
	// fmt.Println(i +i2)

	fmt.Printf("%T\n", i2)

	fmt.Printf("%T\n", int32(i2))
	//型変換

	fmt.Println(i + int(i2))
	//型変換は型を変えることができるので、エラーにならない
}
