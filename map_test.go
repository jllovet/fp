package fp

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func ExampleMap() {
	addOne := func(x int) int { return x + 1 }
	ns := []int{0, 1, 2, 3}
	xs := Map(addOne, ns)
	xStrings := make([]string, len(xs))
	for i, x := range xs {
		xStrings[i] = strconv.Itoa(x)
	}
	fmt.Println(strings.Join(xStrings, ", "))
	// Output: 1, 2, 3, 4
}

func TestMap(t *testing.T) {
	t.Run("Map with Identity func returns original slice", func(t *testing.T) {
		got := Map(Id, []int{1, 2, 3, 4, 5})
		want := []int{1, 2, 3, 4, 5}
		assertEqual(t, got, want)
	})
	t.Run("Map with arithmetic func returns mapped slice", func(t *testing.T) {
		got := Map(func(x int) int { return x + 1 }, []int{1, 2, 3, 4, 5})
		want := []int{2, 3, 4, 5, 6}
		assertEqual(t, got, want)
	})
	t.Run("Map returns slice of same length", func(t *testing.T) {
		got := Map(Id, []int{1, 2, 3, 4, 5})
		want := []int{1, 2, 3, 4, 5, 6}
		assertUnequal(t, got, want)
	})
	t.Run("Map does not modify the original slice", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5}
		got := Map(func(x int) int { return x + 1 }, original)
		want := []int{1, 2, 3, 4, 5}
		assertUnequal(t, got, original)
		assertEqual(t, want, original)
	})
}
