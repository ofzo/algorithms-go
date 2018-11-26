package test

import (
	"testing"

	"github.com/roadhub/algorithms/sort"

	"github.com/roadhub/algorithms/array"
)

//TestSelection test Selection
func TestSelection(t *testing.T) {
	a := array.Array{}
	a.Init(100000, 0, 100000)
	array.Call(sort.Selection, a)
	for index := 0; index < a.Length-1; index++ {
		if a.Values[index] > a.Values[index+1] {
			t.Errorf("error: a.Values[%d] not greater than a.Values[%d]", index, index+1)
		}
	}
}

//TestBubble test bubble
func TestBubble(t *testing.T) {
	a := array.Array{}
	a.Init(100000, 0, 100000)
	array.Call(sort.Bubble, a)
	for index := 0; index < a.Length-1; index++ {
		if a.Values[index] > a.Values[index+1] {
			t.Errorf("error: a.Values[%d] not greater than a.Values[%d]", index, index+1)
		}
	}
}

//TestInsert test insert
func TestInsert(t *testing.T) {
	a := array.Array{}
	a.Init(100000, 0, 100000)
	array.Call(sort.Insert, a)
	for index := 0; index < a.Length-1; index++ {
		if a.Values[index] > a.Values[index+1] {
			t.Errorf("error: a.Values[%d] not greater than a.Values[%d]", index, index+1)
		}
	}
}

func TestQuick(t *testing.T) {
	a := array.Array{}
	a.Init(100000, 0, 100000)
	array.Call(sort.Quick, a)
	for index := 0; index < a.Length-1; index++ {
		if a.Values[index] > a.Values[index+1] {
			t.Errorf("error: a.Values[%d] not greater than a.Values[%d]", index, index+1)
		}
	}
}

//TestHeap test heap
func TestHeap(t *testing.T) {
	a := array.Array{}
	a.Init(100000, 0, 100000)
	array.Call(sort.Heap, a)
	for index := 0; index < a.Length-1; index++ {
		if a.Values[index] > a.Values[index+1] {
			t.Errorf("error: a.Values[%d] not greater than a.Values[%d]", index, index+1)
		}
	}
}

//TestHeap test heap
func TestMerge(t *testing.T) {
	a := array.Array{}
	a.Init(100000, 0, 100000)
	array.Call(sort.Merge, a)
	for index := 0; index < a.Length-1; index++ {
		if a.Values[index] > a.Values[index+1] {
			t.Errorf("error: a.Values[%d] not greater than a.Values[%d]", index, index+1)
		}
	}
}
