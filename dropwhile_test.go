package ranges

import "testing"

func TestDropWhile(t *testing.T) {
	t.Parallel()

	result := Slice(DropWhile(I(F(B(Only(1, 2, 3, 4)))), func(x int) bool { return x < 3 }))

	assertEqual(t, result, []int{3, 4})
}

func TestDropWhileF(t *testing.T) {
	t.Parallel()

	assertHasSaveableFront(t, DropWhileF(F(B(Only(1, 2, 3, 4))), func(x int) bool { return x < 3 }), 3)
}
