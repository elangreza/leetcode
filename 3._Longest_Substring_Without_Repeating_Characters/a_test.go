package a

import "testing"

func Test_lengthOfLongestSubstring2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s: "zxyzxyz",
			},
			want: 3,
		},

		{
			name: "",
			args: args{
				s: "thequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsovert",
			},
			want: 17,
		},

		{
			name: "",
			args: args{
				s: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
