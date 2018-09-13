package sort

import "github.com/roadhub/algorithms/array"

func merge(a *array.Array, l, r int) {
	if r-l <= 10 {
		insert(a, l, r)
		return
	}
	mid := (r-l)/2 + l
	merge(a, l, mid)
	merge(a, mid+1, r)
	arrLen := r - l + 1
	arr := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		arr[i] = a.Values[l+i]
	}
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid {
			a.Values[k] = arr[j-l]
			j++
		} else if j > r {
			a.Values[k] = arr[i-l]
			i++
		} else if arr[i-l] > arr[j-l] {
			a.Values[k] = arr[j-l]
			j++
		} else if arr[i-l] <= arr[j-l] {
			a.Values[k] = arr[i-l]
			i++
		}
	}
}

//Merge merge sort
func Merge(a *array.Array) {
	merge(a, 0, len(a.Values)-1)
}
