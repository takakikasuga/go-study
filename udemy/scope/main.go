package main

import (
	. "fmt"
	f "fmt"
	"scope/foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	b = s
	{
		b := "BBBB"
		f.Println(b)
	}
	f.Println(b)
	return b
}

// @desc スコープ
func main() {
	f.Println(foo.Max)
	// f.Println(foo.min) // undefined: foo.min ⇨ 一文字目が大文字：public / 一文字目が小文字：private
	f.Println(foo.ReturnMin())
	Println(foo.ReturnMin())
	Println(appName())
	// Println(AppName, Version) // undefined: AppName / undefined: Version
	// f.Println(Do("AAA")) // b redeclared in this block
	f.Println(Do("AAA"))
}
