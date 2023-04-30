package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}
type List []Entry

func (l List) Len() int {
	return len(l)
}
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return (l[i].Name < l[j].Name)
	} else {
		return (l[i].Value < l[j].Value)
	}
}

// @desc sort
func main() {

	i := []int{5, 3, 2, 4, 5, 6, 4, 8, 9, 8, 7, 10}
	s := []string{"a", "z", "j"}

	fmt.Println(i, s)

	sort.Ints(i)
	sort.Strings(s)

	fmt.Println(i, s)

	el := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}
	el2 := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}

	fmt.Println(el)

	// Slice
	sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name })

	sort.Slice(el, func(i, j int) bool { return el[i].Value < el[j].Value })

	fmt.Println("START --------- Slice")
	fmt.Println(el)
	fmt.Println("END --------- Slice")

	// SliceStable
	sort.SliceStable(el2, func(i, j int) bool { return el2[i].Name < el2[j].Name })

	sort.SliceStable(el2, func(i, j int) bool { return el2[i].Value < el2[j].Value })

	fmt.Println("START --------- SliceStable")
	fmt.Println(el2)
	fmt.Println("END --------- SliceStable")

	m := map[string]int{"S": 1, "J": 4, "A": 3, "N": 3}
	lt := List{}
	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e)
	}

	sort.Sort(lt)
	fmt.Println(lt)

	sort.Sort(sort.Reverse(lt))
	fmt.Println(lt)

}
