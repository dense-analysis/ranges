package ranges

// Mismatch2 eagerly advances all ranges in lockstep until `eq` returns `false`.
func Mismatch2[T any](
	eq func(a, b T) bool,
	a InputRange[T],
	b InputRange[T],
) InputRange[Pair[T, T]] {
	r := Zip2(a, b)
	Exhaust(
		TakeWhile(
			r,
			func(item Pair[T, T]) bool {
				return eq(item.A, item.B)
			},
		),
	)

	return r
}

// Mismatch3 eagerly advances all ranges in lockstep until `eq` returns `false`.
func Mismatch3[T any](
	eq func(a, b T) bool,
	a InputRange[T],
	b InputRange[T],
	c InputRange[T],
) InputRange[Triplet[T, T, T]] {
	r := Zip3(a, b, c)
	Exhaust(
		TakeWhile(
			Zip3(a, b, c),
			func(item Triplet[T, T, T]) bool {
				return eq(item.A, item.B) &&
					eq(item.B, item.C)
			},
		),
	)

	return r
}

// Mismatch4 eagerly advances all ranges in lockstep until `eq` returns `false`.
func Mismatch4[T any](
	eq func(a, b T) bool,
	a InputRange[T],
	b InputRange[T],
	c InputRange[T],
	d InputRange[T],
) InputRange[Quartet[T, T, T, T]] {
	r := Zip4(a, b, c, d)
	Exhaust(
		TakeWhile(
			r,
			func(item Quartet[T, T, T, T]) bool {
				return eq(item.A, item.B) &&
					eq(item.B, item.C) &&
					eq(item.C, item.D)
			},
		),
	)

	return r
}

// Mismatch5 eagerly advances all ranges in lockstep until `eq` returns `false`.
func Mismatch5[T any](
	eq func(a, b T) bool,
	a InputRange[T],
	b InputRange[T],
	c InputRange[T],
	d InputRange[T],
	e InputRange[T],
) InputRange[Quintet[T, T, T, T, T]] {
	r := Zip5(a, b, c, d, e)
	Exhaust(
		TakeWhile(
			r,
			func(item Quintet[T, T, T, T, T]) bool {
				return eq(item.A, item.B) &&
					eq(item.B, item.C) &&
					eq(item.C, item.D) &&
					eq(item.D, item.E)
			},
		),
	)

	return r
}

// Mismatch6 eagerly advances all ranges in lockstep until `eq` returns `false`.
func Mismatch6[T any](
	eq func(a, b T) bool,
	a InputRange[T],
	b InputRange[T],
	c InputRange[T],
	d InputRange[T],
	e InputRange[T],
	f InputRange[T],
) InputRange[Sextet[T, T, T, T, T, T]] {
	r := Zip6(a, b, c, d, e, f)
	Exhaust(
		TakeWhile(
			r,
			func(item Sextet[T, T, T, T, T, T]) bool {
				return eq(item.A, item.B) &&
					eq(item.B, item.C) &&
					eq(item.C, item.D) &&
					eq(item.D, item.E) &&
					eq(item.E, item.F)
			},
		),
	)

	return r
}

// Mismatch7 eagerly advances all ranges in lockstep until `eq` returns `false`.
func Mismatch7[T any](
	eq func(a, b T) bool,
	a InputRange[T],
	b InputRange[T],
	c InputRange[T],
	d InputRange[T],
	e InputRange[T],
	f InputRange[T],
	g InputRange[T],
) InputRange[Septet[T, T, T, T, T, T, T]] {
	r := Zip7(a, b, c, d, e, f, g)
	Exhaust(
		TakeWhile(
			r,
			func(item Septet[T, T, T, T, T, T, T]) bool {
				return eq(item.A, item.B) &&
					eq(item.B, item.C) &&
					eq(item.C, item.D) &&
					eq(item.D, item.E) &&
					eq(item.E, item.F) &&
					eq(item.F, item.G)
			},
		),
	)

	return r
}

// Mismatch8 eagerly advances all ranges in lockstep until `eq` returns `false`.
func Mismatch8[T any](
	eq func(a, b T) bool,
	a InputRange[T],
	b InputRange[T],
	c InputRange[T],
	d InputRange[T],
	e InputRange[T],
	f InputRange[T],
	g InputRange[T],
	h InputRange[T],
) InputRange[Octet[T, T, T, T, T, T, T, T]] {
	r := Zip8(a, b, c, d, e, f, g, h)
	Exhaust(
		TakeWhile(
			r,
			func(item Octet[T, T, T, T, T, T, T, T]) bool {
				return eq(item.A, item.B) &&
					eq(item.B, item.C) &&
					eq(item.C, item.D) &&
					eq(item.D, item.E) &&
					eq(item.E, item.F) &&
					eq(item.F, item.G) &&
					eq(item.G, item.H)
			},
		),
	)

	return r
}

// Mismatch9 eagerly advances all ranges in lockstep until `eq` returns `false`.
func Mismatch9[T any](
	eq func(a, b T) bool,
	a InputRange[T],
	b InputRange[T],
	c InputRange[T],
	d InputRange[T],
	e InputRange[T],
	f InputRange[T],
	g InputRange[T],
	h InputRange[T],
	i InputRange[T],
) InputRange[Ennead[T, T, T, T, T, T, T, T, T]] {
	r := Zip9(a, b, c, d, e, f, g, h, i)
	Exhaust(
		TakeWhile(
			r,
			func(item Ennead[T, T, T, T, T, T, T, T, T]) bool {
				return eq(item.A, item.B) &&
					eq(item.B, item.C) &&
					eq(item.C, item.D) &&
					eq(item.D, item.E) &&
					eq(item.E, item.F) &&
					eq(item.F, item.G) &&
					eq(item.G, item.H) &&
					eq(item.H, item.I)
			},
		),
	)

	return r
}

// Mismatch10 eagerly advances all ranges in lockstep until `eq` returns `false`.
func Mismatch10[T any](
	eq func(a, b T) bool,
	a InputRange[T],
	b InputRange[T],
	c InputRange[T],
	d InputRange[T],
	e InputRange[T],
	f InputRange[T],
	g InputRange[T],
	h InputRange[T],
	i InputRange[T],
	j InputRange[T],
) InputRange[Decade[T, T, T, T, T, T, T, T, T, T]] {
	r := Zip10(a, b, c, d, e, f, g, h, i, j)
	Exhaust(
		TakeWhile(
			r,
			func(item Decade[T, T, T, T, T, T, T, T, T, T]) bool {
				return eq(item.A, item.B) &&
					eq(item.B, item.C) &&
					eq(item.C, item.D) &&
					eq(item.D, item.E) &&
					eq(item.E, item.F) &&
					eq(item.F, item.G) &&
					eq(item.G, item.H) &&
					eq(item.H, item.I) &&
					eq(item.I, item.J)
			},
		),
	)

	return r
}
