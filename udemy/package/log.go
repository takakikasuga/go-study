package main

import (
	"log"
	"os"
)

// @desc log
func main() {
	// ログの出力先を変更
	// log.SetOutput(os.Stdout)

	/*
		log.Print("Log\n")
		log.Println("Log2")
		log.Printf("Log%d\n", 3)
	*/

	/*
	   log.Fatal("Log\n")
	   log.Fatalln("Log2")
	   log.Fatalf("Log%d\n", 3)
	*/

	/*
	   log.Panic("Log\n")
	   log.Panicln("Log2")
	   log.Panicf("Log%d\n", 3)
	*/

	/*
	   // 任意のファイルを作成し、出力先を指定。
	   f, err := os.Create("test.log")

	   	if err != nil {
	   		return
	   	}

	   log.SetOutput(f)
	   log.Println("ファイルに書き込み")
	*/

	/*
		// ログのフィーマットを指定する
		// 標準のログフォーマット
		log.SetFlags(log.LstdFlags)
		log.Println("A")

		// マイクロ秒追加
		log.SetFlags(log.Ldate | log.LstdFlags | log.Lmicroseconds)
		log.Println("B")

		// 時刻とファイルの行番号
		log.SetFlags(log.Ldate | log.Lshortfile)
		log.Println("C")

		log.SetFlags(log.LstdFlags)
		// ログのプレフィックス指定
		log.SetPrefix("[LOG]")
		log.Println("E")
	*/

	// ロガーの生成
	logger := log.New(os.Stdout, "[Logger]", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Message")
	log.Println("Message")

	_, err := os.Open("foobarbaz")
	if err != nil {
		// ログの出力
		// log.Fatalln("Exit", err)
		logger.Fatalln("Exit", err)
	}
}
