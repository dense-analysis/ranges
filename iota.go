package ranges

import "math"

const nearlyZero = 1e-12

type iotaResult[T RealNumber] struct {
	current T
	end     T
}

func (ir *iotaResult[T]) Empty() bool                  { return ir.current >= ir.end }
func (ir *iotaResult[T]) Front() T                     { return ir.current }
func (ir *iotaResult[T]) PopFront()                    { ir.current++ }
func (ir *iotaResult[T]) Back() T                      { return ir.end - 1 }
func (ir *iotaResult[T]) PopBack()                     { ir.end-- }
func (ir *iotaResult[T]) Save() ForwardRange[T]        { return ir.SaveB() }
func (ir *iotaResult[T]) SaveB() BidirectionalRange[T] { return IotaStart(ir.current, ir.end) }

// Iota returns a range of values from 0 to `end`, excluding `end`, incrementing by `1`
func Iota[T RealNumber](end T) BidirectionalRange[T] {
	return &iotaResult[T]{0, end}
}

// IotaStart returns a range of values from `begin` to `end`, excluding `end`, incrementing by `1`
func IotaStart[T RealNumber](begin, end T) BidirectionalRange[T] {
	return &iotaResult[T]{begin, end}
}

type iotaStepResult[T RealNumber] struct {
	current T
	end     T
	step    T
}

func (ir *iotaStepResult[T]) lenOfRange() T {
	return 1 + (ir.end-1-ir.current)/ir.step
}

func (ir *iotaStepResult[T]) Empty() bool { return ir.current >= ir.end }
func (ir *iotaStepResult[T]) Front() T    { return ir.current }
func (ir *iotaStepResult[T]) PopFront()   { ir.current += ir.step }
func (ir *iotaStepResult[T]) Back() T {
	remainder := math.Mod(float64(ir.end-ir.current), float64(ir.step))

	if remainder > nearlyZero {
		return ir.end - T(remainder)
	}

	return ir.end - ir.step
}
func (ir *iotaStepResult[T]) PopBack()              { ir.end -= ir.step }
func (ir *iotaStepResult[T]) Save() ForwardRange[T] { return ir.SaveB() }
func (ir *iotaStepResult[T]) SaveB() BidirectionalRange[T] {
	return IotaStep(ir.current, ir.end, ir.step)
}

// IotaStep returns a range of values from `begin` to `end`, excluding `end`, incrementing by `step`
func IotaStep[T RealNumber](begin, end, step T) BidirectionalRange[T] {
	return &iotaStepResult[T]{begin, end, step}
}
