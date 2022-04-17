package linkedlist

import (
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				head: []int{1, 2, 6, 3, 4, 5, 6},
				val:  6,
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildListFromSlice(tt.args.head)
			want := buildListFromSlice(tt.want)
			if got := removeElements(head, tt.args.val); !want.Eq(got) {
				t.Errorf("removeElements() = %s, want %s", got, want)
			}
		})
	}
}
