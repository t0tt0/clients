package tests

import (
	"github.com/HyperService-Consortium/go-ves/ves/test/tester"
	"reflect"
)

var srv *tester.Tester

var intT = 1
var intType = reflect.TypeOf(intT)
var uintType = reflect.TypeOf(uint(1))

func RangeInt(l, r int) []int {
	var x = make([]int, r-l)
	for i := l; i < r; i++ {
		x[i-l] = i
	}
	return x
}
