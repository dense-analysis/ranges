package ranges

// This file contains operators in the style of Python's operator module.

// Lt returns a < b
func Lt[T Ordered](a, b T) bool { return a < b }

// Le returns a <= b
func Le[T Ordered](a, b T) bool { return a <= b }

// Eq returns a == b
func Eq[T comparable](a, b T) bool { return a == b }

// Ne returns a != b
func Ne[T comparable](a, b T) bool { return a != b }

// Ge returns a >= b
func Ge[T Ordered](a, b T) bool { return a >= b }

// Gt returns a > b
func Gt[T Ordered](a, b T) bool { return a > b }
