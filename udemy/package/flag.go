package main

import (
	"flag"
	"fmt"
)

// @desc flag
func main() {
	// コマンドラインのオプション処理
	// コマンドラインを処理するサンプル go run flag.go -n 20 -m message -x

	// オプションの値を格納する変数の定義
	var (
		max int
		msg string
		x   bool
	)

	// IntVar 整数オプション
	flag.IntVar(&max, "n", 32, "処理数の最大値")
	// StringVar 整数オプション
	flag.StringVar(&msg, "m", "", "処理メッセージ")
	// StringVar 整数オプション
	flag.BoolVar(&x, "x", false, "拡張オプション")

	// コマンドラインをパース（指定されているポインタに値が入る。）
	flag.Parse()

	fmt.Println("処理の最大値 === ", max)
	fmt.Println("処理メッセージ === ", msg)
	fmt.Println("拡張オプション === ", x)

}
