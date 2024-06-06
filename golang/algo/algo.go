package algo

import "cmp"

// AllOf checks if all elements in the slice satisfy the predicate f.
func AllOf[T any](data []T, f func(T) bool) bool {
	for _, v := range data {
		if !f(v) {
			return false
		}
	}
	return true
}

// AnyOf checks if any elements in the slice satisfy the predicate f.
func AnyOf[T any](data []T, f func(T) bool) bool {
	for _, v := range data {
		if f(v) {
			return true
		}
	}
	return false
}

// NoneOf checks if none of the elements in the slice satisfy the predicate f.
func NoneOf[T any](data []T, f func(T) bool) bool {
	for _, v := range data {
		if f(v) {
			return false
		}
	}
	return true
}

// Clamp constrains a value x to the range [min, max].
func Clamp[T cmp.Ordered](x, min, max T) T {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

// MinElement returns the smallest element in a slice. If the slice is empty, it returns the zero value of the type.
func MinElement[T cmp.Ordered](data []T) (minEl T, found bool) {
	if len(data) == 0 {
		return minEl, false
	}

	minEl = data[0]
	for _, value := range data[1:] {
		if value < minEl {
			minEl = value
		}
	}
	return minEl, true
}

// MaxElement returns the largest element in a slice. If the slice is empty, it returns the zero value of the type.
func MaxElement[T cmp.Ordered](data []T) (maxEl T, found bool) {
	if len(data) == 0 {
		return maxEl, false
	}

	maxEl = data[0]
	for _, value := range data[1:] {
		if value > maxEl {
			maxEl = value
		}
	}
	return maxEl, true
}

// Map applies a function to each item in a slice and returns a new slice with the results.
func Map[T any, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// Filter returns a new slice holding only the items in s that match the predicate f.
func Filter[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce applies a function of two arguments cumulatively to the items of a slice,
// from left to right, so as to reduce the slice to a single value.
func Reduce[T any, U any](s []T, initializer U, f func(U, T) U) U {
	result := initializer
	for _, v := range s {
		result = f(result, v)
	}
	return result
}

func Reverse(s string) string {
	ans := ""
	for _, e := range s {
		ans = string(e) + ans
	}
	return ans
}
