package ranges

import "testing"

func TestPadRight(t *testing.T) {
	t.Parallel()

	assertEqual(t, String(PadRight[rune](Runes("abc"), 'x', 5)), "abcxx")
	assertEqual(t, String(PadRight[rune](Runes("abcdef"), 'x', 5)), "abcdef")
	assertEqual(t, String(PadRight[rune](Runes(""), 'x', 5)), "xxxxx")
	assertEqual(t, String(PadRight[rune](Runes(""), 'x', 0)), "")
	assertEqual(t, String(PadRight[rune](Runes(""), 'x', -10)), "")
}

func TestPadRightF(t *testing.T) {
	t.Parallel()

	assertEqual(t, String(PadRightF(Runes("abc"), 'x', 5)), "abcxx")
	assertEqual(t, String(PadRightF(Runes("abcdef"), 'x', 5)), "abcdef")

	r := PadRightF(Runes("日"), '本', 3)

	r2 := r.Save()

	r.PopFront()

	r3 := r.Save()

	assertEqual(t, String(r2), "日本本")
	assertEqual(t, String(r3), "本本")
}

func TestPadRightS(t *testing.T) {
	t.Parallel()

	assertEqual(t, String(PadRightS([]rune("abc"), 'x', 5)), "abcxx")
	assertEqual(t, String(PadRightS([]rune("abcdef"), 'x', 5)), "abcdef")
}
