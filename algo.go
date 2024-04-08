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
