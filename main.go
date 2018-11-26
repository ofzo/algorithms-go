package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/logrusorgru/aurora"
	"github.com/roadhub/algorithms/array"
	"github.com/roadhub/algorithms/sort"
)

var (
	splitLine = aurora.Red("-------------------------------------------------------------------------------------")
)

func nearby(l int) {
	a := array.Array{}
	a.NearBy(l, 10)
	fmt.Println(aurora.Green("a.Nearby, a array include some(~10) unorder int data:"))
	call(a)
}
func _init(l int) {
	a := array.Array{}
	a.Init(l, 0, 10000000)
	fmt.Println(aurora.Green("a.Init, a array include full random int data:"))
	call(a)
}
func equal(l int) {
	a := array.Array{}
	a.Init(l, 0, 10)
	fmt.Println(aurora.Green("a.Init, a array include large equal int data:"))
	call(a)
}
func call(a array.Array) {
	a2 := a.Clone()
	a3 := a.Clone()
	a4 := a.Clone()
	a5 := a.Clone()
	a6 := a.Clone()
	array.Call(sort.Insert, a)
	array.Call(sort.Bubble, a2)
	array.Call(sort.Selection, a3)
	array.Call(sort.Merge, a4)
	array.Call(sort.Quick, a5)
	array.Call(sort.Heap, a6)
}
func main() {

	// a := array.Array{}
	// a.Init(100000, -1000, 10000000000)
	// array.Call(sort.Quick, a)

	l := 100
	if len(os.Args) > 1 {
		l, _ = strconv.Atoi(os.Args[1])
	}
	fmt.Printf("\na.Length = %d\n", aurora.Red(l))
	fmt.Println(splitLine)
	// a.Init(l, 0, 100)
	_init(l)
	fmt.Println(splitLine)
	nearby(l)
	fmt.Println(splitLine)
	equal(l)
}
