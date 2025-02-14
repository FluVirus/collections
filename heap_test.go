package collections

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createNewMinIntHeap() *Heap[int] {
	cmp := func(a, b int) int {
		if a < b {
			return -1
		}

		if a > b {
			return 1
		}

		return 0
	}

	return NewHeap(cmp)
}

func createNewMaxIntHeap() *Heap[int] {
	cmp := func(a, b int) int {
		if a > b {
			return -1
		}

		if a < b {
			return 1
		}

		return 0
	}

	return NewHeap(cmp)
}

func TestMinIntHeap(t *testing.T) {
	type TestCase struct {
		Name   string
		Input  []int
		Expect []int
	}

	var largeSlice []int
	for i := 0; i < 10000; i++ {
		num := rand.Intn(10000)
		largeSlice = append(largeSlice, num)
	}

	sortedLargeSlice := make([]int, len(largeSlice))
	copy(sortedLargeSlice, largeSlice)
	slices.Sort(sortedLargeSlice)

	testCases := []TestCase{
		{
			Name:   "single element",
			Input:  []int{1},
			Expect: []int{1},
		},
		{
			Name:   "two elements",
			Input:  []int{2, 1},
			Expect: []int{1, 2},
		},
		{
			Name:   "three elements",
			Input:  []int{2, 1, 3},
			Expect: []int{1, 2, 3},
		},
		{
			Name:   "large slice",
			Input:  largeSlice,
			Expect: sortedLargeSlice,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			heap := createNewMinIntHeap()
			for _, num := range testCase.Input {
				heap.Push(num)
			}

			slice := make([]int, 0, len(testCase.Input))
			for heap.Len() > 0 {
				value := heap.Pop()
				slice = append(slice, value)
			}

			assert.Equal(t, testCase.Expect, slice)
		})
	}
}

func TestMaxIntHeap(t *testing.T) {
	type TestCase struct {
		Name   string
		Input  []int
		Expect []int
	}

	var largeSlice []int
	for i := 0; i < 10000; i++ {
		num := rand.Intn(10000)
		largeSlice = append(largeSlice, num)
	}

	sortedLargeSlice := make([]int, len(largeSlice))
	copy(sortedLargeSlice, largeSlice)
	slices.SortFunc(sortedLargeSlice, func(a, b int) int {
		if a > b {
			return -1
		}

		if a < b {
			return 1
		}

		return 0
	})

	testCases := []TestCase{
		{
			Name:   "single element",
			Input:  []int{1},
			Expect: []int{1},
		},
		{
			Name:   "two elements",
			Input:  []int{2, 1},
			Expect: []int{2, 1},
		},
		{
			Name:   "three elements",
			Input:  []int{2, 1, 3},
			Expect: []int{3, 2, 1},
		},
		{
			Name:   "large slice",
			Input:  largeSlice,
			Expect: sortedLargeSlice,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			heap := createNewMaxIntHeap()
			for _, num := range testCase.Input {
				heap.Push(num)
			}

			slice := make([]int, 0, len(testCase.Input))
			for heap.Len() > 0 {
				value := heap.Pop()
				slice = append(slice, value)
			}

			assert.Equal(t, testCase.Expect, slice)
		})
	}
}
