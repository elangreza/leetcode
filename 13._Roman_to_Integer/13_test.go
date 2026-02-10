package romantointeger

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "",
		// 	args: args{
		// 		s: "IV",
		// 	},
		// 	want: 4,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		s: "III",
		// 	},
		// 	want: 3,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		s: "LIX",
		// 	},
		// 	want: 59,
		// },
		{
			name: "",
			args: args{
				s: "MCDLXXVI",
			},
			want: 1476,
		},
		{
			name: "",
			args: args{
				s: "XXVII",
			},
			want: 27,
		},
		{
			name: "",
			args: args{
				s: "MCMXCIX",
			},
			want: 1999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
