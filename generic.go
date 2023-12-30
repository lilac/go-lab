package main

import (
	"fmt"
	"sort"
)

func genericFun[T any](i, j T) {
	fmt.Printf("i: %v, j: %v\n", i, j)
}

func main() {
	genericFun(5, 6)
	genericFun[float32](10.5, 11.8)
	res := Sort[int](IntOrd{})([]int{1, 3})
	fmt.Println(res)
	ints := SortGen([]MyInt{5, 1, 4})
	fmt.Println(ints)
}

type Ord[T any] interface {
	compare(a, b T) int
}

type IntOrd struct{}

func (o IntOrd) compare(a, b int) int {
	return a - b
}

func Sort[T any](ord Ord[T]) func([]T) []T {
	return func(data []T) []T {
		sort.SliceStable(data, func(i, j int) bool {
			a := data[i]
			b := data[j]
			return ord.compare(a, b) < 0
		})
		return data
	}
}

type MyInt int

func (i MyInt) compare(a, b MyInt) int {
	return int(a) - int(b)
}

func SortGen[T Ord[T]](data []T) []T {
	head := data[0]
	sort.SliceStable(data, func(i, j int) bool {
		a := data[i]
		b := data[j]
		return head.compare(a, b) < 0
	})
	return data
}
