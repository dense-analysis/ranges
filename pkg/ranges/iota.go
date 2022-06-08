package ranges

type iotaResult[T RealNumber] struct {
	current T
	end     T
}

func (ir *iotaResult[T]) Empty() bool           { return ir.current >= ir.end }
func (ir *iotaResult[T]) Front() T              { return ir.current }
func (ir *iotaResult[T]) PopFront()             { ir.current++ }
func (ir *iotaResult[T]) Save() ForwardRange[T] { return IotaStart(ir.current, ir.end) }

// Iota returns a range of values from 0 to `end`, excluding `end`, incrementing by `1`
func Iota[T RealNumber](end T) ForwardRange[T] {
	return &iotaResult[T]{0, end}
}

// IotaStart returns a range of values from `begin` to `end`, excluding `end`, incrementing by `1`
func IotaStart[T RealNumber](begin, end T) ForwardRange[T] {
	return &iotaResult[T]{begin, end}
}

type iotaStepResult[T RealNumber] struct {
	current T
	end     T
	step    T
}

func (ir *iotaStepResult[T]) Empty() bool           { return ir.current >= ir.end }
func (ir *iotaStepResult[T]) Front() T              { return ir.current }
func (ir *iotaStepResult[T]) PopFront()             { ir.current += ir.step }
func (ir *iotaStepResult[T]) Save() ForwardRange[T] { return IotaStep(ir.current, ir.end, ir.step) }

// IotaStep returns a range of values from `begin` to `end`, excluding `end`, incrementing by `step`
func IotaStep[T RealNumber](begin, end, step T) ForwardRange[T] {
	return &iotaStepResult[T]{begin, end, step}
}
