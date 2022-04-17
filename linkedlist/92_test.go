package linkedlist

import (
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				head:  []int{1, 2, 3, 4, 5},
				left:  2,
				right: 4,
			},
			want: []int{1, 4, 3, 2, 5},
		},
		{
			name: "basic 2",
			args: args{
				head:  []int{5},
				left:  1,
				right: 1,
			},
			want: []int{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildListFromSlice(tt.args.head)
			want := buildListFromSlice(tt.want)

			if got := reverseBetweenIter(head, tt.args.left, tt.args.right); !got.Eq(want) {
				t.Errorf("reverseBetween() = %s, want %s", got, want)
			}
		})
	}
}

func Test_reverseNIter(t *testing.T) {
	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				head: []int{1, 2, 3, 4, 5, 6, 7},
				n:    3,
			},
			want: []int{3, 2, 1, 4, 5, 6, 7},
		},
		{
			name: "basic 2",
			args: args{
				head: []int{1, 2, 3, 4, 5, 6, 7},
				n:    7,
			},
			want: []int{7, 6, 5, 4, 3, 2, 1},
		},
		{
			name: "basic 2",
			args: args{
				head: []int{1, 2},
				n:    1,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildListFromSlice(tt.args.head)
			want := buildListFromSlice(tt.want)
			if got := reverseNIter(head, tt.args.n); !want.Eq(got) {
				t.Errorf("reverseNIter() = %s, want %s", got, want)
			}
		})
	}
}
