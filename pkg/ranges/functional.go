package ranges

func Compose2[A, B, C any](f1 func(x B) C, f2 func(x A) B) func(x A) C {
	return func(x A) C { return f1(f2(x)) }
}

func Compose3[A, B, C, D any](
	f1 func(x C) D,
	f2 func(x B) C,
	f3 func(x A) B,
) func(x A) D {
	return Compose2(Compose2(f1, f2), f3)
}

func Compose4[A, B, C, D, E any](
	f1 func(x D) E,
	f2 func(x C) D,
	f3 func(x B) C,
	f4 func(x A) B,
) func(x A) E {
	return Compose2(Compose3(f1, f2, f3), f4)
}

func Compose5[A, B, C, D, E, F any](
	f1 func(x E) F,
	f2 func(x D) E,
	f3 func(x C) D,
	f4 func(x B) C,
	f5 func(x A) B,
) func(x A) F {
	return Compose2(Compose4(f1, f2, f3, f4), f5)
}

func Compose6[A, B, C, D, E, F, G any](
	f1 func(x F) G,
	f2 func(x E) F,
	f3 func(x D) E,
	f4 func(x C) D,
	f5 func(x B) C,
	f6 func(x A) B,
) func(x A) G {
	return Compose2(Compose5(f1, f2, f3, f4, f5), f6)
}

func Compose7[A, B, C, D, E, F, G, H any](
	f1 func(x G) H,
	f2 func(x F) G,
	f3 func(x E) F,
	f4 func(x D) E,
	f5 func(x C) D,
	f6 func(x B) C,
	f7 func(x A) B,
) func(x A) H {
	return Compose2(Compose6(f1, f2, f3, f4, f5, f6), f7)
}

func Compose8[A, B, C, D, E, F, G, H, I any](
	f1 func(x H) I,
	f2 func(x G) H,
	f3 func(x F) G,
	f4 func(x E) F,
	f5 func(x D) E,
	f6 func(x C) D,
	f7 func(x B) C,
	f8 func(x A) B,
) func(x A) I {
	return Compose2(Compose7(f1, f2, f3, f4, f5, f6, f7), f8)
}

func Compose9[A, B, C, D, E, F, G, H, I, J any](
	f1 func(x I) J,
	f2 func(x H) I,
	f3 func(x G) H,
	f4 func(x F) G,
	f5 func(x E) F,
	f6 func(x D) E,
	f7 func(x C) D,
	f8 func(x B) C,
	f9 func(x A) B,
) func(x A) J {
	return Compose2(Compose8(f1, f2, f3, f4, f5, f6, f7, f8), f9)
}

func Compose10[A, B, C, D, E, F, G, H, I, J, K any](
	f1 func(x J) K,
	f2 func(x I) J,
	f3 func(x H) I,
	f4 func(x G) H,
	f5 func(x F) G,
	f6 func(x E) F,
	f7 func(x D) E,
	f8 func(x C) D,
	f9 func(x B) C,
	f10 func(x A) B,
) func(x A) K {
	return Compose2(Compose9(f1, f2, f3, f4, f5, f6, f7, f8, f9), f10)
}

func Pipe2[A, B, C any](f1 func(x A) B, f2 func(x B) C) func(x A) C {
	return func(x A) C { return f2(f1(x)) }
}

func Pipe3[A, B, C, D any](
	f1 func(x A) B,
	f2 func(x B) C,
	f3 func(x C) D,
) func(x A) D {
	return Pipe2(Pipe2(f1, f2), f3)
}

func Pipe4[A, B, C, D, E any](
	f1 func(x A) B,
	f2 func(x B) C,
	f3 func(x C) D,
	f4 func(x D) E,
) func(x A) E {
	return Pipe2(Pipe3(f1, f2, f3), f4)
}

func Pipe5[A, B, C, D, E, F any](
	f1 func(x A) B,
	f2 func(x B) C,
	f3 func(x C) D,
	f4 func(x D) E,
	f5 func(x E) F,
) func(x A) F {
	return Pipe2(Pipe4(f1, f2, f3, f4), f5)
}

func Pipe6[A, B, C, D, E, F, G any](
	f1 func(x A) B,
	f2 func(x B) C,
	f3 func(x C) D,
	f4 func(x D) E,
	f5 func(x E) F,
	f6 func(x F) G,
) func(x A) G {
	return Pipe2(Pipe5(f1, f2, f3, f4, f5), f6)
}

func Pipe7[A, B, C, D, E, F, G, H any](
	f1 func(x A) B,
	f2 func(x B) C,
	f3 func(x C) D,
	f4 func(x D) E,
	f5 func(x E) F,
	f6 func(x F) G,
	f7 func(x G) H,
) func(x A) H {
	return Pipe2(Pipe6(f1, f2, f3, f4, f5, f6), f7)
}

func Pipe8[A, B, C, D, E, F, G, H, I any](
	f1 func(x A) B,
	f2 func(x B) C,
	f3 func(x C) D,
	f4 func(x D) E,
	f5 func(x E) F,
	f6 func(x F) G,
	f7 func(x G) H,
	f8 func(x H) I,
) func(x A) I {
	return Pipe2(Pipe7(f1, f2, f3, f4, f5, f6, f7), f8)
}

func Pipe9[A, B, C, D, E, F, G, H, I, J any](
	f1 func(x A) B,
	f2 func(x B) C,
	f3 func(x C) D,
	f4 func(x D) E,
	f5 func(x E) F,
	f6 func(x F) G,
	f7 func(x G) H,
	f8 func(x H) I,
	f9 func(x I) J,
) func(x A) J {
	return Pipe2(Pipe8(f1, f2, f3, f4, f5, f6, f7, f8), f9)
}

func Pipe10[A, B, C, D, E, F, G, H, I, J, K any](
	f1 func(x A) B,
	f2 func(x B) C,
	f3 func(x C) D,
	f4 func(x D) E,
	f5 func(x E) F,
	f6 func(x F) G,
	f7 func(x G) H,
	f8 func(x H) I,
	f9 func(x I) J,
	f10 func(x J) K,
) func(x A) K {
	return Pipe2(Pipe9(f1, f2, f3, f4, f5, f6, f7, f8, f9), f10)
}
