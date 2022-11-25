package ranges

import "testing"

func TestChain(t *testing.T) {
	t.Parallel()

	emptyChain := Chain[int]()

	if !emptyChain.Empty() {
		t.Fatal("An empty chain was not empty")
	}

	sliceCopy := Slice(
		Chain(
			Chain[int](Only(1, 2), Null[int](), Only(3, 4)),
			Chain[int](Null[int](), Only(5, 6), Null[int]()),
		),
	)

	assertEqual(t, sliceCopy, []int{1, 2, 3, 4, 5, 6})
}

func TestChainIsLazy(t *testing.T) {
	t.Parallel()

	// This will panic if it's not lazy.
	Chain[int](nil)
}

func TestFlatten(t *testing.T) {
	t.Parallel()

	result := Slice(
		Flatten[int](
			Only(
				I(F(Only(1, 2, 3))),
				I(F(Only(4, 5))),
			),
		),
	)

	assertEqual(t, result, []int{1, 2, 3, 4, 5})
}

func TestChainF(t *testing.T) {
	t.Parallel()

	emptyChain := ChainF[int]()

	if !emptyChain.Empty() {
		t.Fatal("An empty chain was not empty")
	}

	if !emptyChain.Save().Empty() {
		t.Fatal("An saved empty chain was not empty!")
	}

	chain := ChainF(
		ChainF(F(Only(1, 2)), F(Null[int]()), F(Only(3, 4))),
		ChainF(F(Null[int]()), F(Only(5, 6)), F(Null[int]())),
	)

	chain.PopFront()
	chain.PopFront()
	chain.PopFront()

	savedChain := chain.Save()

	assertEqual(t, SliceF(chain), []int{4, 5, 6})
	assertEqual(t, SliceF(savedChain), []int{4, 5, 6})
}

func TestChainFIsLazy(t *testing.T) {
	t.Parallel()

	// This will panic if it's not lazy.
	ChainF[int](nil)
}

func TestChainB(t *testing.T) {
	t.Parallel()

	emptyChain := ChainB[int]()

	if !emptyChain.Empty() {
		t.Fatal("An empty chain was not empty")
	}

	if !emptyChain.Save().Empty() {
		t.Fatal("An saved empty chain was not empty!")
	}

	chain := ChainB(
		ChainB(Only(1, 2), Null[int](), Only(3, 4)),
		ChainB(Null[int](), Only(5, 6), Null[int]()),
	)

	chain.PopBack()
	chain.PopBack()
	chain.PopBack()

	savedChain := chain.SaveB()

	assertEqual(t, SliceB(chain), []int{1, 2, 3})
	assertEqual(t, SliceB(savedChain), []int{1, 2, 3})
}

func TestChainBIsLazy(t *testing.T) {
	t.Parallel()

	// This will panic if it's not lazy.
	ChainB[int](nil)
}

func TestFrontTransversal(t *testing.T) {
	t.Parallel()

	empty := FrontTransversal[int](Null[InputRange[int]]())

	if !empty.Empty() {
		t.Fatal("An empty transversal was not empty")
	}

	sliceCopy := Slice(
		FrontTransversal[int](
			Only(
				I(F(Only(1, 2))),
				I(F(Only[int]())),
				I(F(Only(3, 4))),
				I(F(Only[int]())),
				I(F(Only(5, 6))),
				I(F(Only[int]())),
			),
		),
	)

	assertEqual(t, sliceCopy, []int{1, 3, 5})
}

func TestFrontTransversalF(t *testing.T) {
	t.Parallel()

	empty := FrontTransversalF(F(Null[ForwardRange[int]]()))

	if !empty.Empty() {
		t.Fatal("An empty transversal was not empty")
	}

	transversal := FrontTransversalF(
		F(Only(
			F(Only(1, 2)),
			F(Only[int]()),
			F(Only(3, 4)),
			F(Only[int]()),
			F(Only(5, 6)),
			F(Only[int]()),
		)),
	)

	transversal.PopFront()
	savedTransversal := transversal.Save()

	assertEqual(t, SliceF(transversal), []int{3, 5})
	assertEqual(t, SliceF(savedTransversal), []int{3, 5})
}
