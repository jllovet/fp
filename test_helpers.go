package fp

import (
	"reflect"
	"testing"
)

func assertEqual[T any](t *testing.T, got, want T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertUnequal[T any](t *testing.T, got, want T) {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
