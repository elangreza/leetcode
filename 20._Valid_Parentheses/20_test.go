package validparentheses

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s: "(([))",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				s: "(([])[])",
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				s: "]",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidV2(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
