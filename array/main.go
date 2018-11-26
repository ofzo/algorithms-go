package array

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

//Array a array
type Array struct {
	Values []int
	Length int
}

//Clone clone a array
func (a *Array) Clone() (arr Array) {
	arr = Array{
		Values: make([]int, a.Length),
		Length: a.Length,
	}
	for index := 0; index < a.Length; index++ {
		arr.Values[index] = a.Values[index]
	}
	return
}

// Init create a random array
func (a *Array) Init(len int, min int, max int) {
	a.Values = make([]int, len)
	rand.Seed(time.Now().Unix())
	for index := 0; index < len; index++ {
		a.Values[index] = rand.Intn(max-min) + min
	}
	a.Length = len
}

//Call call a sort
func Call(sort func(*Array), a Array) {
	name := runtime.FuncForPC(reflect.ValueOf(sort).Pointer()).Name()
	startTime := float64(time.Now().UnixNano())
	sort(&a)
	endTime := float64(time.Now().UnixNano())
	for index := 1; index < a.Length; index++ {
		if a.Values[index-1] > a.Values[index] {
			fmt.Println(name, "sort filed")
			return
		}
	}
	fmt.Printf("%48s: %16f ms\n", name, (endTime-startTime)/1000/1000)
}

//NearBy create nearby array
func (a *Array) NearBy(len int, count int) {
	a.Values = make([]int, len)
	a.Length = len
	for index := 0; index < len; index++ {
		a.Values[index] = index
	}
	for index := 0; index < count; index++ {
		i, j := rand.Intn(len), rand.Intn(len)
		a.Values[i], a.Values[j] = a.Values[j], a.Values[i]
	}
}
