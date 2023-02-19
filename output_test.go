package ranges

import "testing"

func TestNullSink(t *testing.T) {
	t.Parallel()

	sink := NullSink[int]()

	if sink.Put(3) != nil {
		t.Fatal("Put() on NullSink failed")
	}
}

func TestAssignSink(t *testing.T) {
	t.Parallel()

	output := make([]int, 3)

	sink := AssignSink[int](SlicePtrRange(output))

	sink.Put(1)

	sinkRef := sink

	sinkRef.Put(2)
	sinkRef.Put(3)

	assertEqual(t, output, []int{1, 2, 3})
}

func TestSliceSink(t *testing.T) {
	t.Parallel()

	output := make([]int, 0, 3)

	sink := SliceSink(&output)
	sink.Put(1)
	sink.Put(2)
	sink.Put(3)

	assertEqual(t, output, []int{1, 2, 3})
}
