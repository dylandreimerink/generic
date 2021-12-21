package datastruct

import (
	"fmt"
	"strconv"
	"testing"
)

type TestStruct struct {
	A int64
	B float64
	C string
}

func BenchmarkSinglyLinkedListGrowHead(b *testing.B) {
	for i := int64(1); i <= 12; i++ {
		n := int64(128) << i
		b.Run(fmt.Sprintf("%s_%d", b.Name(), n), func(bb *testing.B) {
			benchmarkSLLGrowthHead(bb, n)
		})
	}
}

func benchmarkSLLGrowthHead(b *testing.B, n int64) {
	for i := 0; i < b.N; i++ {
		var ll SinglyLinkedList[TestStruct]
		for ii := int64(0); ii < n; ii++ {
			ll.InsertHead(TestStruct{
				A: ii,
				B: float64(ii),
				C: strconv.FormatInt(ii, 10),
			})
		}
	}
}

func BenchmarkSinglyLinkedListGrowTail(b *testing.B) {
	for i := int64(1); i <= 12; i++ {
		n := int64(128) << i
		b.Run(fmt.Sprintf("%s_%d", b.Name(), n), func(bb *testing.B) {
			benchmarkSLLGrowthTail(bb, n)
		})
	}
}

func benchmarkSLLGrowthTail(b *testing.B, n int64) {
	for i := 0; i < b.N; i++ {
		var ll SinglyLinkedList[TestStruct]
		for ii := int64(0); ii < n; ii++ {
			ll.InsertTail(TestStruct{
				A: ii,
				B: float64(ii),
				C: strconv.FormatInt(ii, 10),
			})
		}
	}
}

func BenchmarkSliceGrowthTail(b *testing.B) {
	for i := int64(1); i <= 12; i++ {
		n := int64(128) << i
		b.Run(fmt.Sprintf("%s_%d", b.Name(), n), func(bb *testing.B) {
			benchmarkSliceGrowthTail(bb, n)
		})
	}
}

func benchmarkSliceGrowthTail(b *testing.B, n int64) {
	for i := 0; i < b.N; i++ {
		var ll []TestStruct
		for ii := int64(0); ii < n; ii++ {
			ll = append(ll, TestStruct{
				A: ii,
				B: float64(ii),
				C: strconv.FormatInt(ii, 10),
			})
		}
	}
}
