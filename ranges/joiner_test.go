package ranges

import (
	"testing"
)

func TestJoiner(t *testing.T) {
	t.Parallel()

	result := string(Slice(
		Joiner(
			Cache(
				Map(
					I(Only("Mary", "has", "a", "little", "lamb")),
					func(x string) InputRange[rune] { return I(Runes(x)) },
				),
			),
			Runes("..."),
		),
	))

	assertEqual(t, result, "Mary...has...a...little...lamb")
}

func TestJoinerF(t *testing.T) {
	t.Parallel()

	result := string(SliceF(
		JoinerF(
			CacheF(
				MapF(
					Only("Mary", "has", "a", "little", "lamb"),
					func(x string) ForwardRange[byte] { return Bytes(x) },
				),
			),
			Bytes("..."),
		),
	))

	assertEqual(t, result, "Mary...has...a...little...lamb")
}

func TestJoinerS(t *testing.T) {
	t.Parallel()

	result := string(SliceF(
		JoinerS(
			[]ForwardRange[rune]{
				Runes("Mary"),
				Runes("has"),
				Runes("a"),
				Runes("little"),
				Runes("lamb"),
			},
			Runes("..."),
		),
	))

	assertEqual(t, result, "Mary...has...a...little...lamb")
}

func TestJoinerSS(t *testing.T) {
	t.Parallel()

	result := string(SliceF(
		JoinerSS(
			[][]rune{
				[]rune("Mary"),
				[]rune("has"),
				[]rune("a"),
				[]rune("little"),
				[]rune("lamb"),
			},
			[]rune("..."),
		),
	))

	assertEqual(t, result, "Mary...has...a...little...lamb")
}

func TestJoinStrings(t *testing.T) {
	t.Parallel()

	result := JoinStrings(
		Only("Mary", "has", "a", "little", "lamb"),
		"...",
	)

	assertEqual(t, result, "Mary...has...a...little...lamb")
}
