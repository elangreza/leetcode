package a

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 4,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				nums:   []int{3, 1},
				target: 3,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums:   []int{3, 1},
				target: 1,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums:   []int{3, 5, 1},
				target: 5,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums:   []int{5, 1, 3},
				target: 3,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				target: -1,
			},
			want: -1,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				target: 1,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				target: 8,
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				target: 9,
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				target: 2,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				target: 5,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				target: 11,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
