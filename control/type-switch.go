package main

import "fmt"

func anything(a interface{}) {
	switch v := a.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println("I don't know")
	}
}

// @desc 型スイッチ
func main() {
	anything("aaa")
	anything(121313)

	var x interface{} = 3
	i, isInt := x.(int) // 型アサーション
	fmt.Println(i, isInt)
	f, isFloat64 := x.(float64) // 型アサーション
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't know")
	}

	switch x.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println("I don't know")
	}
}
