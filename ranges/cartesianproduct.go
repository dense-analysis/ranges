package ranges

type cartesianProduct2Result[A, B any] struct {
	r1     ForwardRange[A]
	r2     ForwardRange[B]
	r2Save ForwardRange[B]
}

func (cpr *cartesianProduct2Result[A, B]) Empty() bool {
	return cpr.r2Save.Empty() &&
		cpr.r1.Empty()
}

func (cpr *cartesianProduct2Result[A, B]) Front() Pair[A, B] {
	return MakePair(
		cpr.r1.Front(),
		cpr.r2Save.Front(),
	)
}

func (cpr *cartesianProduct2Result[A, B]) PopFront() {
	cpr.r2Save.PopFront()

	if cpr.r2Save.Empty() {
		cpr.r1.PopFront()

		if !cpr.r1.Empty() {
			cpr.r2Save = cpr.r2.Save()
		}
	}
}

func (cpr *cartesianProduct2Result[A, B]) Save() ForwardRange[Pair[A, B]] {
	return &cartesianProduct2Result[A, B]{
		cpr.r1.Save(),
		cpr.r2.Save(),
		cpr.r2Save.Save(),
	}
}

// CartesianProduct2 produces the cartersian product r1 X r2
func CartesianProduct2[A, B any](
	r1 ForwardRange[A],
	r2 ForwardRange[B],
) ForwardRange[Pair[A, B]] {
	return &cartesianProduct2Result[A, B]{
		r1,
		r2,
		r2.Save(),
	}
}

// CartesianProduct3 produces the cartersian product r1 X r2 X r3
func CartesianProduct3[A, B, C any](
	r1 ForwardRange[A],
	r2 ForwardRange[B],
	r3 ForwardRange[C],
) ForwardRange[Triplet[A, B, C]] {
	product2 := CartesianProduct2(r2, r3)
	product3 := CartesianProduct2(r1, product2)

	return MapF(
		product3,
		func(x Pair[A, Pair[B, C]]) Triplet[A, B, C] {
			return MakeTriplet(x.A, x.B.A, x.B.B)
		},
	)
}

// CartesianProduct4 produces the cartersian product r1 X r2 X r3 X r4
func CartesianProduct4[A, B, C, D any](
	r1 ForwardRange[A],
	r2 ForwardRange[B],
	r3 ForwardRange[C],
	r4 ForwardRange[D],
) ForwardRange[Quartet[A, B, C, D]] {
	product3 := CartesianProduct3(r2, r3, r4)
	product4 := CartesianProduct2(r1, product3)

	return MapF(
		product4,
		func(x Pair[A, Triplet[B, C, D]]) Quartet[A, B, C, D] {
			return MakeQuartet(x.A, x.B.A, x.B.B, x.B.C)
		},
	)
}

// CartesianProduct5 produces the cartersian product r1 X r2 X r3 X r4 X r5
func CartesianProduct5[A, B, C, D, E any](
	r1 ForwardRange[A],
	r2 ForwardRange[B],
	r3 ForwardRange[C],
	r4 ForwardRange[D],
	r5 ForwardRange[E],
) ForwardRange[Quintet[A, B, C, D, E]] {
	product4 := CartesianProduct4(r2, r3, r4, r5)
	product5 := CartesianProduct2(r1, product4)

	return MapF(
		product5,
		func(x Pair[A, Quartet[B, C, D, E]]) Quintet[A, B, C, D, E] {
			return MakeQuintet(x.A, x.B.A, x.B.B, x.B.C, x.B.D)
		},
	)
}

// CartesianProduct6 produces the cartersian product r1 X r2 X r3 X r4 X r5 X r6
func CartesianProduct6[A, B, C, D, E, F any](
	r1 ForwardRange[A],
	r2 ForwardRange[B],
	r3 ForwardRange[C],
	r4 ForwardRange[D],
	r5 ForwardRange[E],
	r6 ForwardRange[F],
) ForwardRange[Sextet[A, B, C, D, E, F]] {
	product5 := CartesianProduct5(r2, r3, r4, r5, r6)
	product6 := CartesianProduct2(r1, product5)

	return MapF(
		product6,
		func(x Pair[A, Quintet[B, C, D, E, F]]) Sextet[A, B, C, D, E, F] {
			return MakeSextet(x.A, x.B.A, x.B.B, x.B.C, x.B.D, x.B.E)
		},
	)
}

// CartesianProduct7 produces the cartersian product r1 X r2 X r3 X r4 X r5 X r6 x r7
func CartesianProduct7[A, B, C, D, E, F, G any](
	r1 ForwardRange[A],
	r2 ForwardRange[B],
	r3 ForwardRange[C],
	r4 ForwardRange[D],
	r5 ForwardRange[E],
	r6 ForwardRange[F],
	r7 ForwardRange[G],
) ForwardRange[Septet[A, B, C, D, E, F, G]] {
	product6 := CartesianProduct6(r2, r3, r4, r5, r6, r7)
	product7 := CartesianProduct2(r1, product6)

	return MapF(
		product7,
		func(x Pair[A, Sextet[B, C, D, E, F, G]]) Septet[A, B, C, D, E, F, G] {
			return MakeSeptet(x.A, x.B.A, x.B.B, x.B.C, x.B.D, x.B.E, x.B.F)
		},
	)
}

// CartesianProduct8 produces the cartersian product r1 X r2 X r3 X r4 X r5 X r6 x r7 x r8
func CartesianProduct8[A, B, C, D, E, F, G, H any](
	r1 ForwardRange[A],
	r2 ForwardRange[B],
	r3 ForwardRange[C],
	r4 ForwardRange[D],
	r5 ForwardRange[E],
	r6 ForwardRange[F],
	r7 ForwardRange[G],
	r8 ForwardRange[H],
) ForwardRange[Octet[A, B, C, D, E, F, G, H]] {
	product7 := CartesianProduct7(r2, r3, r4, r5, r6, r7, r8)
	product8 := CartesianProduct2(r1, product7)

	return MapF(
		product8,
		func(x Pair[A, Septet[B, C, D, E, F, G, H]]) Octet[A, B, C, D, E, F, G, H] {
			return MakeOctet(x.A, x.B.A, x.B.B, x.B.C, x.B.D, x.B.E, x.B.F, x.B.G)
		},
	)
}

// CartesianProduct9 produces the cartersian product r1 X r2 X r3 X r4 X r5 X r6 x r7 x r8 x r9
func CartesianProduct9[A, B, C, D, E, F, G, H, I any](
	r1 ForwardRange[A],
	r2 ForwardRange[B],
	r3 ForwardRange[C],
	r4 ForwardRange[D],
	r5 ForwardRange[E],
	r6 ForwardRange[F],
	r7 ForwardRange[G],
	r8 ForwardRange[H],
	r9 ForwardRange[I],
) ForwardRange[Ennead[A, B, C, D, E, F, G, H, I]] {
	product8 := CartesianProduct8(r2, r3, r4, r5, r6, r7, r8, r9)
	product9 := CartesianProduct2(r1, product8)

	return MapF(
		product9,
		func(x Pair[A, Octet[B, C, D, E, F, G, H, I]]) Ennead[A, B, C, D, E, F, G, H, I] {
			return MakeEnnead(x.A, x.B.A, x.B.B, x.B.C, x.B.D, x.B.E, x.B.F, x.B.G, x.B.H)
		},
	)
}

// CartesianProduct10 produces the cartersian product r1 X r2 X r3 X r4 X r5 X r6 x r7 x r8 x r9 x r10
func CartesianProduct10[A, B, C, D, E, F, G, H, I, J any](
	r1 ForwardRange[A],
	r2 ForwardRange[B],
	r3 ForwardRange[C],
	r4 ForwardRange[D],
	r5 ForwardRange[E],
	r6 ForwardRange[F],
	r7 ForwardRange[G],
	r8 ForwardRange[H],
	r9 ForwardRange[I],
	r10 ForwardRange[J],
) ForwardRange[Decade[A, B, C, D, E, F, G, H, I, J]] {
	product9 := CartesianProduct9(r2, r3, r4, r5, r6, r7, r8, r9, r10)
	product10 := CartesianProduct2(r1, product9)

	return MapF(
		product10,
		func(x Pair[A, Ennead[B, C, D, E, F, G, H, I, J]]) Decade[A, B, C, D, E, F, G, H, I, J] {
			return MakeDecade(x.A, x.B.A, x.B.B, x.B.C, x.B.D, x.B.E, x.B.F, x.B.G, x.B.H, x.B.I)
		},
	)
}
