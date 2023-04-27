package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// @desc ioutil
func main() {
	// 入力全体を読み込む

	f, _ := os.Open("foo.txt")
	bs, _ := ioutil.ReadAll(f)
	fmt.Println(bs)
	fmt.Println(string(bs))

	// ファイルの書き込み
	if err := ioutil.WriteFile("bar.txt", bs, 0666); err != nil {
		log.Fatalln(err)
	}
}
