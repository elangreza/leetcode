package binarysearch

import (
	"reflect"
	"testing"
)

// func Test_binarySearch(t *testing.T) {
// 	type args struct {
// 		arr    []int
// 		target int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			name: "",
// 			args: args{
// 				arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
// 				target: -1,
// 			},
// 			want: -1,
// 		},
// 		{
// 			name: "",
// 			args: args{
// 				arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
// 				target: 1,
// 			},
// 			want: 0,
// 		},
// 		{
// 			name: "",
// 			args: args{
// 				arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
// 				target: 8,
// 			},
// 			want: 7,
// 		},
// 		{
// 			name: "",
// 			args: args{
// 				arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
// 				target: 9,
// 			},
// 			want: 8,
// 		},
// 		{
// 			name: "",
// 			args: args{
// 				arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
// 				target: 2,
// 			},
// 			want: 1,
// 		},
// 		{
// 			name: "",
// 			args: args{
// 				arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
// 				target: 5,
// 			},
// 			want: 4,
// 		},
// 		{
// 			name: "",
// 			args: args{
// 				arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
// 				target: 11,
// 			},
// 			want: -1,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := binarySearch(tt.args.arr, 0, len(tt.args.arr), tt.args.target); got != tt.want {
// 				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	tests := []struct {
		target int
		want   int
	}{
		{1, 0},   // first element
		{5, 4},   // middle element
		{9, 8},   // last element
		{10, -1}, // beyond end (triggers infinite loop in your version!)
		{0, -1},  // before start
		{6, 5},   // arbitrary element
	}

	for _, tt := range tests {
		got := binarySearch(arr, 0, len(arr), tt.target)
		if got != tt.want {
			t.Errorf("binarySearch(%d) = %d; want %d", tt.target, got, tt.want)
		}
	}
}

func TestBinarySearchIterative(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	tests := []struct {
		target int
		want   int
	}{
		{1, 0},   // first element
		{5, 4},   // middle element
		{9, 8},   // last element
		{10, -1}, // beyond end (triggers infinite loop in your version!)
		{0, -1},  // before start
		{6, 5},   // arbitrary element
	}

	for _, tt := range tests {
		got := binarySearchIterative(arr, tt.target)
		if got != tt.want {
			t.Errorf("binarySearchIterative(%d) = %d; want %d", tt.target, got, tt.want)
		}
	}
}

func Test_binarySearchIterativeWithDuplicate(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				arr:    []int{10, 20, 20, 20, 30},
				target: 20,
			},
			want: []int{1, 3},
		},
		// {
		// 	name: "",
		// 	args: args{
		// 		arr:    []int{5, 7, 7, 8, 8, 10},
		// 		target: 8,
		// 	},
		// 	want: []int{3, 4},
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		arr:    []int{5, 7, 7, 8, 8, 10},
		// 		target: 6,
		// 	},
		// 	want: []int{-1, -1},
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		arr:    []int{},
		// 		target: 0,
		// 	},
		// 	want: []int{-1, -1},
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		arr:    []int{1},
		// 		target: 1,
		// 	},
		// 	want: []int{0, 0},
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		arr:    []int{2, 2},
		// 		target: 2,
		// 	},
		// 	want: []int{0, 1},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchIterativeWithDuplicate(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binarySearchIterativeWithDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
