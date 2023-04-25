package main

import (
	"fmt"
	"time"
)

// @desc time
func main() {

	/*
		// 時間の生成
		t := time.Now()
		fmt.Println(t)

		t2 := time.Date(2023, 4, 26, 8, 25, 30, 555, time.Local)
		fmt.Println(t2)
		fmt.Println(t2.Year())
		fmt.Println(t2.YearDay())
		fmt.Println(t2.Month())
		fmt.Println(t2.Weekday())
		fmt.Println(t2.Day())
		fmt.Println(t2.Hour())
		fmt.Println(t2.Minute())
		fmt.Println(t2.Second())
		fmt.Println(t2.Nanosecond())
		fmt.Println(t2.Zone())
	*/

	/*
	   // 時間の感覚を表現
	   fmt.Println(time.Hour)
	   fmt.Printf("%T\n", time.Hour)
	   fmt.Println(time.Minute)
	   fmt.Println(time.Second)
	   fmt.Println(time.Millisecond)
	   fmt.Println(time.Microsecond)
	   fmt.Println(time.Nanosecond)

	   // time.Duration型を文字列から生成する。
	   d, _ := time.ParseDuration("2h30m")
	   fmt.Println(d)

	   // 現在自国の2分30秒後を表すtime.Time型の取得
	   t3 := time.Now()
	   fmt.Println(t3)
	   t3 = t3.Add(2*time.Minute + 15*time.Second)
	   fmt.Println(t3)
	*/

	/*
		// 時間の差分を比較する。
		t5 := time.Date(2024, 7, 24, 0, 0, 0, 0, time.Local)
		t6 := time.Now()
		fmt.Println(t6)

		// t5 -  t6
		d2 := t5.Sub(t6)
		fmt.Println(d2)

		// 時刻を比較する。
		fmt.Println(t6.Before(t5))
		fmt.Println(t6.After(t5))
		fmt.Println(t5.Before(t6))
		fmt.Println(t5.After(t6))
	*/

	// 指定時間のスリープ
	time.Sleep(5 * time.Second)
	fmt.Println("5秒停止後表示")
}
