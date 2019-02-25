package markdown

import "strings"

type replaceFunc func(string) string

func chainsAll(replaceFuncs ...replaceFunc) replaceFunc {
	f := replaceFuncs[0]
	for i := 1; i < len(replaceFuncs); i++ {
		f = chains(f, replaceFuncs[i])
	}
	return f
}

func chains(f1, f2 replaceFunc) replaceFunc {
	return func(input string) string {
		return f2(f1(input))
	}
}

func each(replaceFuncs ...replaceFunc) replaceFunc {
	return func(input string) string {
		output := ""
		f := chainsAll(replaceFuncs...)
		lines := strings.Split(input, "\n")
		for _, line := range lines {
			output += f(line)
		}
		return output
	}
}
