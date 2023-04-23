package main

import (
	"fmt"
)

/*
type error interface {
	Error() string
}
*/

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました。", ErrCode: 1234}
}

// @desc カスタムエラー
func main() {
	err := RaiseError()
	fmt.Println(err.Error())
	fmt.Println(err.Error())

	e, ok := err.(*MyError)

	if ok {
		fmt.Println(e.ErrCode)
	}
}
