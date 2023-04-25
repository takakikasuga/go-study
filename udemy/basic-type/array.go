package main

import "fmt"

// @desc 配列型
func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)
	fmt.Printf("%T\n", arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)
	fmt.Printf("%T\n", arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])

	arr2[2] = "C"
	fmt.Println(arr2[2])

	// var arr5 [4]int
	// arr5 = arr1 //  cannot use arr1 (variable of type [3]int) as [4]int value in assignment
	// fmt.Println(arr5)

	fmt.Println(len(arr1))
}
