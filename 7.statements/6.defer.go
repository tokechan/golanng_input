package main

import (
	"fmt"
	"os"
)

// defer

func RunDefer() {
	defer fmt.Println("world")
	defer fmt.Println("hello")
	defer fmt.Println("end")
}

func main() {
	// deferは関数の終了時に実行される シンプル版
	/*
		defer fmt.Println("world")
		fmt.Println("hello")
	*/

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write([]byte("Hello"))
}
