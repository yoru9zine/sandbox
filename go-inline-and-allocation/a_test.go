package a

import "testing"

func somefunc(interface{}) {
}

func BenchmarkAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		somefunc(nil)
	}
}

func BenchmarkAllocationSmallSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make([]int, 1)
		somefunc(data)
	}
}

func BenchmarkAllocationLargeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make([]int, 10000)
		somefunc(data)
	}
}

func getSmallSlice() []int {
	return make([]int, 1)
}

func BenchmarkAllocationSmallSliceByFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := getSmallSlice()
		somefunc(data)
	}
}

func BenchmarkAllocationSmallSliceByAnonFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := func() []int {
			return make([]int, 1)
		}()
		somefunc(data)
	}
}
