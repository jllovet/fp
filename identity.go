package fp

// Id is the identity function. It returns whatever is passed to it as an argument.
func Id[T any](x T) T {
	return x
}
