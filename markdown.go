package main

import (
	"fmt"
	"regexp"
	"strings"
)

type rule struct {
	reg         *regexp.Regexp
	replaceFunc replaceFunc
}

type replaceFunc func(string) string

var (
	h6 = newRule(`^###### .*`, trimPrefix("###### "), withTag("h6"))
	h5 = newRule(`^##### .*`, trimPrefix("##### "), withTag("h5"))
	h4 = newRule(`^#### .*`, trimPrefix("####  "), withTag("h4"))
	h3 = newRule(`^### .*`, trimPrefix("### "), withTag("h3"))
	h2 = newRule(`^## .*`, trimPrefix("## "), withTag("h2"))
	h1 = newRule(`^# .*`, trimPrefix("# "), withTag("h1"))

	strongEm = newRule(`\*\*\*.*\*\*\*`, trim("*"), withTag("em", "strong"))
	strong   = newRule(`\*\*.*\*\*`, trim("*"), withTag("strong"))
	em       = newRule(`\*.*\*`, trim("*"), withTag("em"))

	strongEm2 = newRule(`\_\_\_.*\_\_\_`, trim("_"), withTag("em", "strong"))
	strong2   = newRule(`\_\_.*\_\_`, trim("_"), withTag("strong"))
	em2       = newRule(`\_.*\_`, trim("_"), withTag("em"))

	ul        = newRuleGreedy(`^- .*\n`)
	ullist    = newRule(`- `, trimPrefix("_ "), withTag("li"))
	ollistPre = newRule(`^[0-9]*. `)
)

const (
	greedyMode = iota
	nonGreedymode
)

func newRule(regStr string, replaceFuncs ...replaceFunc) *rule {
	return initRule(regStr, false, replaceFuncs...)
}

func newRuleGreedy(regStr string, replaceFuncs ...replaceFunc) *rule {
	return initRule(regStr, true, replaceFuncs...)
}

func initRule(regStr string, isGreedy bool, replaceFuncs ...replaceFunc) *rule {
	r := regexp.MustCompile(regStr)
	if isGreedy {
		r.Longest()
	}
	f := chainsAll(replaceFuncs...)
	return &rule{r, f}
}

func markdownToHTML(input string) string {

	return chains(
		replaceByRules(strongEm),
		replaceByRules(strong),
	)(input)
}

func replaceByRules(rule *rule) replaceFunc {
	return func(input string) string {
		return rule.reg.ReplaceAllStringFunc(input, rule.replaceFunc)
	}
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
