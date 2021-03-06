package reverse_linked_list

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestReverseLinkedList1(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect *ListNode
	}{
		{SliceToList([]int{}), SliceToList([]int{})},
		{SliceToList([]int{1}), SliceToList([]int{1})},
		{SliceToList([]int{1, 2, 3, 4, 5}), SliceToList([]int{5, 4, 3, 2, 1})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := ReverseLinkedList1(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					ListToSlice(c.expect), got, ListToSlice(c.input))
			}
		})
	}
}

func TestReverseLinkedList2(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect *ListNode
	}{
		{SliceToList([]int{}), SliceToList([]int{})},
		{SliceToList([]int{1}), SliceToList([]int{1})},
		{SliceToList([]int{1, 2, 3, 4, 5}), SliceToList([]int{5, 4, 3, 2, 1})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := ReverseLinkedList2(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					ListToSlice(c.expect), got, ListToSlice(c.input))
			}
		})
	}
}
