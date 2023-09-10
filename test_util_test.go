package ranges

import (
	"reflect"
	"testing"
)

func assertEqual[T any](t *testing.T, v1, v2 T) {
	if !reflect.DeepEqual(v1, v2) {
		t.Helper()
		t.Fatalf("%+v != %+v", v1, v2)
	}
}

func assertPanic(t *testing.T, expectedValue any, cb func()) {
	defer func() {
		if val := recover(); val == nil {
			t.Helper()
			t.Fatalf("Code did not panic!")
		} else if !reflect.DeepEqual(val, expectedValue) {
			t.Helper()
			t.Fatalf("Panic result: %+v != %+v", val, expectedValue)
		}
	}()

	cb()
}

func assertEmpty[T any](t *testing.T, r InputRange[T]) {
	t.Helper()

	if !r.Empty() {
		t.Fatal("The range was not empty")
	}
}

func assertEmptyF[T any](t *testing.T, r ForwardRange[T]) {
	t.Helper()
	assertEmpty[T](t, r)
}

func assertEmptyB[T any](t *testing.T, r BidirectionalRange[T]) {
	t.Helper()
	assertEmpty[T](t, r)
}

func assertEmptyR[T any](t *testing.T, r RandomAccessRange[T]) {
	t.Helper()
	assertEmpty[T](t, r)
}

func assertNotEmpty[T any](t *testing.T, r InputRange[T]) {
	t.Helper()

	if r.Empty() {
		t.Fatal("The range was empty")
	}
}

func assertNotEmptyF[T any](t *testing.T, r ForwardRange[T]) {
	t.Helper()
	assertNotEmpty[T](t, r)
}

func assertNotEmptyB[T any](t *testing.T, r BidirectionalRange[T]) {
	t.Helper()
	assertNotEmpty[T](t, r)
}

func assertNotEmptyR[T any](t *testing.T, r RandomAccessRange[T]) {
	t.Helper()
	assertNotEmpty[T](t, r)
}

func assertHasFront[T any](t *testing.T, r InputRange[T], value T) {
	t.Helper()

	if r.Empty() {
		t.Fatalf("Range was empty, did not start with %v", value)
	} else {
		front := r.Front()

		if !reflect.DeepEqual(front, value) {
			t.Fatalf("Range started with %v not %v", front, value)
		}
	}
}

func assertHasFrontF[T any](t *testing.T, r ForwardRange[T], value T) {
	t.Helper()
	assertHasFront[T](t, r, value)
}

func assertHasFrontB[T any](t *testing.T, r BidirectionalRange[T], value T) {
	t.Helper()
	assertHasFront[T](t, r, value)
}

func assertHasFrontR[T any](t *testing.T, r RandomAccessRange[T], value T) {
	t.Helper()
	assertHasFront[T](t, r, value)
}

func assertHasBack[T any](t *testing.T, r BidirectionalRange[T], value T) {
	t.Helper()

	if r.Empty() {
		t.Fatalf("Range was empty, did not end with %v", value)
	} else {
		back := r.Back()

		if !reflect.DeepEqual(back, value) {
			t.Fatalf("Range ended with %v not %v", back, value)
		}
	}
}

func assertHasBackR[T any](t *testing.T, r RandomAccessRange[T], value T) {
	t.Helper()
	assertHasBack(t, r, value)
}

func assertHasSaveableFront[T any](t *testing.T, r ForwardRange[T], value T) {
	t.Helper()
	assertHasFront[T](t, r, value)

	rSave := r.Save()
	r.PopFront()

	assertHasFront[T](t, rSave, value)
}

func assertHasSaveableFrontB[T any](t *testing.T, r BidirectionalRange[T], value T) {
	t.Helper()
	assertHasSaveableFront(t, r, value)
}

func assertHasSaveableFrontR[T any](t *testing.T, r RandomAccessRange[T], value T) {
	t.Helper()
	assertHasSaveableFront(t, r, value)
}

func assertHasSaveableBack[T any](t *testing.T, r BidirectionalRange[T], value T) {
	t.Helper()
	assertHasBack(t, r, value)

	rSave := r.SaveB()
	r.PopBack()

	assertHasBack(t, rSave, value)
}

func assertHasSaveableBackR[T any](t *testing.T, r RandomAccessRange[T], value T) {
	t.Helper()
	assertHasSaveableBack(t, r, value)
}

// lengthOnlyRange panics on all calls accept a Length() check.
// This can be used in tests to ensure we run length check optimizations.
type lengthOnlyRange[T any] struct {
	length int
}

func (r *lengthOnlyRange[T]) Len() int                     { return r.length }
func (r *lengthOnlyRange[T]) Front() T                     { panic("Front() not implemented") }
func (r *lengthOnlyRange[T]) PopFront()                    { panic("PopFront() not implemented") }
func (r *lengthOnlyRange[T]) Empty() bool                  { panic("Empty() not implemented") }
func (r *lengthOnlyRange[T]) Back() T                      { panic("Back() not implemented") }
func (r *lengthOnlyRange[T]) PopBack()                     { panic("PopBack() not implemented") }
func (r *lengthOnlyRange[T]) Get(index int) T              { panic("Get() not implemented") }
func (r *lengthOnlyRange[T]) Save() ForwardRange[T]        { panic("Save() not implemented") }
func (r *lengthOnlyRange[T]) SaveB() BidirectionalRange[T] { panic("SaveB() not implemented") }
func (r *lengthOnlyRange[T]) SaveR() RandomAccessRange[T]  { panic("SaveB() not implemented") }

// lengthOnly returns a RandomAccessRange that only implements Len()
// This can be used in tests to ensure we run length check optimizations.
func lengthOnly[T any](length int) RandomAccessRange[T] {
	return &lengthOnlyRange[T]{length}
}

// badLengthRangeImpl implements badLengthRange
type badLengthRangeImpl[T any] struct {
	RandomAccessRange[T]
	badLength int
}

func (r *badLengthRangeImpl[T]) Len() int { return r.badLength }

// badLengthRange returns a RandomAccess range where the length can set improperly.
// this can be used to test how algorithms behave with flawed input.
func badLengthRange[T any](r RandomAccessRange[T], badLength int) RandomAccessRange[T] {
	return &badLengthRangeImpl[T]{r, badLength}
}
