package array

import "testing"

func TestArrayInit(t *testing.T) {
	a := Array{}
	a.Init(10, 10, 100)
	if a.Length != 10 {
		t.Errorf("the length should be equal %d", 10)
	}
	for _, v := range a.Values {
		if v < 10 || v > 100 {
			t.Errorf("value should > 10 and < 100, but got %d", v)
		}
	}
}

func TestArrayClone(t *testing.T) {
	a := Array{}
	b := a.Clone()
	for i := range a.Values {
		if b.Values[i] == a.Values[i] {
			t.Errorf("a.Values[%d] should equal b.Values[%d]", i, i)
		}
	}
}
