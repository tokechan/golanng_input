package main

import "fmt"

func main() {
	var a float64 = 2.4
	fmt.Println(a)

	fl := 3.2
	//暗黙的な型変換

	fmt.Println(a + fl)
	fmt.Printf("%T, %T\n", a, fl)
	//型を表示する

	var b float32 = 1.2
	fmt.Printf("%T\n", b)

	fmt.Printf("%T\n", float64(b))
	//型変換

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)
}
