package a

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ss   string
		want int
	}{
		{
			name: "first",
			ss:   "abcabcbb",
			want: 3,
		},
		{
			name: "second",
			ss:   "bbbbb",
			want: 1,
		},
		{
			name: "third",
			ss:   "pwwkew",
			want: 3,
		},
		{
			name: "fourth",
			ss:   "bdb",
			want: 2,
		},
		{
			name: "fifth",
			ss:   "dvdf",
			want: 3,
		},
		{
			name: "sixth",
			ss:   "abba",
			want: 2,
		},
		{
			name: "seven",
			ss:   " ",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.ss)
			t.Log(got)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
