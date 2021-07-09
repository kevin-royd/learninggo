package main

import (
	"math"
	"testing"
)

/*
表格驱动测试
*/

func TestTriangle(t *testing.T) {
	tests := []struct {
		a, b, c int32
	}{{1, 2, 3}, {0, 2, 2}, {0, 0, 0}, {-1, 1, 0}, {math.MaxInt32, 1, math.MinInt32}}
	for _, test := range tests {
		if actual := calcTriangle(test.a, test.b); actual != 1 {
			t.Errorf("calcTrinagle(%d,%d);"+"got %d; expected %d", test.a, test.b, actual, test.c)
		}
	}
}

func calcTriangle(a int32, b int32) (actual int32) {
	return a + b
}
