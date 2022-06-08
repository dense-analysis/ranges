package ranges

import "testing"

func TestIota(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceF(Iota(0)), []int{})
	assertEqual(t, SliceF(Iota(1)), []int{0})
	assertEqual(t, SliceF(Iota(2)), []int{0, 1})
}

func TestIotaStart(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceF(IotaStart(2, 2)), []int{})
	assertEqual(t, SliceF(IotaStart(2, 3)), []int{2})
	assertEqual(t, SliceF(IotaStart(2, 4)), []int{2, 3})
}

func TestIotaStep(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceF(IotaStep(2.0, 2.0, 0.5)), []float64{})
	assertEqual(t, SliceF(IotaStep(2.0, 3.0, 0.5)), []float64{2.0, 2.5})
	assertEqual(t, SliceF(IotaStep(2.0, 4.0, 0.5)), []float64{2.0, 2.5, 3.0, 3.5})
}
