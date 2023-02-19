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
					I(F(Only("Mary", "has", "a", "little", "lamb"))),
					func(x string) InputRange[rune] { return I(F(Runes(x))) },
				),
			),
			F(Runes("...")),
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
					F(Only("Mary", "has", "a", "little", "lamb")),
					func(x string) ForwardRange[byte] { return Bytes(x) },
				),
			),
			F(Bytes("...")),
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
			F(Runes("...")),
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
