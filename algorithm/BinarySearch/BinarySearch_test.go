package binarysearch

import (
	"reflect"
	"testing"
)

//A binary search is a search strategy used to find elements
//within a list by consistently reducing the amount of data
//to be searched and thereby increasing the rate at which
//the search term is found. To use a binary search algorithm,
//the list to be operated on must have already been sorted.

// input array, target int, return bool

func Test_BinarySearch(t *testing.T) {

	type args struct {
		inputArr []int
		target   int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{
				inputArr: []int{20, 31, 45, 63, 70, 100, 1, 2, 9},
				target:   63,
			},
			want: true,
		},
		{
			name: "Example1",
			args: args{
				inputArr: []int{20, 31, 40, 1, 2, 9},
				target:   53,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.inputArr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Binary Search %v want %v", got, tt.want)
			}
		})
	}

}
