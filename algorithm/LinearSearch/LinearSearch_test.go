package linearsearch

import (
	"reflect"
	"testing"
)

// Linear Search given an array and target number, return bool
func Test_LinearSearch(t *testing.T) {
	// place input parameter
	type args struct {
		inputArr []int
		target   int
	}
	// Place information of test, including return value
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{
				inputArr: []int{95, 78, 46, 58, 45, 86, 99, 251, 320},
				target:   58,
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				inputArr: []int{95, 78, 46, 58, 45, 86, 99, 251, 320},
				target:   8,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinearSearch(tt.args.inputArr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Linear Search got is %v, want %v", got, tt.want)
			}
		})
	}

}
