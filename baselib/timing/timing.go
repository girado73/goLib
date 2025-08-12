// Package timing is a library to extend the timing functionality of the Go programming language
// It consists of higher order functions to work with time durations and intervals
package timing

import "time"

func StopTime[T any](fn func() T) (int64, T) {
	start := time.Now()
	result := fn()
	elapsed := time.Since(start).Milliseconds()
	return elapsed, result
}
