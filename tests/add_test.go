package tests

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
测试用例分为单元测试、性能测试、和示例测试
运行该文件内的所有单元测试 go test --bench=".*"
案例结果： 输出ns/op的结果越小越好
BenchmarkStringF-16         1370            874218 ns/op
BenchmarkAdd-16               51          23091759 ns/op
BenchmarkBuilder-16         4580            270619 ns/op

*/

// 表格单元驱动测试
func TestAdd(t *testing.T) {
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 2, 3},
		{100, 100, 200},
		{-1, -9, -10},
		{0, 0, 0},
	}
	for _, v := range dataset {
		result := add(v.a, v.b)
		if result != v.out {
			t.Errorf("export %d sctual %d", result, v.out)
		}
	}
}

const countName = 10000

// 性能测试 测试string的3种拼接方式
func BenchmarkStringF(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < countName; j++ {
			fmt.Sprintf("%s%d", str, j)
		}
	}
	b.StopTimer()
}

func BenchmarkAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < countName; j++ {
			str = str + strconv.Itoa(j)
		}
	}
	b.StopTimer()
}

func BenchmarkBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < countName; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	b.StopTimer()
}
