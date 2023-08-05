package ranges

// Generated with: go run internal/cmd/generatecode/main.go

// Pair holds 2 items
type Pair[A, B any] struct {
	A A
	B B
}

// Get returns items as a Go tuple
func (p Pair[A, B]) Get() (A, B) {
	return p.A, p.B
}

// MakePair creates a Pair
func MakePair[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{a, b}
}

// Triplet holds 3 items
type Triplet[A, B, C any] struct {
	A A
	B B
	C C
}

// Get returns items as a Go tuple
func (t Triplet[A, B, C]) Get() (A, B, C) {
	return t.A, t.B, t.C
}

// MakeTriplet creates a Triplet
func MakeTriplet[A, B, C any](a A, b B, c C) Triplet[A, B, C] {
	return Triplet[A, B, C]{a, b, c}
}

// Quartet holds 4 items
type Quartet[A, B, C, D any] struct {
	A A
	B B
	C C
	D D
}

// Get returns items as a Go tuple
func (q Quartet[A, B, C, D]) Get() (A, B, C, D) {
	return q.A, q.B, q.C, q.D
}

// MakeQuartet creates a Quartet
func MakeQuartet[A, B, C, D any](a A, b B, c C, d D) Quartet[A, B, C, D] {
	return Quartet[A, B, C, D]{a, b, c, d}
}

// Quintet holds 5 items
type Quintet[A, B, C, D, E any] struct {
	A A
	B B
	C C
	D D
	E E
}

// Get returns items as a Go tuple
func (q Quintet[A, B, C, D, E]) Get() (A, B, C, D, E) {
	return q.A, q.B, q.C, q.D, q.E
}

// MakeQuintet creates a Quintet
func MakeQuintet[A, B, C, D, E any](a A, b B, c C, d D, e E) Quintet[A, B, C, D, E] {
	return Quintet[A, B, C, D, E]{a, b, c, d, e}
}

// Sextet holds 6 items
type Sextet[A, B, C, D, E, F any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
}

// Get returns items as a Go tuple
func (s Sextet[A, B, C, D, E, F]) Get() (A, B, C, D, E, F) {
	return s.A, s.B, s.C, s.D, s.E, s.F
}

// MakeSextet creates a Sextet
func MakeSextet[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) Sextet[A, B, C, D, E, F] {
	return Sextet[A, B, C, D, E, F]{a, b, c, d, e, f}
}

// Septet holds 7 items
type Septet[A, B, C, D, E, F, G any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
}

// Get returns items as a Go tuple
func (s Septet[A, B, C, D, E, F, G]) Get() (A, B, C, D, E, F, G) {
	return s.A, s.B, s.C, s.D, s.E, s.F, s.G
}

// MakeSeptet creates a Septet
func MakeSeptet[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) Septet[A, B, C, D, E, F, G] {
	return Septet[A, B, C, D, E, F, G]{a, b, c, d, e, f, g}
}

// Octet holds 8 items
type Octet[A, B, C, D, E, F, G, H any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
}

// Get returns items as a Go tuple
func (o Octet[A, B, C, D, E, F, G, H]) Get() (A, B, C, D, E, F, G, H) {
	return o.A, o.B, o.C, o.D, o.E, o.F, o.G, o.H
}

// MakeOctet creates a Octet
func MakeOctet[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) Octet[A, B, C, D, E, F, G, H] {
	return Octet[A, B, C, D, E, F, G, H]{a, b, c, d, e, f, g, h}
}

// Ennead holds 9 items
type Ennead[A, B, C, D, E, F, G, H, I any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
}

// Get returns items as a Go tuple
func (e Ennead[A, B, C, D, E, F, G, H, I]) Get() (A, B, C, D, E, F, G, H, I) {
	return e.A, e.B, e.C, e.D, e.E, e.F, e.G, e.H, e.I
}

// MakeEnnead creates a Ennead
func MakeEnnead[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) Ennead[A, B, C, D, E, F, G, H, I] {
	return Ennead[A, B, C, D, E, F, G, H, I]{a, b, c, d, e, f, g, h, i}
}

// Decade holds 10 items
type Decade[A, B, C, D, E, F, G, H, I, J any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
	J J
}

// Get returns items as a Go tuple
func (d Decade[A, B, C, D, E, F, G, H, I, J]) Get() (A, B, C, D, E, F, G, H, I, J) {
	return d.A, d.B, d.C, d.D, d.E, d.F, d.G, d.H, d.I, d.J
}

// MakeDecade creates a Decade
func MakeDecade[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) Decade[A, B, C, D, E, F, G, H, I, J] {
	return Decade[A, B, C, D, E, F, G, H, I, J]{a, b, c, d, e, f, g, h, i, j}
}
