package ranges

import "testing"

func TestIota(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceB(Iota(0)), []int{})
	assertEqual(t, SliceB(Iota(1)), []int{0})
	assertEqual(t, SliceB(Iota(2)), []int{0, 1})
}

func TestIotaStart(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceB(IotaStart(2, 2)), []int{})
	assertEqual(t, SliceB(IotaStart(2, 3)), []int{2})
	assertEqual(t, SliceB(IotaStart(2, 4)), []int{2, 3})
	assertEqual(t, SliceB(Retro(IotaStart(2, 4))), []int{3, 2})
}

func TestIotaStep(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceB(IotaStep(2.0, 2.0, 0.5)), []float64{})
	assertEqual(t, SliceB(IotaStep(2.0, 3.0, 0.5)), []float64{2.0, 2.5})
	assertEqual(t, SliceB(IotaStep(2.0, 4.0, 0.5)), []float64{2.0, 2.5, 3.0, 3.5})
	assertEqual(t, SliceB(IotaStep(2.0, 4.0, 1.2)), []float64{2.0, 3.2})
	assertEqual(t, SliceB(Retro(IotaStep(2.0, 4.0, 1.2))), []float64{3.2, 2.0})
	assertEqual(t, SliceB(IotaStep(2.0, 4.4, 1.2)), []float64{2, 3.2})
	assertEqual(t, SliceB(Retro(IotaStep(2.0, 4.4, 1.2))), []float64{3.2, 2.0})
	assertEqual(t, SliceB(IotaStep(2.0, 4.5, 1.2)), []float64{2.0, 3.2, 4.4})
	assertEqual(t, SliceB(Retro(IotaStep(2.0, 4.5, 1.2))), []float64{4.4, 3.2, 2.0})
}
