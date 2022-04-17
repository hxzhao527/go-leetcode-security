package linkedlist

import (
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    2,
			},
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "basic 2",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    3,
			},
			want: []int{3, 2, 1, 4, 5},
		},
		{
			name: "basic 3",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    1,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "basic 4",
			args: args{
				head: []int{1},
				k:    1,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildListFromSlice(tt.args.head)
			want := buildListFromSlice(tt.want)
			if got := reverseKGroup(head, tt.args.k); !want.Eq(got) {
				t.Errorf("reverseKGroup() = %s, want %s", got, want)
			}
		})
	}
}
