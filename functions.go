package fp

type Transform[T any] func(x T) T
type Predicate[T any] func(x T) bool
