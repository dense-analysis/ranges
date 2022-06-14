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

// Flatten combines a range of ranges into one straight range
func Flatten[T any](r InputRange[InputRange[T]]) InputRange[T] {
	return &flattenResult[T]{baseChain[T]{r, false}}
}

// FlattenF is `Flatten` where the range can be saved.
func FlattenF[T any](r ForwardRange[ForwardRange[T]]) ForwardRange[T] {
	return &flattenForwardResult[T]{baseForwardChain[T]{r, false}}
}

// FlattenS is `FlattenF` accepting a slice.
func FlattenS[T any](r []ForwardRange[T]) ForwardRange[T] {
	return FlattenF(SliceRange(r))
}

// FlattenSS is `FlattenS` accepting a slice of slices.
func FlattenSS[T any](r [][]T) ForwardRange[T] {
	return FlattenF(MapS(r, SliceRange[T]))
}

// FrontTransversal yields the first value in each range, skipping empty ranges.
func FrontTransversal[T any](r InputRange[InputRange[T]]) InputRange[T] {
	return &frontTransversalResult[T]{baseChain[T]{r, false}}
}

// FrontTransversalF is `FrontTransversal` where the range can be saved.
func FrontTransversalF[T any](r ForwardRange[ForwardRange[T]]) ForwardRange[T] {
	return &frontTransversalForwardResult[T]{baseForwardChain[T]{r, false}}
}

// Chain produces the results of all the ranges together in a sequence.
func Chain[T any](ranges ...InputRange[T]) InputRange[T] {
	return Flatten[T](SliceRange(ranges))
}

// ChainF is `Chain` where the range can be saved.
func ChainF[T any](ranges ...ForwardRange[T]) ForwardRange[T] {
	return FlattenF(SliceRange(ranges))
}

// ChainS is `ChainF` accepting many slices.
func ChainS[T any](ranges ...[]T) ForwardRange[T] {
	return FlattenF(MapS(ranges, SliceRange[T]))
}
