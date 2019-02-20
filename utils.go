package main

func composeAll(replaceFuncs ...replaceFunc) replaceFunc {
	f := replaceFuncs[0]
	for i := 1; i < len(replaceFuncs); i++ {
		f = compose(f, replaceFuncs[i])
	}
	return f
}

func compose(f1, f2 replaceFunc) replaceFunc {
	return func(input string) string {
		return f2(f1(input))
	}
}
