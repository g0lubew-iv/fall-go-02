package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		wantSize int
	}{
		{
			name:     "отрицательный размер",
			size:     -10,
			wantSize: 0,
		},
		{
			name:     "нулевой размер",
			size:     0,
			wantSize: 0,
		},
		{
			name:     "один элемент",
			size:     1,
			wantSize: 1,
		},
		{
			name:     "стандартный размер",
			size:     100,
			wantSize: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.wantSize {
				t.Errorf("generateRandomElements(%d) длина = %d, ожидалось %d", tt.size, len(got), tt.wantSize)
			}

			for i, v := range got {
				if v < 0 {
					t.Errorf("generateRandomElements(%d) элемент [%d] = %d, ожидалось положительное число", tt.size, i, v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "nil слайс",
			data: nil,
			want: 0,
		},
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
		{
			name: "один элемент",
			data: []int{42},
			want: 42,
		},
		{
			name: "все элементы одинаковые",
			data: []int{7, 7, 7, 7, 7},
			want: 7,
		},
		{
			name: "максимум в начале",
			data: []int{100, 50, 20, 10, 1},
			want: 100,
		},
		{
			name: "максимум в конце",
			data: []int{1, 10, 20, 50, 100},
			want: 100,
		},
		{
			name: "максимум в середине",
			data: []int{10, 20, 999, 50, 30},
			want: 999,
		},
		{
			name: "слайс с нулями",
			data: []int{0, 0, 0},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "nil слайс",
			data: nil,
			want: 0,
		},
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
		{
			name: "один элемент (меньше CHUNKS)",
			data: []int{55},
			want: 55,
		},
		{
			name: "элементов меньше чем CHUNKS",
			data: []int{1, 5, 3, 9, 2},
			want: 9,
		},
		{
			name: "количество элементов равно CHUNKS (8)",
			data: []int{10, 20, 30, 80, 50, 60, 70, 40},
			want: 80,
		},
		{
			name: "некратное CHUNKS количество элементов (максимум в хвосте)",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 999},
			want: 999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}
