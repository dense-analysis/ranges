package ranges

import "testing"

func TestRepeat(t *testing.T) {
	t.Parallel()

	repeater := Repeat(3)

	for i := 0; i < 3; i++ {
		assertHasFront(t, repeater, 3)
		assertNotEmpty(t, repeater)
		repeater.PopFront()
	}

	repeater = repeater.SaveB()

	for i := 0; i < 3; i++ {
		assertHasBack(t, repeater, 3)
		assertNotEmpty(t, repeater)
		repeater.PopBack()
	}
}

func TestGenerate(t *testing.T) {
	t.Parallel()

	generator := Generate(func() int { return 3 })

	for i := 0; i < 3; i++ {
		assertHasFront(t, generator, 3)
		assertNotEmpty(t, generator)
		generator.PopFront()
	}

	generator = generator.SaveB()

	for i := 0; i < 3; i++ {
		assertHasBack(t, generator, 3)
		assertNotEmpty(t, generator)
		generator.PopBack()
	}
}
