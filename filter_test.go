package fp

import (
	"fmt"
	"testing"
)

func ExampleFilter() {
	isEven := func(x int) bool { return x%2 == 0 }
	ns := []int{1, 2, 3, 4, 5}
	xs := Filter(isEven, ns)
	fmt.Println(xs)
	// Output: [2 4]
}

func TestFilter(t *testing.T) {
	t.Run("Filter with trivial true func returns original slice", func(t *testing.T) {
		got := Filter(func(x int) bool { return true }, []int{1, 2, 3, 4, 5})
		want := []int{1, 2, 3, 4, 5}
		assertEqual(t, got, want)
	})
	t.Run("Filter with trivial false func returns empty slice", func(t *testing.T) {
		got := Filter(func(x int) bool { return false }, []int{1, 2, 3, 4, 5})
		want := []int{}
		assertEqual(t, got, want)
	})
	t.Run("Filter for even ints returns slice with even ints", func(t *testing.T) {
		got := Filter(func(x int) bool { return x%2 == 0 }, []int{1, 2, 3, 4, 5})
		want := []int{2, 4}
		assertEqual(t, got, want)
	})
	t.Run("Filter for structs with non-empty field returns slice with those structs", func(t *testing.T) {
		type entity struct {
			a string
		}
		got := Filter(func(x entity) bool { return x.a != "" }, []entity{{a: ""}, {a: "foo"}, {a: "bar"}})
		want := []entity{{a: "foo"}, {a: "bar"}}
		assertEqual(t, got, want)
	})
}
