package fibonacci

import (
	"reflect"
	"testing"
)

type testStruct struct {
	name  string
	input uint64
	want  uint64
}

func TestFindRecursive(t *testing.T) {
	tests := []testStruct{
		{name: "first", input: 20, want: 4181},
		{name: "second", input: 45, want: 701408733},
		{name: "third", input: 12, want: 89},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FindRecursive(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []testStruct{
		{name: "first", input: 30, want: 514229},
		{name: "second", input: 100, want: 16008811023750101250},
		{name: "third", input: 45, want: 701408733},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Find(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

var findRecursiveRes uint64

func BenchmarkFindRecursive(t *testing.B) {
	var res uint64

	for i := 0; i < t.N; i++ {
		res = FindRecursive(30)
	}

	findRecursiveRes = res
}

var findRes uint64

func BenchmarkFind30(t *testing.B) {
	var res uint64

	for i := 0; i < t.N; i++ {
		res = Find(30)
	}

	findRes = res
}
