package sort

import (
	"github.com/roadhub/algorithms/array"
)

//Selection selection sort
func Selection(a *array.Array) {
	for index := 0; index < a.Length; index++ {
		minIndex := index
		for j := index + 1; j < a.Length; j++ {
			if a.Values[j] < a.Values[minIndex] {
				minIndex = j
			}
		}
		a.Values[minIndex], a.Values[index] = a.Values[index], a.Values[minIndex]
	}
}

//Bubble sort a array
func Bubble(a *array.Array) {
	for i := 0; i < a.Length-1; i++ {
		for j := 0; j < a.Length-i-1; j++ {
			if a.Values[j] > a.Values[j+1] {
				a.Values[j], a.Values[j+1] = a.Values[j+1], a.Values[j]
			}
		}
	}
}

//Insert insert sort
func Insert(a *array.Array) {
	insert(a, 0, a.Length-1)
}

func insert(a *array.Array, l, r int) {
	for index := l; index <= r; index++ {
		e := a.Values[index]
		var j int
		for j = index; j > 0 && e < a.Values[j-1]; j-- {
			a.Values[j] = a.Values[j-1]
		}
		a.Values[j] = e
	}
}
