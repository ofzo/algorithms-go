package sort

import (
	"github.com/roadhub/algorithms/array"
)

type heap struct {
	data   []int
	length int
}

/* heap struct-------------------------------------*
 *						98(0)                      *
 *			 98(1)             	      84(2)        *
 *	  93(3)          90(4)		 8(5)	    18(6)  *
 * 49(7)  50(8)    76(9)					       *
 *-------------------------------------------------*/

func (h *heap) init(cap int) {
	h.data = make([]int, cap)
	h.length = cap
}

func (h *heap) from(a *array.Array) {
	h.data = make([]int, a.Length)
	for i, v := range a.Values {
		h.data[i] = v
	}
	h.length = a.Length
	for t := (h.length - 2) / 2; t >= 0; t-- {
		// fmt.Printf("%d => %+v\n", t, h)
		h.shiftDown(t)
	}
}

func (h *heap) shiftUp(index int) {
	for index >= 1 {
		t := (index - 1) / 2 // 父元素index
		if h.data[t] < h.data[index] {
			h.data[index], h.data[t] = h.data[t], h.data[index]
		}
		index = t
	}
}

func (h *heap) shiftDown(index int) {
	for 2*index+1 < h.length {
		t := 2*index + 1
		if t+1 < h.length && h.data[t] < h.data[t+1] {
			t = t + 1
		}
		if h.data[index] < h.data[t] {
			h.data[index], h.data[t] = h.data[t], h.data[index]
		}
		index = t
	}
}
func (h *heap) pop() int {
	t := h.length - 1
	h.data[0], h.data[t] = h.data[t], h.data[0]
	h.length--
	h.shiftDown(0)
	return h.data[t]
}

// Heap heap sort
func Heap(a *array.Array) {
	h := heap{}
	h.from(a)
	for i := 0; i < a.Length; i++ {
		a.Values[a.Length-i-1] = h.pop()
	}
}
