package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnqueue(t *testing.T) {
	type TestCase[TIn any] struct {
		Name  string
		Input []TIn
	}

	var inputForExtraLargeTest []int
	for i := 0; i < 890126; i++ {
		inputForExtraLargeTest = append(inputForExtraLargeTest, ((33*i)^0x03BF)>>3)
	}

	testCases := []TestCase[int]{
		{
			Name:  "full buffer",
			Input: []int{1, 5, 7, 9, 13, 22, 23, 24},
		},
		{
			Name:  "not full buffer",
			Input: []int{1, 5, 7, 9, 13, 22},
		},
		{
			Name:  "empty buffer",
			Input: []int{},
		},
		{
			Name:  "single buffer",
			Input: []int{1},
		},
		{
			Name:  "two elements",
			Input: []int{1, 5},
		},
		{
			Name:  "extra large",
			Input: inputForExtraLargeTest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			dequeued := make([]int, 0)
			queue := NewQueue[int]()

			for _, input := range testCase.Input {
				queue.Enqueue(input)
			}

			queueLengthFull := queue.Len()

			for i := 0; i < len(testCase.Input); i++ {
				d := queue.Dequeue()
				dequeued = append(dequeued, d)
			}

			queueLengthEmpty := queue.Len()

			assert.Equal(t, 0, queueLengthEmpty)
			assert.Equal(t, len(testCase.Input), queueLengthFull)
			assert.Equal(t, testCase.Input, dequeued)
		})
	}
}

func TestDequeue(t *testing.T) {
	require.Panics(t, func() {
		queue := NewQueue[int]()

		queue.Dequeue()
	})
}

func TestPeek(t *testing.T) {
	t.Run("empty queue", func(t *testing.T) {
		require.Panics(t, func() {
			queue := NewQueue[int]()

			queue.Peek()
		})
	})

	type TestCase[T any] struct {
		Name     string
		Input    []T
		Expected T
	}

	testCases := []TestCase[int]{
		{
			Name:     "single element",
			Input:    []int{5},
			Expected: 5,
		},
		{
			Name:     "two elements",
			Input:    []int{5, 10},
			Expected: 5,
		},
		{
			Name:     "three elements",
			Input:    []int{5, 10, 11},
			Expected: 5,
		},
		{
			Name:     "many elements",
			Input:    []int{5, 10, 11, 12, 13, 14, 15},
			Expected: 5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			queue := NewQueue[int]()

			for _, element := range testCase.Input {
				queue.Enqueue(element)
			}
			actual := queue.Peek()

			assert.Equal(t, testCase.Expected, actual)
		})
	}
}

func TestGrow(t *testing.T) {
	t.Run("empty queue", func(t *testing.T) {
		queue := NewQueue[int]()
		expectedQueueCap := 3

		queue.Grow(expectedQueueCap)
		actualQueueCap := queue.Cap()

		assert.Equal(t, expectedQueueCap, actualQueueCap)
	})

	t.Run("negative value", func(t *testing.T) {
		require.Panics(t, func() {
			queue := NewQueue[int]()
			expectedQueueCap := -1

			queue.Grow(expectedQueueCap)
		})
	})

	t.Run("typical use case", func(t *testing.T) {
		var contents []int
		for i := 0; i < 4997; i++ {
			contents = append(contents, (3*i)>>1)
		}

		queue := NewQueue[int]()
		expectedQueueCap := 4999

		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Grow(expectedQueueCap)

		// queue.Cap() should be the same as expectedQueueCap
		queueCapAfterGrow := queue.Cap()

		//  add all contents to the queue, shouldn't cause reallocation
		for _, content := range contents {
			queue.Enqueue(content)
		}

		// queue.Cap() should be the same as expectedQueueCap
		queueCapAfterEnqueue := queue.Cap()

		assert.Equal(t, expectedQueueCap, queueCapAfterGrow)
		assert.Equal(t, expectedQueueCap, queueCapAfterEnqueue)
	})
}
