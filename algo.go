package algo

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
func Clamp[T comparable](x, min, max T) T {
    if x < min {
        return min
    }
    if x > max {
        return max
    }
    return x
}

// Max returns the larger of a and b for any orderable type.
func Max[T comparable](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Min returns the smaller of a and b for any orderable type.
func Min[T comparable](a, b T) T {
    if a < b {
        return a
    }
    return b
}

