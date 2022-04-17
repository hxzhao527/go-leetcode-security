package btree

import (
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				root: "[4,2,7,1,3,6,9]",
			},
			want: "[4,7,2,9,6,3,1]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeByLevel(tt.args.root)
			want := buildTreeByLevel(tt.want)
			if got := invertTree(root); !want.Eq(got) {
				t.Errorf("invertTree() = %v, want %v", got, want)
			}
		})
	}
}
