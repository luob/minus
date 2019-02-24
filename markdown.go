package main

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

type rule struct {
	reg         *regexp.Regexp
	replaceFunc replaceFunc
}

type replaceFunc func(string) string

func newRule(regStr string, replaceFuncs ...replaceFunc) *rule {
	r := regexp.MustCompile(regStr)
	f := chainsAll(replaceFuncs...)
	return &rule{r, f}
}

var rules = []*rule{

	// headings
	newRule(`^###### .*`, trimPrefix("###### "), withTag("h6")),
	newRule(`^##### .*`, trimPrefix("##### "), withTag("h5")),
	newRule(`^#### .*`, trimPrefix("####  "), withTag("h4")),
	newRule(`^### .*`, trimPrefix("### "), withTag("h3")),
	newRule(`^## .*`, trimPrefix("## "), withTag("h2")),
	newRule(`^# .*`, trimPrefix("# "), withTag("h1")),

	// inline tags
	newRule(`\*\*\*.*\*\*\*`, trim("*"), withTag("em", "strong")),
	newRule(`\*\*.*\*\*`, trim("*"), withTag("strong")),
	newRule(`\*.*\*`, trim("*"), withTag("em")),

	// inline tags type 2
	newRule(`\_\_\_.*\_\_\_`, trim("_"), withTag("em", "strong")),
	newRule(`\_\_.*\_\_`, trim("_"), withTag("strong")),
	newRule(`\_.*\_`, trim("_"), withTag("em")),

	// list
	newRule(`^- `, each(trimPrefix("- "), withTag("li")), withTag("ol")),
	newRule(`^[0-9]+. `, each(trim("0123456789. "), withTag("li")), withTag("ul")),

	// code blocks
	newRule(`^`+"```", trim("`"), asCode()),

	// html
	newRule(`^<.*>`, pass()),

	// paragragh
	newRule(`.*`, withTag("p")),
}

var (
	splitRegexp = regexp.MustCompile(`\n{2,}`)
	langRegexp  = regexp.MustCompile(`\b.*\b`)
)

func markdownToHTML(input string) template.HTML {
	s := ""
	for _, block := range splitRegexp.Split(input, -1) {
		s += parseBlock(block)
	}
	return template.HTML(s)
}

func parseBlock(input string) string {
	for _, rule := range rules {
		if rule.reg.MatchString(input) {
			return rule.replaceFunc(input)
		}
	}
	return input
}

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
