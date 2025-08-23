package utils

import (
	"reflect"
	"testing"
)

var sortTests = []struct {
	name  string
	input []string
	want  []string
}{
	{
		name:  "empty slice",
		input: []string{},
		want:  []string{},
	},
	{
		name:  "single element",
		input: []string{"a"},
		want:  []string{"a"},
	},
	{
		name:  "sorted slice",
		input: []string{"a", "b", "c"},
		want:  []string{"a", "b", "c"},
	},
	{
		name:  "reverse sorted slice",
		input: []string{"c", "b", "a"},
		want:  []string{"a", "b", "c"},
	},
	{
		name:  "random slice",
		input: []string{"b", "c", "a"},
		want:  []string{"a", "b", "c"},
	},
	{
		name:  "slice with duplicates",
		input: []string{"b", "a", "c", "a", "b"},
		want:  []string{"a", "a", "b", "b", "c"},
	},
}

func TestMergeSort(t *testing.T) {
	tests := sortTests

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy to avoid modifying the original
			inputCopy := make([]string, len(tt.input))
			copy(inputCopy, tt.input)

			got := MergeSort(inputCopy)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	tests := sortTests

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy to avoid modifying the original
			inputCopy := make([]string, len(tt.input))
			copy(inputCopy, tt.input)

			got := QuickSort(inputCopy)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	tests := sortTests

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy to avoid modifying the original
			inputCopy := make([]string, len(tt.input))
			copy(inputCopy, tt.input)

			got := HeapSort(inputCopy)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
