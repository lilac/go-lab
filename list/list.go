package list

import (
	GList "github.com/zyedidia/generic/list"
)

type Number interface {
	int | float64 | float32
}

func NewList[T Number]() *GList.List[T] {
	list := GList.New[T]()
	list.PushFront(1)
	list.PushBack(1.2)
	return list
}
