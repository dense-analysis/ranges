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
				I(F(B(Only(1, 2, 3)))),
				I(F(B(Only(4, 5)))),
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
		ChainF(F(B(Only(1, 2))), F(B(Null[int]())), F(B(Only(3, 4)))),
		ChainF(F(B(Null[int]())), F(B(Only(5, 6))), F(B(Null[int]()))),
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
		ChainB(B(Only(1, 2)), B(Null[int]()), B(Only(3, 4))),
		ChainB(B(Null[int]()), B(Only(5, 6)), B(Null[int]())),
	)

	chain.PopBack()
	chain.PopBack()
	chain.PopBack()

	savedChain := chain.SaveB()

	assertEqual(t, SliceB(chain), []int{1, 2, 3})
	assertEqual(t, SliceB(savedChain), []int{1, 2, 3})

	chain2 := ChainB(B(Only(1, 2)), B(Null[int]()), B(Only(3, 4)), B(Null[int]()))

	assertEqual(t, SliceB(Retro(chain2)), []int{4, 3, 2, 1})
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
				I(F(B(Only(1, 2)))),
				I(F(B(Only[int]()))),
				I(F(B(Only(3, 4)))),
				I(F(B(Only[int]()))),
				I(F(B(Only(5, 6)))),
				I(F(B(Only[int]()))),
			),
		),
	)

	assertEqual(t, sliceCopy, []int{1, 3, 5})
}

func TestFrontTransversalF(t *testing.T) {
	t.Parallel()

	empty := FrontTransversalF(F(B(Null[ForwardRange[int]]())))

	if !empty.Empty() {
		t.Fatal("An empty transversal was not empty")
	}

	transversal := FrontTransversalF(
		F(B(Only(
			F(B(Only(1, 2))),
			F(B(Only[int]())),
			F(B(Only(3, 4))),
			F(B(Only[int]())),
			F(B(Only(5, 6))),
			F(B(Only[int]())),
		))),
	)

	transversal.PopFront()
	savedTransversal := transversal.Save()

	assertEqual(t, SliceF(transversal), []int{3, 5})
	assertEqual(t, SliceF(savedTransversal), []int{3, 5})
}

func TestFrontTransversalB(t *testing.T) {
	t.Parallel()

	empty := FrontTransversalB(B(Null[BidirectionalRange[int]]()))

	if !empty.Empty() {
		t.Fatal("An empty transversal was not empty")
	}

	transversal := FrontTransversalB(
		B(Only(
			B(Only(1, 2)),
			B(Only[int]()),
			B(Only(3, 4)),
			B(Only[int]()),
			B(Only(5, 6)),
			B(Only[int]()),
		)),
	)

	transversal.PopFront()
	savedTransversal := transversal.SaveB()

	assertEqual(t, SliceB(transversal), []int{3, 5})
	assertEqual(t, SliceB(savedTransversal), []int{3, 5})

	transversal2 := FrontTransversalB(
		B(Only(
			B(Only(1, 2)),
			B(Only[int]()),
			B(Only(3, 4)),
			B(Only[int]()),
			B(Only(5, 6)),
			B(Only[int]()),
		)),
	)

	transversal2.PopBack()
	savedTransversal2 := transversal2.SaveB()

	assertEqual(t, SliceB(Retro(transversal2)), []int{3, 1})
	assertEqual(t, SliceB(Retro(savedTransversal2)), []int{3, 1})
}
