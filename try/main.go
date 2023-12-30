package main

import (
	"fmt"
	"github.com/jenazads/gotry"
)

var try = gotry.Try
var throw = gotry.Throw

func main() {
	var obj interface{}
	obj = "hi"
	try(func() {
		text := obj.(string)
		fmt.Println("Try ---> ", text)
	}).Catch(func(err gotry.Exception) {
		fmt.Println("Catch ---> exception caught #1:", err)
		//throw(err)
	}).Finally(func() {
		fmt.Println("Finally ---> This always print after all try block #1")
	})
}
