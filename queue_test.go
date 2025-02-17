package collections

import (
	"slices"
	"testing"
)

func testEnqueue[T comparable](t *testing.T, contents []T) {
	//arrange
	var dequeued []T
	var queue = NewQueue[T]()
	var queueLengthFull, queueLengthEmpty int

	//act
	for i := 0; i < len(contents); i++ {
		queue.Enqueue(contents[i])
	}

	queueLengthFull = queue.Len()

	for i := 0; i < len(contents); i++ {
		d, err := queue.Dequeue()
		if err != nil {
			t.Error(err)
			return
		}

		dequeued = append(dequeued, d)
	}

	queueLengthEmpty = queue.Len()

	//assert
	if !slices.Equal(dequeued, contents) {
		t.Error("slices are not equal; contents slice: ", contents, " dequeued slice: ", dequeued)
	}

	if queueLengthFull != len(contents) {
		t.Error("filled queue length is unexpected: got ", queueLengthFull, " expected ", len(contents))
	}

	if queueLengthEmpty != 0 {
		t.Error("empty queue length is unexpected: got ", queueLengthEmpty, " expected ", 0)
	}
}

func TestQueue_Enqueue_FullBuffer(t *testing.T) {
	var contents = []int{1, 5, 7, 9, 13, 22, 23, 24}
	testEnqueue(t, contents)
}

func TestQueue_Enqueue_NotFullBuffer(t *testing.T) {
	var contents = []int{1, 5, 7, 9, 13, 22}
	testEnqueue(t, contents)
}

func TestQueue_Enqueue_EmptyBuffer(t *testing.T) {
	var contents []int
	testEnqueue(t, contents)
}

func TestQueue_Enqueue_SingleBuffer(t *testing.T) {
	var contents = []int{1}
	testEnqueue(t, contents)
}

func TestQueue_Enqueue_TwoElementsBuffer(t *testing.T) {
	var contents = []int{10, 20}
	testEnqueue(t, contents)
}

func TestQueue_Enqueue_ExtraLarge(t *testing.T) {
	var contents []int

	for i := 0; i < 89012; i++ {
		contents = append(contents, ((33*i)^0x03BF)>>3)
	}

	testEnqueue(t, contents)
}

func TestQueue_Grow(t *testing.T) {
	//arrange
	var contents []int
	for i := 0; i < 4997; i++ {
		contents = append(contents, (3*i)>>1)
	}

	var queue = NewQueue[int]()
	var queueCapAfterGrow, queueCapAfterEnqueue int
	var expectedQueueCap = 4999

	//act
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Grow(expectedQueueCap)

	queueCapAfterGrow = queue.Cap()

	for i := 0; i < len(contents); i++ {
		queue.Enqueue(contents[i])
	}

	queueCapAfterEnqueue = queue.Cap()

	//assert
	if queueCapAfterGrow != queueCapAfterEnqueue {
		t.Errorf("queue capacity after enquqeue (%d) is not equal to capacity before (%d)", queueCapAfterEnqueue, queueCapAfterGrow)
	}

	if queueCapAfterEnqueue != expectedQueueCap {
		t.Errorf("unexpected queue capacity: got %d, expected: %d", queueCapAfterEnqueue, expectedQueueCap)
	}
}
