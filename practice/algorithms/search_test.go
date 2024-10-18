package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type test struct {
	name     string
	testCase func(t *testing.T)
}

func Test_BinarySearch(t *testing.T) {
	tests := []test{
		{
			name: "best case",
			testCase: func(t *testing.T) {
				arr := []int{5, 4, 3, 2, 1}
				target := 5
				expected := 0
				if binary_search(arr, target) != expected {
					t.Fail()
				}
				if linear_search(arr, target) != expected {
					t.Fail()
				}
			},
		},
		{
			name: "worst case: present",
			testCase: func(t *testing.T) {
				arr := []int{1, 2, 3, 4, 5}
				target := 5
				expected := 4
				if binary_search(arr, target) != expected {
					t.Fail()
				}
				if linear_search(arr, target) != expected {
					t.Fail()
				}
			},
		},
		{
			name: "worst case: absent",
			testCase: func(t *testing.T) {
				arr := []int{5, 4, 3, 2, 1}
				target := 0
				expected := -1
				if binary_search(arr, target) != expected {
					t.Fail()
				}
				if linear_search(arr, target) != expected {
					t.Fail()
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.testCase)
	}
}

func Benchmark_Search(b *testing.B) {
	fmt.Println("Benchmarking Linear Search:")
	testing.Benchmark(Benchmark_LinearSearch)
	fmt.Println("Benchmarking Binary Search:")
	testing.Benchmark(Benchmark_LinearSearch)
}

func Benchmark_BinarySearch(b *testing.B) {
	arr := generateRandomArray(1000000)
	target := rand.Intn(1000000)
	for i := 0; i < b.N; i++ {
		binary_search(arr, target)
	}
}

// Number of Gouroutines			Iterations	Average Time		Bytes/Operation    Avg Mem alloc/Operation
// Benchmark_BinarySearch-8   	    3740	    309515 ns/op	    2139 B/op	       0 allocs/op
// Benchmark_LinearSearch-8   	    8787	    253771 ns/op	     911 B/op	       0 allocs/op

func Benchmark_LinearSearch(b *testing.B) {
	arr := generateRandomArray(1000000)
	target := rand.Intn(1000000)
	for i := 0; i < b.N; i++ {
		linear_search(arr, target)
	}
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)

	rand.NewSource(time.Now().UnixMilli())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size)
	}
	return arr
}
