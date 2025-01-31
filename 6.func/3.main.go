package main

import "fmt"

// クロージャーの実装
// クロージャーは、関数がその関数の外で宣言された変数を参照できるようにするもの
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("Goodbye"))
	fmt.Println(f(""))
	fmt.Println(f("good"))
	fmt.Println(f("morning"))
}
