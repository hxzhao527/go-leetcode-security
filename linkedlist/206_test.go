package linkedlist

import (
	"testing"
)

func Test_reverseList(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "basic",
			args: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "basic 2",
			args: nil,
			want: nil,
		},
		// {
		// 	name: "error 1",
		// 	args: []int{1, 2, 3, 4},
		// 	want: []int{1, 2, 3, 4},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildListFromSlice(tt.args)
			want := buildListFromSlice(tt.want)

			if got := reverseList(head); !got.Eq(want) {
				t.Errorf("reverseList() = %s, want %s", got, want)
			}
		})
	}
}
