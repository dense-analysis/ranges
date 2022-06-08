package ranges

// Generated with: go run internal/cmd/generatecode/main.go

type zipPairResult[A, B any] struct {
	a InputRange[A]
	b InputRange[B]
}

func (z *zipPairResult[A, B]) Empty() bool {
	return z.a.Empty() ||
		z.b.Empty()
}

func (z *zipPairResult[A, B]) Front() Pair[A, B] {
	return Pair[A, B]{
		z.a.Front(),
		z.b.Front(),
	}
}

func (z *zipPairResult[A, B]) PopFront() {
	z.a.PopFront()
	z.b.PopFront()
}

// Zip2 produces items from 2 ranges in parallel.
// The range will be empty when any of the ranges are empty.
func Zip2[A, B any](
	a InputRange[A],
	b InputRange[B],
) InputRange[Pair[A, B]] {
	return &zipPairResult[A, B]{a, b}
}

type zipTripletResult[A, B, C any] struct {
	a InputRange[A]
	b InputRange[B]
	c InputRange[C]
}

func (z *zipTripletResult[A, B, C]) Empty() bool {
	return z.a.Empty() ||
		z.b.Empty() ||
		z.c.Empty()
}

func (z *zipTripletResult[A, B, C]) Front() Triplet[A, B, C] {
	return Triplet[A, B, C]{
		z.a.Front(),
		z.b.Front(),
		z.c.Front(),
	}
}

func (z *zipTripletResult[A, B, C]) PopFront() {
	z.a.PopFront()
	z.b.PopFront()
	z.c.PopFront()
}

// Zip3 produces items from 3 ranges in parallel.
// The range will be empty when any of the ranges are empty.
func Zip3[A, B, C any](
	a InputRange[A],
	b InputRange[B],
	c InputRange[C],
) InputRange[Triplet[A, B, C]] {
	return &zipTripletResult[A, B, C]{a, b, c}
}

type zipQuartetResult[A, B, C, D any] struct {
	a InputRange[A]
	b InputRange[B]
	c InputRange[C]
	d InputRange[D]
}

func (z *zipQuartetResult[A, B, C, D]) Empty() bool {
	return z.a.Empty() ||
		z.b.Empty() ||
		z.c.Empty() ||
		z.d.Empty()
}

func (z *zipQuartetResult[A, B, C, D]) Front() Quartet[A, B, C, D] {
	return Quartet[A, B, C, D]{
		z.a.Front(),
		z.b.Front(),
		z.c.Front(),
		z.d.Front(),
	}
}

func (z *zipQuartetResult[A, B, C, D]) PopFront() {
	z.a.PopFront()
	z.b.PopFront()
	z.c.PopFront()
	z.d.PopFront()
}

// Zip4 produces items from 4 ranges in parallel.
// The range will be empty when any of the ranges are empty.
func Zip4[A, B, C, D any](
	a InputRange[A],
	b InputRange[B],
	c InputRange[C],
	d InputRange[D],
) InputRange[Quartet[A, B, C, D]] {
	return &zipQuartetResult[A, B, C, D]{a, b, c, d}
}

type zipQuintetResult[A, B, C, D, E any] struct {
	a InputRange[A]
	b InputRange[B]
	c InputRange[C]
	d InputRange[D]
	e InputRange[E]
}

func (z *zipQuintetResult[A, B, C, D, E]) Empty() bool {
	return z.a.Empty() ||
		z.b.Empty() ||
		z.c.Empty() ||
		z.d.Empty() ||
		z.e.Empty()
}

func (z *zipQuintetResult[A, B, C, D, E]) Front() Quintet[A, B, C, D, E] {
	return Quintet[A, B, C, D, E]{
		z.a.Front(),
		z.b.Front(),
		z.c.Front(),
		z.d.Front(),
		z.e.Front(),
	}
}

func (z *zipQuintetResult[A, B, C, D, E]) PopFront() {
	z.a.PopFront()
	z.b.PopFront()
	z.c.PopFront()
	z.d.PopFront()
	z.e.PopFront()
}

// Zip5 produces items from 5 ranges in parallel.
// The range will be empty when any of the ranges are empty.
func Zip5[A, B, C, D, E any](
	a InputRange[A],
	b InputRange[B],
	c InputRange[C],
	d InputRange[D],
	e InputRange[E],
) InputRange[Quintet[A, B, C, D, E]] {
	return &zipQuintetResult[A, B, C, D, E]{a, b, c, d, e}
}

type zipSextetResult[A, B, C, D, E, F any] struct {
	a InputRange[A]
	b InputRange[B]
	c InputRange[C]
	d InputRange[D]
	e InputRange[E]
	f InputRange[F]
}

func (z *zipSextetResult[A, B, C, D, E, F]) Empty() bool {
	return z.a.Empty() ||
		z.b.Empty() ||
		z.c.Empty() ||
		z.d.Empty() ||
		z.e.Empty() ||
		z.f.Empty()
}

func (z *zipSextetResult[A, B, C, D, E, F]) Front() Sextet[A, B, C, D, E, F] {
	return Sextet[A, B, C, D, E, F]{
		z.a.Front(),
		z.b.Front(),
		z.c.Front(),
		z.d.Front(),
		z.e.Front(),
		z.f.Front(),
	}
}

func (z *zipSextetResult[A, B, C, D, E, F]) PopFront() {
	z.a.PopFront()
	z.b.PopFront()
	z.c.PopFront()
	z.d.PopFront()
	z.e.PopFront()
	z.f.PopFront()
}

// Zip6 produces items from 6 ranges in parallel.
// The range will be empty when any of the ranges are empty.
func Zip6[A, B, C, D, E, F any](
	a InputRange[A],
	b InputRange[B],
	c InputRange[C],
	d InputRange[D],
	e InputRange[E],
	f InputRange[F],
) InputRange[Sextet[A, B, C, D, E, F]] {
	return &zipSextetResult[A, B, C, D, E, F]{a, b, c, d, e, f}
}

type zipSeptetResult[A, B, C, D, E, F, G any] struct {
	a InputRange[A]
	b InputRange[B]
	c InputRange[C]
	d InputRange[D]
	e InputRange[E]
	f InputRange[F]
	g InputRange[G]
}

func (z *zipSeptetResult[A, B, C, D, E, F, G]) Empty() bool {
	return z.a.Empty() ||
		z.b.Empty() ||
		z.c.Empty() ||
		z.d.Empty() ||
		z.e.Empty() ||
		z.f.Empty() ||
		z.g.Empty()
}

func (z *zipSeptetResult[A, B, C, D, E, F, G]) Front() Septet[A, B, C, D, E, F, G] {
	return Septet[A, B, C, D, E, F, G]{
		z.a.Front(),
		z.b.Front(),
		z.c.Front(),
		z.d.Front(),
		z.e.Front(),
		z.f.Front(),
		z.g.Front(),
	}
}

func (z *zipSeptetResult[A, B, C, D, E, F, G]) PopFront() {
	z.a.PopFront()
	z.b.PopFront()
	z.c.PopFront()
	z.d.PopFront()
	z.e.PopFront()
	z.f.PopFront()
	z.g.PopFront()
}

// Zip7 produces items from 7 ranges in parallel.
// The range will be empty when any of the ranges are empty.
func Zip7[A, B, C, D, E, F, G any](
	a InputRange[A],
	b InputRange[B],
	c InputRange[C],
	d InputRange[D],
	e InputRange[E],
	f InputRange[F],
	g InputRange[G],
) InputRange[Septet[A, B, C, D, E, F, G]] {
	return &zipSeptetResult[A, B, C, D, E, F, G]{a, b, c, d, e, f, g}
}

type zipOctetResult[A, B, C, D, E, F, G, H any] struct {
	a InputRange[A]
	b InputRange[B]
	c InputRange[C]
	d InputRange[D]
	e InputRange[E]
	f InputRange[F]
	g InputRange[G]
	h InputRange[H]
}

func (z *zipOctetResult[A, B, C, D, E, F, G, H]) Empty() bool {
	return z.a.Empty() ||
		z.b.Empty() ||
		z.c.Empty() ||
		z.d.Empty() ||
		z.e.Empty() ||
		z.f.Empty() ||
		z.g.Empty() ||
		z.h.Empty()
}

func (z *zipOctetResult[A, B, C, D, E, F, G, H]) Front() Octet[A, B, C, D, E, F, G, H] {
	return Octet[A, B, C, D, E, F, G, H]{
		z.a.Front(),
		z.b.Front(),
		z.c.Front(),
		z.d.Front(),
		z.e.Front(),
		z.f.Front(),
		z.g.Front(),
		z.h.Front(),
	}
}

func (z *zipOctetResult[A, B, C, D, E, F, G, H]) PopFront() {
	z.a.PopFront()
	z.b.PopFront()
	z.c.PopFront()
	z.d.PopFront()
	z.e.PopFront()
	z.f.PopFront()
	z.g.PopFront()
	z.h.PopFront()
}

// Zip8 produces items from 8 ranges in parallel.
// The range will be empty when any of the ranges are empty.
func Zip8[A, B, C, D, E, F, G, H any](
	a InputRange[A],
	b InputRange[B],
	c InputRange[C],
	d InputRange[D],
	e InputRange[E],
	f InputRange[F],
	g InputRange[G],
	h InputRange[H],
) InputRange[Octet[A, B, C, D, E, F, G, H]] {
	return &zipOctetResult[A, B, C, D, E, F, G, H]{a, b, c, d, e, f, g, h}
}

type zipEnneadResult[A, B, C, D, E, F, G, H, I any] struct {
	a InputRange[A]
	b InputRange[B]
	c InputRange[C]
	d InputRange[D]
	e InputRange[E]
	f InputRange[F]
	g InputRange[G]
	h InputRange[H]
	i InputRange[I]
}

func (z *zipEnneadResult[A, B, C, D, E, F, G, H, I]) Empty() bool {
	return z.a.Empty() ||
		z.b.Empty() ||
		z.c.Empty() ||
		z.d.Empty() ||
		z.e.Empty() ||
		z.f.Empty() ||
		z.g.Empty() ||
		z.h.Empty() ||
		z.i.Empty()
}

func (z *zipEnneadResult[A, B, C, D, E, F, G, H, I]) Front() Ennead[A, B, C, D, E, F, G, H, I] {
	return Ennead[A, B, C, D, E, F, G, H, I]{
		z.a.Front(),
		z.b.Front(),
		z.c.Front(),
		z.d.Front(),
		z.e.Front(),
		z.f.Front(),
		z.g.Front(),
		z.h.Front(),
		z.i.Front(),
	}
}

func (z *zipEnneadResult[A, B, C, D, E, F, G, H, I]) PopFront() {
	z.a.PopFront()
	z.b.PopFront()
	z.c.PopFront()
	z.d.PopFront()
	z.e.PopFront()
	z.f.PopFront()
	z.g.PopFront()
	z.h.PopFront()
	z.i.PopFront()
}

// Zip9 produces items from 9 ranges in parallel.
// The range will be empty when any of the ranges are empty.
func Zip9[A, B, C, D, E, F, G, H, I any](
	a InputRange[A],
	b InputRange[B],
	c InputRange[C],
	d InputRange[D],
	e InputRange[E],
	f InputRange[F],
	g InputRange[G],
	h InputRange[H],
	i InputRange[I],
) InputRange[Ennead[A, B, C, D, E, F, G, H, I]] {
	return &zipEnneadResult[A, B, C, D, E, F, G, H, I]{a, b, c, d, e, f, g, h, i}
}

type zipDecadeResult[A, B, C, D, E, F, G, H, I, J any] struct {
	a InputRange[A]
	b InputRange[B]
	c InputRange[C]
	d InputRange[D]
	e InputRange[E]
	f InputRange[F]
	g InputRange[G]
	h InputRange[H]
	i InputRange[I]
	j InputRange[J]
}

func (z *zipDecadeResult[A, B, C, D, E, F, G, H, I, J]) Empty() bool {
	return z.a.Empty() ||
		z.b.Empty() ||
		z.c.Empty() ||
		z.d.Empty() ||
		z.e.Empty() ||
		z.f.Empty() ||
		z.g.Empty() ||
		z.h.Empty() ||
		z.i.Empty() ||
		z.j.Empty()
}

func (z *zipDecadeResult[A, B, C, D, E, F, G, H, I, J]) Front() Decade[A, B, C, D, E, F, G, H, I, J] {
	return Decade[A, B, C, D, E, F, G, H, I, J]{
		z.a.Front(),
		z.b.Front(),
		z.c.Front(),
		z.d.Front(),
		z.e.Front(),
		z.f.Front(),
		z.g.Front(),
		z.h.Front(),
		z.i.Front(),
		z.j.Front(),
	}
}

func (z *zipDecadeResult[A, B, C, D, E, F, G, H, I, J]) PopFront() {
	z.a.PopFront()
	z.b.PopFront()
	z.c.PopFront()
	z.d.PopFront()
	z.e.PopFront()
	z.f.PopFront()
	z.g.PopFront()
	z.h.PopFront()
	z.i.PopFront()
	z.j.PopFront()
}

// Zip10 produces items from 10 ranges in parallel.
// The range will be empty when any of the ranges are empty.
func Zip10[A, B, C, D, E, F, G, H, I, J any](
	a InputRange[A],
	b InputRange[B],
	c InputRange[C],
	d InputRange[D],
	e InputRange[E],
	f InputRange[F],
	g InputRange[G],
	h InputRange[H],
	i InputRange[I],
	j InputRange[J],
) InputRange[Decade[A, B, C, D, E, F, G, H, I, J]] {
	return &zipDecadeResult[A, B, C, D, E, F, G, H, I, J]{a, b, c, d, e, f, g, h, i, j}
}
