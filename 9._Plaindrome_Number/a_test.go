package plaindromenumber

import "testing"

func Test_isPalindromeV2(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				x: 23456,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				x: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeV2(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
