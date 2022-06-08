package ranges

import "testing"

func TestRepeat(t *testing.T) {
	t.Parallel()

	repeater := Repeat(3)

	assertHasFrontF(t, repeater, 3)
	assertNotEmptyF(t, repeater)
	repeater.PopFront()

	assertHasFrontF(t, repeater, 3)
	assertNotEmptyF(t, repeater)
	repeater.PopFront()

	assertHasFrontF(t, repeater, 3)
	assertNotEmptyF(t, repeater)
	repeater.PopFront()

	repeater = repeater.Save()

	assertHasFrontF(t, repeater, 3)
	assertNotEmptyF(t, repeater)
	repeater.PopFront()
}

func TestGenerate(t *testing.T) {
	t.Parallel()

	generator := Generate(func() int { return 3 })

	assertHasFrontF(t, generator, 3)
	assertNotEmptyF(t, generator)
	generator.PopFront()

	assertHasFrontF(t, generator, 3)
	assertNotEmptyF(t, generator)
	generator.PopFront()

	assertHasFrontF(t, generator, 3)
	assertNotEmptyF(t, generator)
	generator.PopFront()

	generator = generator.Save()

	assertHasFrontF(t, generator, 3)
	assertNotEmptyF(t, generator)
	generator.PopFront()
}
