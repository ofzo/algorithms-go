package sort

import (
	"math/rand"

	"github.com/roadhub/algorithms/array"
)

func quick(a *array.Array, l, r int) {
	if r-l <= 10 {
		insert(a, l, r)
		return
	}
	e := rand.Intn(r-l+1) + l
	a.Values[e], a.Values[l] = a.Values[l], a.Values[e]
	b := a.Values[l]
	// fmt.Println(b)
	// [l...i] [i+1...j-1(e)] [j...r]
	e, i, j := l+1, l, r
	for e <= j {
		if a.Values[e] > b {
			a.Values[j], a.Values[e] = a.Values[e], a.Values[j]
			j--
		} else if a.Values[e] < b {
			i++
			a.Values[i], a.Values[e] = a.Values[e], a.Values[i]
			e++
		} else {
			e++
		}
	}
	a.Values[i], a.Values[l] = a.Values[l], a.Values[i]
	i--
	quick(a, l, i)
	quick(a, j, r)
}

//Quick quick sort
func Quick(a *array.Array) {
	quick(a, 0, a.Length-1)
}
