package markdown

import (
	"fmt"
	"strings"
)

func pass() replaceFunc {
	return func(string) string {
		return ""
	}
}

func trim(cutset string) replaceFunc {
	return func(input string) string {
		return strings.Trim(input, cutset)
	}
}

func trimPrefix(prefix string) replaceFunc {
	return func(input string) string {
		return strings.TrimPrefix(input, prefix)
	}
}

func trimSuffix(suffix string) replaceFunc {
	return func(input string) string {
		return strings.TrimSuffix(input, suffix)
	}
}

func withTag(tags ...string) replaceFunc {
	return func(input string) string {
		for _, tag := range tags {
			input = fmt.Sprintf("<%s>%s</%s>", tag, input, tag)
		}
		return input
	}
}

func asCode() func(string) string {
	return func(input string) string {
		// lang := langRegexp.FindString(input)
		arr := strings.Split(input, "\n")
		lang := arr[0]
		input = strings.Join(arr[1:], "\n")
		return fmt.Sprintf("<pre><code class=\"%s\">%s</code></pre>", lang, input)
	}
}
