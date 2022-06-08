package ranges

import "testing"

func TestEq(t *testing.T) {
	t.Parallel()

	if !Eq(1, 1) {
		t.Error("1 == 2 was false")
	}

	if !Eq("abc", "abc") {
		t.Error("abc == abc was false")
	}

	if Eq("abc", "xyz") {
		t.Error("abc == xyz was true")
	}
}

func TestLe(t *testing.T) {
	t.Parallel()

	if !Le(1, 1) {
		t.Error("1 <= 1 was false")
	}

	if !Le(1, 2) {
		t.Error("1 <= 2 was false")
	}

	if Le(2, 1) {
		t.Error("2 <= 1 was true")
	}
}

func TestLt(t *testing.T) {
	t.Parallel()

	if Lt(1, 1) {
		t.Error("1 < 1 was true")
	}

	if !Lt(1, 2) {
		t.Error("1 < 2 was false")
	}

	if Lt(2, 1) {
		t.Error("2 < 1 was true")
	}
}

func TestNeq(t *testing.T) {
	t.Parallel()

	if Ne(1, 1) {
		t.Error("1 != 1 was true")
	}

	if Ne("abc", "abc") {
		t.Error("abc != abc was true")
	}

	if !Ne("abc", "xyz") {
		t.Error("abc != xyz was false")
	}
}

func TestGe(t *testing.T) {
	t.Parallel()

	if !Ge(1, 1) {
		t.Error("1 >= 1 was false")
	}

	if Ge(1, 2) {
		t.Error("1 >= 2 was true")
	}

	if !Ge(2, 1) {
		t.Error("2 >= 1 was false")
	}
}

func TestGt(t *testing.T) {
	t.Parallel()

	if Gt(1, 1) {
		t.Error("1 > 1 was true")
	}

	if Gt(1, 2) {
		t.Error("1 > 2 was true")
	}

	if !Gt(2, 1) {
		t.Error("2 > 1 was false")
	}
}
