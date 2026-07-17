package linkedlist

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty list", []int{}, nil},
		{"single node", []int{1}, []int{1}},
		{"multiple nodes", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"two nodes", []int{1, 2}, []int{2, 1}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			head := FromSlice(tc.in)
			got := ToSlice(Reverse(head))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Reverse(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
