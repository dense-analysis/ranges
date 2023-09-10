package ranges

type baseChain[T any] struct {
	r        InputRange[InputRange[T]]
	isPrimed bool
}

func (bc *baseChain[T]) prime() {
	if !bc.isPrimed {
		for !bc.r.Empty() && bc.r.Front().Empty() {
			bc.r.PopFront()
		}

		bc.isPrimed = true
	}
}

func (bc *baseChain[T]) Empty() bool {
	bc.prime()

	return bc.r.Empty()
}

func (bc *baseChain[T]) Front() T {
	bc.prime()

	return bc.r.Front().Front()
}

// flattenResult implments Flatten
type flattenResult[T any] struct {
	baseChain[T]
}

func (fr *flattenResult[T]) PopFront() {
	fr.prime()

	fr.r.Front().PopFront()
	fr.isPrimed = false
}

// frontTransversalResult implments FrontTransversal
type frontTransversalResult[T any] struct {
	baseChain[T]
}

func (ftr *frontTransversalResult[T]) PopFront() {
	ftr.prime()
	ftr.r.PopFront()
	ftr.isPrimed = false
}

type baseForwardChain[T any] struct {
	r        ForwardRange[ForwardRange[T]]
	isPrimed bool
}

func (bc *baseForwardChain[T]) prime() {
	if !bc.isPrimed {
		for !bc.r.Empty() && bc.r.Front().Empty() {
			bc.r.PopFront()
		}

		bc.isPrimed = true
	}
}

func (bc *baseForwardChain[T]) Empty() bool {
	bc.prime()

	return bc.r.Empty()
}

func (bc *baseForwardChain[T]) Front() T {
	bc.prime()

	return bc.r.Front().Front()
}

func (bc *baseForwardChain[T]) saveOuter() ForwardRange[ForwardRange[T]] {
	newList := make([]ForwardRange[T], 0)
	r := bc.r.Save()

	for !r.Empty() {
		newList = append(newList, r.Front().Save())
		r.PopFront()
	}

	return SliceRange(newList)
}

// flattenForwardResult implements FlattenF
type flattenForwardResult[T any] struct {
	baseForwardChain[T]
}

func (fr *flattenForwardResult[T]) PopFront() {
	fr.prime()

	fr.r.Front().PopFront()
	fr.isPrimed = false
}

func (fr *flattenForwardResult[T]) Save() ForwardRange[T] {
	return &flattenForwardResult[T]{baseForwardChain[T]{fr.saveOuter(), fr.isPrimed}}
}

// frontTransversalForwardResult implments FrontTraveralF
type frontTransversalForwardResult[T any] struct {
	baseForwardChain[T]
}

func (ftr *frontTransversalForwardResult[T]) PopFront() {
	ftr.prime()
	ftr.r.PopFront()
	ftr.isPrimed = false
}

func (ftr *frontTransversalForwardResult[T]) Save() ForwardRange[T] {
	return &frontTransversalForwardResult[T]{baseForwardChain[T]{ftr.saveOuter(), ftr.isPrimed}}
}

type baseBidrectionalChain[T any] struct {
	r        BidirectionalRange[BidirectionalRange[T]]
	isPrimed bool
}

func (bc *baseBidrectionalChain[T]) prime() {
	if !bc.isPrimed {
		for !bc.r.Empty() && bc.r.Front().Empty() {
			bc.r.PopFront()
		}

		for !bc.r.Empty() && bc.r.Back().Empty() {
			bc.r.PopBack()
		}

		bc.isPrimed = true
	}
}

func (bc *baseBidrectionalChain[T]) Empty() bool {
	bc.prime()

	return bc.r.Empty()
}

func (bc *baseBidrectionalChain[T]) Front() T {
	bc.prime()

	return bc.r.Front().Front()
}

func (bc *baseBidrectionalChain[T]) saveOuter() BidirectionalRange[BidirectionalRange[T]] {
	newList := make([]BidirectionalRange[T], 0)
	r := bc.r.Save()

	for !r.Empty() {
		newList = append(newList, r.Front().SaveB())
		r.PopFront()
	}

	return SliceRange(newList)
}

// flattenBidirectionalResult implements FlattenB
type flattenBidirectionalResult[T any] struct {
	baseBidrectionalChain[T]
}

func (fr *flattenBidirectionalResult[T]) PopFront() {
	fr.prime()

	fr.r.Front().PopFront()
	fr.isPrimed = false
}

func (fr *flattenBidirectionalResult[T]) Back() T {
	fr.prime()

	return fr.r.Back().Back()
}

func (fr *flattenBidirectionalResult[T]) PopBack() {
	fr.prime()

	fr.r.Back().PopBack()
	fr.isPrimed = false
}

func (fr *flattenBidirectionalResult[T]) Save() ForwardRange[T] {
	return fr.SaveB()
}

func (fr *flattenBidirectionalResult[T]) SaveB() BidirectionalRange[T] {
	return &flattenBidirectionalResult[T]{baseBidrectionalChain[T]{fr.saveOuter(), fr.isPrimed}}
}

// flattenRandomAccessResult implements FlattenR
type flattenRandomAccessResult[T any] struct {
	flattenBidirectionalResult[T]
}

// frontTransversalBidirectionalResult implments FrontTraveralB
type frontTransversalBidirectionalResult[T any] struct {
	baseBidrectionalChain[T]
}

func (fr *frontTransversalBidirectionalResult[T]) PopFront() {
	fr.prime()
	fr.r.PopFront()
	fr.isPrimed = false
}

func (fr *frontTransversalBidirectionalResult[T]) Back() T {
	fr.prime()

	return fr.r.Back().Front()
}

func (fr *frontTransversalBidirectionalResult[T]) PopBack() {
	fr.prime()
	fr.r.PopBack()
	fr.isPrimed = false
}

func (fr *frontTransversalBidirectionalResult[T]) Save() ForwardRange[T] {
	return fr.SaveB()
}

func (fr *frontTransversalBidirectionalResult[T]) SaveB() BidirectionalRange[T] {
	return &frontTransversalBidirectionalResult[T]{baseBidrectionalChain[T]{fr.saveOuter(), fr.isPrimed}}
}

// Flatten combines a range of ranges into one straight range
func Flatten[T any](r InputRange[InputRange[T]]) InputRange[T] {
	return &flattenResult[T]{baseChain[T]{r, false}}
}

// FlattenF is `Flatten` where the range can be saved.
func FlattenF[T any](r ForwardRange[ForwardRange[T]]) ForwardRange[T] {
	return &flattenForwardResult[T]{baseForwardChain[T]{r, false}}
}

// FlattenB is `FlattenF` that can be shrunk from the back.
func FlattenB[T any](r BidirectionalRange[BidirectionalRange[T]]) BidirectionalRange[T] {
	return &flattenBidirectionalResult[T]{baseBidrectionalChain[T]{r, false}}
}

// FlattenS is `FlattenF` accepting a slice.
func FlattenS[T any](r []ForwardRange[T]) ForwardRange[T] {
	return FlattenF(SliceRange(r))
}

// FlattenSB is `FlattenS` for bidirectional ranges.
func FlattenSB[T any](r []BidirectionalRange[T]) BidirectionalRange[T] {
	return FlattenB(SliceRange(r))
}

// FlattenSS is `FlattenB` accepting a slice of slices.
func FlattenSS[T any](r [][]T) BidirectionalRange[T] {
	return FlattenB(MapS(r, func(s []T) BidirectionalRange[T] { return SliceRange(s) }))
}

// FrontTransversal yields the first value in each range, skipping empty ranges.
func FrontTransversal[T any](r InputRange[InputRange[T]]) InputRange[T] {
	return &frontTransversalResult[T]{baseChain[T]{r, false}}
}

// FrontTransversalF is `FrontTransversal` where the range can be saved.
func FrontTransversalF[T any](r ForwardRange[ForwardRange[T]]) ForwardRange[T] {
	return &frontTransversalForwardResult[T]{baseForwardChain[T]{r, false}}
}

// FrontTransversalB is `FrontTransversalF` that can be shrunk from the back.
func FrontTransversalB[T any](r BidirectionalRange[BidirectionalRange[T]]) BidirectionalRange[T] {
	return &frontTransversalBidirectionalResult[T]{baseBidrectionalChain[T]{r, false}}
}

// Chain produces the results of all the ranges together in a sequence.
func Chain[T any](ranges ...InputRange[T]) InputRange[T] {
	return Flatten[T](SliceRange(ranges))
}

// ChainF is `Chain` where the range can be saved.
func ChainF[T any](ranges ...ForwardRange[T]) ForwardRange[T] {
	return FlattenF(SliceRange(ranges))
}

// ChainB is `ChainF` that can be shrunk from the back.
func ChainB[T any](ranges ...BidirectionalRange[T]) BidirectionalRange[T] {
	return FlattenB(SliceRange(ranges))
}

// ChainS is `ChainB` accepting many slices.
func ChainS[T any](ranges ...[]T) BidirectionalRange[T] {
	return FlattenSS(ranges)
}
