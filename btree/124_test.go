package btree

import "testing"

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				root: "[1, 2, 3]",
			},
			want: 6,
		},
		{
			name: "fix 1",
			args: args{
				root: "[1, -2, 3]",
			},
			want: 4,
		},
		{
			name: "basic 1",
			args: args{
				root: "[-10,9,20,null,null,15,7]",
			},
			want: 42,
		},
		{
			name: "fix 2",
			args: args{
				root: "[-1,-2,10,-6,null,-3,-6]",
			},
			want: 10,
		},
		{
			name: "fix 3",
			args: args{
				root: "[-3]",
			},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeByLevel(tt.args.root)
			if got := maxPathSum(root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
