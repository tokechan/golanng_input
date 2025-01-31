package main

import (
	"fmt"
	"strconv"
)

// if文
// エラーハンドリング
func main() {
	// 文字列変数 s を宣言し、"A"を代入
	var s string = "A"

	// strconv.Atoi関数を使って文字列を整数に変換
	// 戻り値は2つ：変換後の整数(i)とエラー情報(err)
	i, err := strconv.Atoi(s)

	// エラーチェック
	// err が nil でない場合（エラーが発生した場合）の処理
	if err != nil {
		fmt.Println(err)
	}

	// 変換結果の整数を出力
	fmt.Println(i)
}
