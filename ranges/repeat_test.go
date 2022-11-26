package ranges

import "testing"

func TestRepeat(t *testing.T) {
	t.Parallel()

	repeater := Repeat(3)

	for i := 0; i < 3; i++ {
		assertHasFrontB(t, repeater, 3)
		assertNotEmptyB(t, repeater)
		repeater.PopFront()
	}

	repeater = repeater.SaveB()

	for i := 0; i < 3; i++ {
		assertHasBack(t, repeater, 3)
		assertNotEmptyB(t, repeater)
		repeater.PopBack()
	}
}

func TestGenerate(t *testing.T) {
	t.Parallel()

	generator := Generate(func() int { return 3 })

	for i := 0; i < 3; i++ {
		assertHasFrontB(t, generator, 3)
		assertNotEmptyB(t, generator)
		generator.PopFront()
	}

	generator = generator.SaveB()

	for i := 0; i < 3; i++ {
		assertHasBack(t, generator, 3)
		assertNotEmptyB(t, generator)
		generator.PopBack()
	}
}
