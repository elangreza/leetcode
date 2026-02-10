package longestpalindromicsubstring

import (
	"testing"
)

// func Test_palindrome(t *testing.T) {
// 	type args struct {
// 		s string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		{
// 			name: "1",
// 			args: args{
// 				s: "ava",
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "1",
// 			args: args{
// 				s: "gava",
// 			},
// 			want: false,
// 		},
// 		{
// 			name: "1",
// 			args: args{
// 				s: "a",
// 			},
// 			want: true,
// 		},
// 		{
// 			name: "1",
// 			args: args{
// 				s: "aa",
// 			},
// 			want: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := palindrome(tt.args.s); got != tt.want {
// 				t.Errorf("palindrome() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_longestPalindromeV3(t *testing.T) {
	type args struct {
		ss string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				ss: "bab",
			},
			want: "bab",
		},
		{
			name: "",
			args: args{
				ss: "bb",
			},
			want: "bb",
		},
		{
			name: "",
			args: args{
				ss: "abb",
			},
			want: "bb",
		},
		{
			name: "",
			args: args{
				ss: "zzabbabbx",
			},
			want: "bbabb",
		},
		{
			name: "",
			args: args{
				ss: "vazaazav",
			},
			want: "vazaazav",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.ss); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
