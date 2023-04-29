package main

import (
	"fmt"
	"sync"
)

/*
	var st struct{ A, B, C int }

	// @desc sync（ミューテックスによる同期処理）
	func UpdateAndPrint(n int) {
		st.A = n
		time.Sleep(time.Microsecond)
		st.B = n
		time.Sleep(time.Microsecond)
		st.C = n
		time.Sleep(time.Microsecond)
		fmt.Println(st)
	}
*/

/*
var st2 struct{ A, B, C int }

var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	mutex.Lock() // ロック - ロックされている間は1つのgoroutin歯科処理を行うことはできない。

	st2.A = n
	time.Sleep(time.Microsecond)
	st2.B = n
	time.Sleep(time.Microsecond)
	st2.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st2)

	mutex.Unlock() // アンロック
}

func main() {
	mutex = new(sync.Mutex)
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}
	for {

	}
}

*/

// ゴルーチンの終了を待ち受ける。
func main() {
	/* sync.WaitGroupを生成 */
	wg := new(sync.WaitGroup)
	/* 待ち受けるゴルーチンの数は3 */
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
		wg.Done() // 完了
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
		wg.Done() // 完了
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
		wg.Done() // 完了
	}()

	/* ゴルーチンの完了を待ち受ける */
	wg.Wait()
	// Doneが3つ完了するまで待つ
}
