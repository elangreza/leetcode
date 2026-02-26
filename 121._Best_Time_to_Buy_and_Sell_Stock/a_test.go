package a

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				prices: []int{10, 1, 5, 6, 7, 1},
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				prices: []int{10, 8, 7, 5, 2},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				prices: []int{3, 4, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
