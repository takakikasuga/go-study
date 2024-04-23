// package main

// import "fmt"

// func main() {
// 	a := []int{1, 2, 3, 4}
// 	b := make([]int, len(a)) // スライスtの骨組みを作成する
// 	copy(b, a)
// 	a[0] = 100
// 	fmt.Printf("Value: %v, Capacity: %v, Length: %v\n", a, cap(a), len(a)) // Value: [100 2 3 4], Capacity: 4, Length: 4
// 	fmt.Printf("Value: %v, Capacity: %v, Length: %v\n", b, cap(b), len(b)) // Value: [1 2 3 4], Capacity: 4, Length: 4
// 	fmt.Printf("a == b: %v\n", &a == &b)                                   // a == b: false
// }

// package main

// import "fmt"

// func main() {
// 	var a [1]int
// 	b := &a
// 	b[0] = 1
// 	fmt.Printf("a:%v, b:%v, a == b:%v\n", a, b, a == *b)
// 	// a:[1], b:&[1], a == b:true
// 	fmt.Printf("pointer a:%p, pointer b:%p, pointer a[0]:%p, pointer b[0]:%p\n", &a, b, &a[0], &b[0])
// 	// pointer a:0xc000110018, pointer b:0xc000110018, pointer a[0]:0xc000110018, pointer b[0]:0xc000110018
// }

package main

import "fmt"

func main() {
	var a [3][2]int
	b := &(a[0]) // 配列が参照しているポインタの共有
	b[0] = 1
	b[1] = 5
	fmt.Printf("a:%v, b:%v, a[0] == b:%v\n", a, *b, a[0] == *b)
	// a:[[1 5] [0 0] [0 0]], b:[1 5], a[0] == b:true
	fmt.Printf("pointer a[0]:%p, pointer b:%p, pointer a[0][0]:%p, pointer b[0]:%p,  pointer a[0][1]:%p, pointer b[1]:%p\n", &a[0], b, &a[0][0], &(*b)[0], &a[0][1], &(*b)[1])
	// pointer a[0]:0xc000016180, pointer b:0xc000016180, pointer a[0][0]:0xc000016180, pointer b[0]:0xc000016180,  pointer a[0][1]:0xc000016188, pointer b[1]:0xc000016188
}
