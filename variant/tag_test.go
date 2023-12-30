package Variant

import (
	"fmt"
	"testing"
)

func TestNewSome(t *testing.T) {
	some := NewSome(3)
	fmt.Printf("%v: %T\n", some, some)
	switch v := some.(type) {
	case Some[int]:
		fmt.Printf("%v\n", v)
	case None:
		fmt.Printf("none: %v", v)
	}
}
