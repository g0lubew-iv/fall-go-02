package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates a slice of random non-negative integers of the given size.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	elements := make([]int, size)
	for i := range elements {
		elements[i] = rand.Int()
	}
	return elements
}

// maximum returns the maximum value in the slice.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxVal := data[0]
	for _, v := range data[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// maxChunks finds the maximum value in the slice concurrently using CHUNKS goroutines.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) < CHUNKS {
		return maximum(data)
	}

	var wg sync.WaitGroup
	results := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	for i := range CHUNKS {
		start := i * chunkSize
		end := start + chunkSize

		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(idx int, slice []int) {
			defer wg.Done()
			results[idx] = maximum(slice)
		}(i, data[start:end])
	}

	wg.Wait()

	return maximum(results)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	elements := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxVal := maximum(elements)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxVal, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxVal = maxChunks(elements)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxVal, elapsed)
}
