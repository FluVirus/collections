package collections

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func toSlice[T any](list *LinkedList[T]) []T {
	slice := make([]T, 0, list.Len())

	for p := list.Head; p != nil; p = p.Next {
		slice = append(slice, p.Value)
	}

	return slice
}

func TestPushFront(t *testing.T) {
	type TestCase[T any] struct {
		Name  string
		Input []T
	}

	var manyElementsInput []int
	for i := 0; i < 932347; i++ {
		manyElementsInput = append(manyElementsInput, (38*i)^((3*i)+1))
	}

	testCases := []TestCase[int]{
		{
			Name:  "single element",
			Input: []int{5},
		},
		{
			Name:  "two elements",
			Input: []int{5, 6},
		},
		{
			Name:  "three elements",
			Input: []int{5, 6, 7},
		},
		{
			Name:  "many elements",
			Input: manyElementsInput,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			list := NewLinkedList[int]()

			for _, element := range testCase.Input {
				list.PushFront(element)
			}

			listContents := make([]int, 0, len(testCase.Input))
			for p := list.Tail; p != nil; p = p.Prev {
				listContents = append(listContents, p.Value)
			}

			assert.Equal(t, testCase.Input, listContents)
		})
	}
}

func TestPushBack(t *testing.T) {
	type TestCase[T any] struct {
		Name  string
		Input []T
	}

	var manyElementsInput []int
	for i := 0; i < 932347; i++ {
		manyElementsInput = append(manyElementsInput, (38*i)^((3*i)+1))
	}

	testCases := []TestCase[int]{
		{
			Name:  "single element",
			Input: []int{5},
		},
		{
			Name:  "two elements",
			Input: []int{5, 6},
		},
		{
			Name:  "three elements",
			Input: []int{5, 6, 7},
		},
		{
			Name:  "many elements",
			Input: manyElementsInput,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			list := NewLinkedList[int]()

			for _, element := range testCase.Input {
				list.PushBack(element)
			}

			listContents := make([]int, 0, len(testCase.Input))
			for p := list.Head; p != nil; p = p.Next {
				listContents = append(listContents, p.Value)
			}

			assert.Equal(t, testCase.Input, listContents)
		})
	}
}

func TestPopFront(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		require.Panics(t, func() {
			list := NewLinkedList[int]()

			list.PopFront()
		})
	})

	type TestCase[T comparable] struct {
		Name     string
		List     []T
		Expected T
	}

	var listForManyElements []int
	for i := 0; i < 932347; i++ {
		listForManyElements = append(listForManyElements, (38*i)^(3*i))
	}

	testCases := []TestCase[int]{
		{
			Name:     "single element",
			List:     []int{20},
			Expected: 20,
		},
		{
			Name:     "two elements",
			List:     []int{25, 33},
			Expected: 25,
		},
		{
			Name:     "three elements",
			List:     []int{26, 33, 61},
			Expected: 26,
		},
		{
			Name:     "many elements",
			List:     listForManyElements,
			Expected: listForManyElements[0],
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			list := NewLinkedList[int]()
			for _, element := range testCase.List {
				list.PushBack(element)
			}

			actual := list.PopFront()

			require.Equal(t, testCase.Expected, actual)
		})
	}
}

func TestPopBack(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		require.Panics(t, func() {
			list := NewLinkedList[int]()

			list.PopBack()
		})
	})

	type TestCase[T comparable] struct {
		Name     string
		List     []T
		Expected T
	}

	var listForManyElements []int
	for i := 0; i < 932347; i++ {
		listForManyElements = append(listForManyElements, (38*i)^(3*i))
	}

	testCases := []TestCase[int]{
		{
			Name:     "single element",
			List:     []int{20},
			Expected: 20,
		},
		{
			Name:     "two elements",
			List:     []int{25, 33},
			Expected: 33,
		},
		{
			Name:     "three elements",
			List:     []int{26, 33, 61},
			Expected: 61,
		},
		{
			Name:     "many elements",
			List:     listForManyElements,
			Expected: listForManyElements[len(listForManyElements)-1],
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			list := NewLinkedList[int]()
			for _, element := range testCase.List {
				list.PushBack(element)
			}

			actual := list.PopBack()

			require.Equal(t, testCase.Expected, actual)
		})
	}
}

func TestPushAndPop(t *testing.T) {
	t.Run("push back in empty list after pop", func(t *testing.T) {
		list := NewLinkedList[int]()
		firstPushBack := 10
		secondPushBack := 20

		list.PushBack(firstPushBack)

		pop := list.PopBack()
		assert.Equal(t, firstPushBack, pop)

		list.PushBack(secondPushBack)

		pop = list.PopBack()
		assert.Equal(t, secondPushBack, pop)
	})

	t.Run("push back in non-empty list after pop", func(t *testing.T) {
		list := NewLinkedList[int]()
		firstPushBack := []int{10, 11, 12}
		secondPushBack := 20

		for _, element := range firstPushBack {
			list.PushBack(element)
		}

		pop := list.PopBack()
		assert.Equal(t, firstPushBack[len(firstPushBack)-1], pop)

		list.PushBack(secondPushBack)
		pop = list.PopBack()

		assert.Equal(t, secondPushBack, pop)
	})
}

func TestIteration(t *testing.T) {
	type TestCase[T comparable] struct {
		Name string
		List []T
	}

	var listForManyElements []int
	for i := 0; i < 932347; i++ {
		listForManyElements = append(listForManyElements, (38*i)^(3*i))
	}

	testCases := []TestCase[int]{
		{
			Name: "empty list",
			List: make([]int, 0),
		},
		{
			Name: "single element",
			List: []int{5},
		},
		{
			Name: "two elements",
			List: []int{5, 6},
		},
		{
			Name: "three elements",
			List: []int{5, 6, 7},
		},
		{
			Name: "many elements",
			List: listForManyElements,
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("forward %s", testCase.Name), func(t *testing.T) {
			list := NewLinkedList[int]()
			for _, element := range testCase.List {
				list.PushBack(element)
			}

			slice := make([]int, 0)
			for p := list.Head; p != nil; p = p.Next {
				slice = append(slice, p.Value)
			}

			assert.Equal(t, testCase.List, slice)
		})
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("reverse %s", testCase.Name), func(t *testing.T) {
			list := NewLinkedList[int]()
			for _, element := range testCase.List {
				list.PushBack(element)
			}

			slice := make([]int, 0)
			for p := list.Tail; p != nil; p = p.Prev {
				slice = append(slice, p.Value)
			}

			expected := make([]int, len(testCase.List))
			copy(expected, testCase.List)
			slices.Reverse(slice)

			assert.Equal(t, expected, slice)
		})
	}
}

func TestLen(t *testing.T) {
	type TestCase[T any] struct {
		Name        string
		List        []T
		ExpectedLen int
	}

	var listForManyElements []int
	listForManyElementsLen := 932347
	for i := 0; i < listForManyElementsLen; i++ {
		listForManyElements = append(listForManyElements, (38*i)^(3*i))
	}

	testCases := []TestCase[int]{
		{
			Name:        "empty list",
			List:        make([]int, 0),
			ExpectedLen: 0,
		},
		{
			Name:        "single element",
			List:        []int{5},
			ExpectedLen: 1,
		},
		{
			Name:        "two elements",
			List:        []int{5, 6},
			ExpectedLen: 2,
		},
		{
			Name:        "three elements",
			List:        []int{5, 6, 7},
			ExpectedLen: 3,
		},
		{
			Name:        "many elements",
			List:        listForManyElements,
			ExpectedLen: listForManyElementsLen,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			list := NewLinkedList[int]()
			for _, element := range testCase.List {
				list.PushBack(element)
			}

			actualLen := list.Len()

			assert.Equal(t, testCase.ExpectedLen, actualLen)
		})
	}
}

func TestGet(t *testing.T) {
	type TestCase[T comparable] struct {
		Name     string
		List     []T
		Indexes  []int
		Expected []T
	}

	type PanicTestCase[T comparable] struct {
		Name    string
		List    []T
		Indexes []int
	}

	testCases := []TestCase[int]{
		{
			Name:     "single element",
			List:     []int{5},
			Indexes:  []int{0},
			Expected: []int{5},
		},
		{
			Name:     "two elements",
			List:     []int{5, 6},
			Indexes:  []int{0, 1},
			Expected: []int{5, 6},
		},
		{
			Name:     "three elements",
			List:     []int{5, 6, 7},
			Indexes:  []int{0, 2},
			Expected: []int{5, 7},
		},
	}

	panicTestCases := []PanicTestCase[int]{
		{
			Name:    "panic empty list",
			List:    []int{},
			Indexes: []int{0, 1},
		},
		{
			Name:    "panic non-empty list out of bounds",
			List:    []int{5, 6, 7},
			Indexes: []int{0, 4},
		},
		{
			Name:    "panic negative index",
			List:    []int{5, 6, 7},
			Indexes: []int{0, -1},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			list := NewLinkedList[int]()
			for _, element := range testCase.List {
				list.PushBack(element)
			}

			var actual []int
			for _, index := range testCase.Indexes {
				value := list.Get(index)
				actual = append(actual, value)
			}

			assert.Equal(t, testCase.Expected, actual)
		})
	}

	for _, panicTestCase := range panicTestCases {
		t.Run(panicTestCase.Name, func(t *testing.T) {
			require.Panics(t, func() {
				list := NewLinkedList[int]()
				for _, element := range panicTestCase.List {
					list.PushBack(element)
				}

				for _, index := range panicTestCase.Indexes {
					list.Get(index)
				}
			})
		})
	}
}

func TestSet(t *testing.T) {
	type Replacement[T any] struct {
		Index int
		Value T
	}

	type TestCase[T comparable] struct {
		Name         string
		List         []T
		Replacements []Replacement[T]
		ExpectedList []T
	}

	type PanicTestCase[T comparable] struct {
		Name         string
		List         []T
		Replacements []Replacement[T]
	}

	testCases := []TestCase[int]{
		{
			Name: "single element",
			List: []int{20},
			Replacements: []Replacement[int]{
				{Index: 0, Value: 25},
			},
			ExpectedList: []int{25},
		},
		{
			Name: "two elements",
			List: []int{20, 30},
			Replacements: []Replacement[int]{
				{Index: 0, Value: 25},
				{Index: 1, Value: 35},
			},
			ExpectedList: []int{25, 35},
		},
		{
			Name: "three elements",
			List: []int{20, 30, 40},
			Replacements: []Replacement[int]{
				{Index: 0, Value: 25},
				{Index: 1, Value: 35},
				{Index: 2, Value: 45},
			},
			ExpectedList: []int{25, 35, 45},
		},
	}

	panicTestCases := []PanicTestCase[int]{
		{
			Name: "panic set in empty list",
			List: make([]int, 0),
			Replacements: []Replacement[int]{
				{Index: 0, Value: 25},
			},
		},
		{
			Name: "panic set in non-empty list out of bounds",
			List: []int{20, 30},
			Replacements: []Replacement[int]{
				{Index: 3, Value: 25},
			},
		},
		{
			Name: "panic negative index",
			List: []int{20, 30},
			Replacements: []Replacement[int]{
				{Index: -1, Value: 25},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			list := NewLinkedList[int]()
			for _, element := range testCase.List {
				list.PushBack(element)
			}

			for _, replacement := range testCase.Replacements {
				list.Set(replacement.Index, replacement.Value)
			}

			assert.Equal(t, testCase.ExpectedList, toSlice(list))
		})
	}

	for _, panicTestCase := range panicTestCases {
		t.Run(panicTestCase.Name, func(t *testing.T) {
			require.Panics(t, func() {
				list := NewLinkedList[int]()
				for _, element := range panicTestCase.List {
					list.PushBack(element)
				}

				for _, replacement := range panicTestCase.Replacements {
					list.Set(replacement.Index, replacement.Value)
				}
			})
		})
	}
}

func TestRemove(t *testing.T) {
	type TestCase[T comparable] struct {
		Name     string
		List     []T
		Indexes  []int
		Expected []T
	}

	type PanicTestCase[T comparable] struct {
		Name    string
		List    []T
		Indexes []int
	}

	testCases := []TestCase[int]{
		{
			Name:     "single element",
			List:     []int{20},
			Indexes:  []int{0},
			Expected: make([]int, 0),
		},
		{
			Name:     "two elements",
			List:     []int{20, 30},
			Indexes:  []int{0},
			Expected: []int{30},
		},
		{
			Name:     "three elements",
			List:     []int{20, 30, 40},
			Indexes:  []int{1},
			Expected: []int{20, 40},
		},
		{
			Name:     "five elements",
			List:     []int{20, 30, 40, 50, 60},
			Indexes:  []int{3, 2, 1},
			Expected: []int{20, 60},
		},
	}

	panicTestCases := []PanicTestCase[int]{
		{
			Name:    "panic in empty list",
			List:    make([]int, 0),
			Indexes: []int{10},
		},
		{
			Name:    "panic in non-empty list out of bounds",
			List:    []int{20, 30},
			Indexes: []int{10},
		},
		{
			Name:    "panic by negative index",
			List:    []int{20, 30},
			Indexes: []int{-1},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			list := NewLinkedList[int]()
			for _, element := range testCase.List {
				list.PushBack(element)
			}

			for _, index := range testCase.Indexes {
				list.Remove(index)
			}

			assert.Equal(t, testCase.Expected, toSlice(list))
		})
	}

	for _, panicTestCase := range panicTestCases {
		t.Run(panicTestCase.Name, func(t *testing.T) {
			require.Panics(t, func() {
				list := NewLinkedList[int]()
				for _, element := range panicTestCase.List {
					list.PushBack(element)
				}

				for _, index := range panicTestCase.Indexes {
					list.Remove(index)
				}
			})
		})
	}
}
