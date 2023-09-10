package ranges

import "testing"

func TestCommonPrefix(t *testing.T) {
	t.Parallel()

	result := String(CommonPrefix(Runes("hello, world"), Runes("hello, there"), Eq[rune]))

	assertEqual(t, result, "hello, ")
}

func TestCommonPrefixF(t *testing.T) {
	t.Parallel()

	result := String(CommonPrefixF(Runes("hello, world"), Runes("hello, there"), Eq[rune]))

	assertEqual(t, result, "hello, ")
}

func TestCommonPrefixS(t *testing.T) {
	t.Parallel()

	result := String(CommonPrefixS([]rune("hello, world"), []rune("hello, there"), Eq[rune]))

	assertEqual(t, result, "hello, ")
}
