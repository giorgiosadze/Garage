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
