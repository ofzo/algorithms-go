package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/roadhub/algorithms/array"
	"github.com/roadhub/algorithms/sort"
)

func nearby(a array.Array, l int) {
	a.NearBy(l, 10)
	fmt.Println("\na.Nearby")
	call(a)
}
func _init(a array.Array, l int) {
	a.Init(l, 0, 10000000)
	fmt.Println("a.Init")
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
	a := array.Array{}
	l := 10
	if len(os.Args) > 1 {
		l, _ = strconv.Atoi(os.Args[1])
	}
	// a.Init(l, 0, 100)
	_init(a, l)
	nearby(a, l)
}
