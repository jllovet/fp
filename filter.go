package fp

// Filter returns a new slice made by iterating over a given slice
// and checking whether each of its elements satisfy the filter function.
// If the function applied to the element returns true, then the element
// is included in the returned slice. If it returns false, then the element
// is not included in the returned slice.
func Filter[T any](f func(x T) bool, xs []T) []T {
	ys := make([]T, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}
