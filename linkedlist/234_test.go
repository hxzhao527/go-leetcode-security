package linkedlist

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "basic",
			args: args{
				head: []int{1, 2, 2, 1},
			},
			want: true,
		},
		{
			name: "basic 2",
			args: args{
				head: []int{1, 2, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildListFromSlice(tt.args.head)
			if got := isPalindrome(head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
