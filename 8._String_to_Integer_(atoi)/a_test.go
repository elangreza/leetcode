package a

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "42",
			},
			want: 42,
		},
		{
			name: "2",
			args: args{
				s: "   -042",
			},
			want: -42,
		},
		{
			name: "3",
			args: args{
				s: "0-1",
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				s: "1337c0d3",
			},
			want: 1337,
		},
		{
			name: "5",
			args: args{
				s: "-+12",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
