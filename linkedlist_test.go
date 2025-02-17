package collections

import (
	"slices"
	"testing"
)

// TODO: add more tests to cover new features and change existing

func TestLinkedList_Prepend(t *testing.T) {
	t.Parallel()

	var list = NewLinkedList[int]()
	var contents = []int{1, 20, 35, 49, 51, 63}
	var listContents = make([]int, 0, len(contents))

	for i := 0; i < len(contents); i++ {
		list.Prepend(contents[i])
	}

	for p := list.Head; p != nil; p = p.Next {
		listContents = append(listContents, p.Value)
	}

	reverseListContents := make([]int, len(listContents))
	copy(reverseListContents, listContents)
	slices.Reverse(reverseListContents)

	if !slices.Equal(reverseListContents, contents) {
		t.Errorf("contents (%v) and list contents (%v) do not match", contents, listContents)
	}
}

func TestLinkedList_Append(t *testing.T) {
	t.Parallel()

	var list = NewLinkedList[int]()
	var contents = []int{1, 20, 35, 49, 51, 63}
	var listContents = make([]int, 0, len(contents))

	for i := 0; i < len(contents); i++ {
		list.Append(contents[i])
	}

	for p := list.Head; p != nil; p = p.Next {
		listContents = append(listContents, p.Value)
	}

	if !slices.Equal(listContents, contents) {
		t.Errorf("contents (%v) and list contents (%v) do not match", contents, listContents)
	}
}

func TestLinkedList_RemoveFromHead_NotAll(t *testing.T) {
	t.Parallel()

	var list = NewLinkedList[int]()
	var contents = []int{1, 20, 35, 49, 51, 63}
	var contentsAfterRemoveExpected = []int{49, 51, 63}
	var contentsAfterRemove = make([]int, 0, len(contentsAfterRemoveExpected))
	var err error

	for i := 0; i < len(contents); i++ {
		list.Append(contents[i])
	}

	for i := 0; i < len(contents)-len(contentsAfterRemoveExpected); i++ {
		_, err = list.RemoveFromHead()

		if err != nil {
			t.Errorf("Remove from head returned error: %s", err.Error())
			return
		}
	}

	for p := list.Head; p != nil; p = p.Next {
		contentsAfterRemove = append(contentsAfterRemove, p.Value)
	}

	if !slices.Equal(contentsAfterRemove, contentsAfterRemoveExpected) {
		t.Errorf("After remove slice (%v) differs from expected (%v)", contentsAfterRemove, contentsAfterRemoveExpected)
	}
}

func TestLinkedList_RemoveFromHead_AllButOne(t *testing.T) {
	t.Parallel()

	var list = NewLinkedList[int]()
	var contents = []int{1, 20, 35, 49, 51, 63}
	var contentsAfterRemoveExpected = []int{63}
	var contentsAfterRemove = make([]int, 0, len(contentsAfterRemoveExpected))
	var err error

	for i := 0; i < len(contents); i++ {
		list.Append(contents[i])
	}

	for i := 0; i < len(contents)-len(contentsAfterRemoveExpected); i++ {
		_, err = list.RemoveFromHead()

		if err != nil {
			t.Errorf("Remove from head returned error: %s", err.Error())
			return
		}
	}

	for p := list.Head; p != nil; p = p.Next {
		contentsAfterRemove = append(contentsAfterRemove, p.Value)
	}

	if !slices.Equal(contentsAfterRemove, contentsAfterRemoveExpected) {
		t.Errorf("After remove slice (%v) differs from expected (%v)", contentsAfterRemove, contentsAfterRemoveExpected)
	}
}

func TestLinkedList_RemoveFromHead_All(t *testing.T) {
	t.Parallel()

	var list = NewLinkedList[int]()
	var contents = []int{1, 20, 35, 49, 51, 63}
	var contentsAfterRemoveExpected []int
	var contentsAfterRemove = make([]int, 0, len(contentsAfterRemoveExpected))
	var err error

	for i := 0; i < len(contents); i++ {
		list.Append(contents[i])
	}

	for i := 0; i < len(contents)-len(contentsAfterRemoveExpected); i++ {
		_, err = list.RemoveFromHead()

		if err != nil {
			t.Errorf("Remove from head returned error: %s", err.Error())
			return
		}
	}

	for p := list.Head; p != nil; p = p.Next {
		contentsAfterRemove = append(contentsAfterRemove, p.Value)
	}

	if !slices.Equal(contentsAfterRemove, contentsAfterRemoveExpected) {
		t.Errorf("After remove slice (%v) differs from expected (%v)", contentsAfterRemove, contentsAfterRemoveExpected)
	}
}

func TestLinkedList_RemoveFromTail_NotAll(t *testing.T) {
	t.Parallel()

	var list = NewLinkedList[int]()
	var contents = []int{1, 20, 35, 49, 51, 63}
	var contentsAfterRemoveExpected = []int{1, 20, 35}
	var contentsAfterRemove = make([]int, 0, len(contentsAfterRemoveExpected))
	var err error

	for i := 0; i < len(contents); i++ {
		list.Append(contents[i])
	}

	for i := 0; i < len(contents)-len(contentsAfterRemoveExpected); i++ {
		_, err = list.RemoveFromTail()

		if err != nil {
			t.Errorf("Remove from head returned error: %s", err.Error())
			return
		}
	}

	for p := list.Head; p != nil; p = p.Next {
		contentsAfterRemove = append(contentsAfterRemove, p.Value)
	}

	if !slices.Equal(contentsAfterRemove, contentsAfterRemoveExpected) {
		t.Errorf("After remove slice (%v) differs from expected (%v)", contentsAfterRemove, contentsAfterRemoveExpected)
	}
}

func TestLinkedList_RemoveFromTail_AllButOne(t *testing.T) {
	t.Parallel()

	var list = NewLinkedList[int]()
	var contents = []int{1, 20, 35, 49, 51, 63}
	var contentsAfterRemoveExpected = []int{1}
	var contentsAfterRemove = make([]int, 0, len(contentsAfterRemoveExpected))
	var err error

	for i := 0; i < len(contents); i++ {
		list.Append(contents[i])
	}

	for i := 0; i < len(contents)-len(contentsAfterRemoveExpected); i++ {
		_, err = list.RemoveFromTail()

		if err != nil {
			t.Errorf("Remove from head returned error: %s", err.Error())
			return
		}
	}

	for p := list.Head; p != nil; p = p.Next {
		contentsAfterRemove = append(contentsAfterRemove, p.Value)
	}

	if !slices.Equal(contentsAfterRemove, contentsAfterRemoveExpected) {
		t.Errorf("After remove slice (%v) differs from expected (%v)", contentsAfterRemove, contentsAfterRemoveExpected)
	}
}

func TestLinkedList_RemoveFromTail_All(t *testing.T) {
	t.Parallel()

	var list = NewLinkedList[int]()
	var contents = []int{1, 20, 35, 49, 51, 63}
	var contentsAfterRemoveExpected []int
	var contentsAfterRemove = make([]int, 0, len(contentsAfterRemoveExpected))
	var err error

	for i := 0; i < len(contents); i++ {
		list.Append(contents[i])
	}

	for i := 0; i < len(contents)-len(contentsAfterRemoveExpected); i++ {
		_, err = list.RemoveFromTail()

		if err != nil {
			t.Errorf("Remove from head returned error: %s", err.Error())
			return
		}
	}

	for p := list.Head; p != nil; p = p.Next {
		contentsAfterRemove = append(contentsAfterRemove, p.Value)
	}

	if !slices.Equal(contentsAfterRemove, contentsAfterRemoveExpected) {
		t.Errorf("After remove slice (%v) differs from expected (%v)", contentsAfterRemove, contentsAfterRemoveExpected)
	}
}

func TestLinkedList_Len_AfterAddSome(t *testing.T) {
	// arrange
	var list = NewLinkedList[int]()
	var contents = []int{1, 20, 35, 49, 51, 63}
	var expectedLen = len(contents)
	var listLen int

	//act
	for i := 0; i < len(contents); i++ {
		list.Append(contents[i])
	}

	listLen = list.Len()

	//assert
	if listLen != expectedLen {
		t.Errorf("Len of list (%d) is expected to be (%d)", contents, expectedLen)
	}
}
