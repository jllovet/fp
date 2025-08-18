package fp

import (
	"fmt"
	"testing"
)

func ExampleId() {
	x := Id(1)
	fmt.Println(x)
	// Output: 1

}

func TestId(t *testing.T) {
	t.Run("Identity func returns input for int", func(t *testing.T) {
		got := Id(1)
		want := 1
		assertEqual(t, got, want)
	})
	t.Run("Identity func returns input for bool", func(t *testing.T) {
		got := Id(true)
		want := true
		assertEqual(t, got, want)
	})
	t.Run("Identity func returns input for slice", func(t *testing.T) {
		got := Id([]int{})
		want := []int{}
		assertEqual(t, got, want)
	})
	t.Run("Identity func returns input for struct", func(t *testing.T) {
		type entity struct {
			a string
		}
		got := Id(entity{a: "I am"})
		want := entity{a: "I am"}
		assertEqual(t, got, want)
	})
	t.Run("Identity func returns input for empty struct", func(t *testing.T) {
		type entity struct {
		}
		got := Id(entity{})
		want := entity{}
		assertEqual(t, got, want)
	})
	t.Run("Identity func does not modify argument", func(t *testing.T) {
		got := Id(0)
		want := 1
		assertUnequal(t, got, want)
	})
	t.Run("Identity func applied to Identity func returns Identity func", func(t *testing.T) {
		f := Id(Id[any])
		g := Id[any]
		// Function comparison is not allowed in Go, so we will need to compare the functions behavior.
		// Maybe there is some unsafe.Pointer dark magic that can be done to determine whether they're equal.
		// Note: A downside to this approach is that there might be some transformation to the original
		// function that produces side effects which are not caught here.
		got := f(1)
		want := g(1)
		assertEqual(t, got, want)
	})
}
