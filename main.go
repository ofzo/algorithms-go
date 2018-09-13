package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/roadhub/algorithms/array"
	"github.com/roadhub/algorithms/sort"
)

func main() {
	a := array.Array{}
	l := 10
	if len(os.Args) > 1 {
		l, _ = strconv.Atoi(os.Args[1])
	}
	a.Init(l, 0, 10000000)
	a2 := a.Clone()
	a3 := a.Clone()
	a4 := a.Clone()
	a5 := a.Clone()

	// fmt.Printf("%+v\n", a)
	fmt.Println("a.Init")
	array.Call(sort.Insert, a)
	array.Call(sort.Bubble, a2)
	array.Call(sort.Selection, a3)
	array.Call(sort.Merge, a4)
	array.Call(sort.Quick, a5)

	a.NearBy(l, 0)
	a2 = a.Clone()
	a3 = a.Clone()
	a4 = a.Clone()
	a5 = a.Clone()
	fmt.Println("\na.Nearby")
	array.Call(sort.Insert, a)
	array.Call(sort.Bubble, a2)
	array.Call(sort.Selection, a3)
	array.Call(sort.Merge, a4)
	array.Call(sort.Quick, a5)
	// fmt.Printf("%+v\n", a)

}
