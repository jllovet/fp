package fp

// Map returns a new slice made by iterating over a given slice
// and applying a function to each of its elements.
// It is meant to be equivalent to the mathematical idea of a map.
func Map[T any](f func(x T) T, xs []T) []T {
	ys := make([]T, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}
