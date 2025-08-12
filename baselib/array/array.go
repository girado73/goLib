// Package array is a library to extend the array / slice functionality of the Go programming language
// It orients itself on the php base library and provides similar functions
package array

import (
	"cmp"
	"math/rand"
	"slices"
	"sync"
)

func End[T any](array []T) T {
	return array[len(array)-1]
}

func Filter[T any](array []T, fn func(T) bool) []T {
	var result []T
	for _, v := range array {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[T any, R any](array []T, fn func(T) R) []R {
	var result []R
	for _, v := range array {
		result = append(result, fn(v))
	}
	return result
}

// MapParallel applies a function to each element of a slice in parallel
// Only use this function if the function is not dependent on the order of execution
// Function is only useful for big arrays or slices
// This is still saldy really slow
func MapParallel[T, R any](input []T, fn func(T) R) []R {
	result := make([]R, len(input))
	var wg sync.WaitGroup

	for i, item := range input {
		wg.Add(1)
		go func(index int, value T) {
			defer wg.Done()
			result[index] = fn(value)
		}(i, item)
	}

	wg.Wait()
	return result
}

func Merge[T any, R any](array1 []T, array2 []R) []any {
	var result []any
	for _, v := range array1 {
		result = append(result, v)
	}
	for _, v := range array2 {
		result = append(result, v)
	}
	return result
}

func MergeMultiple[T any](arrays ...[]T) []T {
	var result []T
	for _, array := range arrays {
		result = append(result, array...)
	}
	return result
}

func All[T any](array []T, fn func(T) bool) bool {
	for _, v := range array {
		if !fn(v) {
			return false
		}
	}
	return true
}

func Any[T any](array []T, fn func(T) bool) bool {
	return slices.ContainsFunc(array, fn)
}

func Pop[T any](array []T) (T, []T) {
	if len(array) == 0 {
		var zero T
		return zero, array
	}
	last := array[len(array)-1]
	array = array[:len(array)-1]
	return last, array
}

// Search finds the index of the first occurrence of a value in an array.
// Use for unsorted arrays.
// Brute force search algorithm
func Search[T comparable](array []T, value T) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

// Replace replaces the first occurrence of a value in an array with a new value.
// Returns the modified array.
// When not found, returns the original array.
func Replace[T comparable](array []T, oldValue T, newValue T) []T {
	searchIndex := Search(array, oldValue)

	if searchIndex == -1 {
		return array // Value not found, return original array
	} else {
		array[searchIndex] = newValue
		return array
	}
}

// SearchSorted finds the index of a value in a sorted array.
// Use for sorted arrays.
// Uses binary search algorithm for efficiency.
func SearchSorted[T cmp.Ordered](array []T, value T) int {
	// Binary search for sorted arrays
	low, high := 0, len(array)-1
	for low <= high {
		mid := (low + high) / 2
		if array[mid] < value {
			low = mid + 1
		} else if array[mid] > value {
			high = mid - 1
		} else {
			return mid // Found the value
		}
	}
	return -1 // Not found
}

func Sort[T cmp.Ordered](array []T) {
	slices.Sort(array)
}

func Rand[T any](array []T) T {
	if len(array) == 0 {
		var zero T
		return zero
	}
	x := rand.Int31n(int32(len(array)))
	return array[x]
}

func Reduce[T any, R any](array []T, fn func(R, T) R, initial R) R {
	result := initial
	for _, v := range array {
		result = fn(result, v)
	}
	return result
}

func Reverse[T any](array []T) []T {
	result := make([]T, len(array))

	for _, v := range array {
		result[len(array)-1] = v
		array = array[:len(array)-1]
	}
	return result
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Product[T Number](array []T) T {
	var result T = 1
	for _, v := range array {
		result *= v
	}
	return result
}

// ProductParallel calculates the product of all elements in an array in parallel.
// This is sadly really slow
func ProductParallel[T Number](array []T) T {
	result := T(1)
	var wg sync.WaitGroup

	for _, item := range array {
		wg.Add(1)
		go func(value T) {
			defer wg.Done()
			result *= value
		}(item)
	}

	wg.Wait()
	return result
}

func Difference[T comparable](array1, array2 []T) []T {
	var result []T
	for _, v := range array1 {
		if !slices.Contains(array2, v) {
			result = append(result, v)
		}
	}
	return result
}

func Compare[T comparable](array1, array2 []T) bool {
	if len(array1) != len(array2) {
		return false
	}
	for i, v := range array1 {
		if v != array2[i] {
			return false
		}
	}
	return true
}
