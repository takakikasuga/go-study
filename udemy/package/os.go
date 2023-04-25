package main

import (
	"fmt"
	"log"
	"os"
)

// @desc os
func main() {
	/*
		os.Exit(1)
		fmt.Println("Start")
	*/

	/*
		defer func() {
			fmt.Println("defer")
		}()
		os.Exit(1)
	*/

	/*
		// log.Fatalln
		_, err := os.Open("A.txt")
		if err != nil {
			log.Fatalln(err) // 引数の値を出力しつつ、「os.Exit(1)」で終了する。
		}
	*/

	/*
		// Args
		fmt.Println(os.Args[0])
		fmt.Println(os.Args[1])
		fmt.Println(os.Args[2])
		fmt.Println(os.Args[3])

		fmt.Printf("length=%d\n", len(os.Args))
		for i, v := range os.Args {
			fmt.Println(i, v)
		}
	*/

	/*
		// ファイル操作（Open）
		f, err := os.Open("test.txt")
		if err != nil {
			log.Fatalln(err) // 引数の値を出力しつつ、「os.Exit(1)」で終了する。
		}

		defer f.Close()
	*/

	/*
	   // ファイル操作（Create）
	   f, _ := os.Create("foo.txt") // 既存で同じ名前がある場合、上書きされてしまう。
	   f.Write([]byte("Hello\n"))
	   f.WriteAt([]byte("Golang"), 6) // ６文字目に移動
	   f.Seek(0, os.SEEK_END)         // 末尾に移動
	   f.WriteString("Yeah")
	*/

	/*
	   // ファイル操作（Read）
	   f, err := os.Open("foo.txt")

	   	if err != nil {
	   		log.Fatalln(err) // 引数の値を出力しつつ、「os.Exit(1)」で終了する。
	   	}

	   defer f.Close()

	   bs := make([]byte, 128)
	   n, err := f.Read(bs)
	   fmt.Println(n)
	   fmt.Println(string(bs))

	   bs2 := make([]byte, 128)
	   nn, err := f.ReadAt(bs2, 10)
	   fmt.Println(nn)
	   fmt.Println(string(bs2))
	*/

	// ファイル操作（OpenFile - O_RDONLY_読み込み専用 / O_WRONLY_書き込み可能 / O_RDWR_読み書き専用 / O_APPEND_ファイル末尾に追記 / O_CREATE_ファイルがなければ作成 / O_TRUNC_可能であればファイルの内容をオープン時にからにする。）
	f, err := os.OpenFile("foo.txt", os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln(err) // 引数の値を出力しつつ、「os.Exit(1)」で終了する。
	}

	defer f.Close()
	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
}
