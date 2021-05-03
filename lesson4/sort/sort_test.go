package sort

import (
	"reflect"
	"testing"
)

func TestByInserts(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		want  []int64
	}{
		{name: "first", input: []int64{45, 3, -56, 0, 23, -4, -128}, want: []int64{-128, -56, -4, 0, 3, 23, 45}},
		{name: "second", input: []int64{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}, want: []int64{0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "third", input: []int64{1111, 555, -777, -333, 999, 222, -444}, want: []int64{-777, -444, -333, 222, 555, 999, 1111}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ByInserts(tc.input...)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
