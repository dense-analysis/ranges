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

func assertHasSaveableFront[T any](t *testing.T, r ForwardRange[T], value T) {
	t.Helper()
	assertHasFront[T](t, r, value)

	rSave := r.Save()
	r.PopFront()

	assertHasFront[T](t, rSave, value)
}

func assertHasSaveableFrontB[T any](t *testing.T, r BidirectionalRange[T], value T) {
	t.Helper()
	assertHasSaveableFront(t, r.(ForwardRange[T]), value)
}

func assertHasSaveableBack[T any](t *testing.T, r BidirectionalRange[T], value T) {
	t.Helper()
	assertHasBack(t, r, value)

	rSave := r.SaveB()
	r.PopFront()

	assertHasBack(t, rSave, value)
}
