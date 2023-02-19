package ranges

import (
	"strconv"
	"testing"
)

func intToFloat(x int) float64       { return float64(x) }
func floatToString(x float64) string { return strconv.FormatFloat(x, 'f', 2, 64) }
func stringToFloat(x string) float64 { val, _ := strconv.ParseFloat(x, 64); return val }
func floatToInt(x float64) int       { return int(float64(x)) }

func TestPipe2(t *testing.T) {
	t.Parallel()

	result := Pipe2(
		intToFloat,
		floatToString,
	)(92)

	assertEqual(t, result, "92.00")
}

func TestPipe3(t *testing.T) {
	t.Parallel()

	result := Pipe3(
		intToFloat,
		floatToString,
		stringToFloat,
	)(92)

	assertEqual(t, result, 92)
}

func TestPipe10(t *testing.T) {
	t.Parallel()

	result := Pipe10(
		intToFloat,
		floatToString,
		stringToFloat,
		floatToInt,
		intToFloat,
		floatToString,
		stringToFloat,
		floatToInt,
		intToFloat,
		floatToString,
	)(92)

	assertEqual(t, result, "92.00")
}

func TestCompose2(t *testing.T) {
	t.Parallel()

	result := Compose2(
		floatToString,
		intToFloat,
	)(92)

	assertEqual(t, result, "92.00")
}

func TestCompose3(t *testing.T) {
	t.Parallel()

	result := Compose3(
		stringToFloat,
		floatToString,
		intToFloat,
	)(92)

	assertEqual(t, result, 92)
}

func TestCompose10(t *testing.T) {
	t.Parallel()

	result := Compose10(
		floatToString,
		intToFloat,
		floatToInt,
		stringToFloat,
		floatToString,
		intToFloat,
		floatToInt,
		stringToFloat,
		floatToString,
		intToFloat,
	)(92)

	assertEqual(t, result, "92.00")
}
