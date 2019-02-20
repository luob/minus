package main

import (
	"fmt"
	"regexp"
)

type rule struct {
	reg         *regexp.Regexp
	replaceFunc replaceFunc
}

type replaceFunc func(string) string

func newRule(regStr string, replaceFuncs ...replaceFunc) *rule {
	return &rule{
		reg:         regexp.MustCompile(regStr),
		replaceFunc: composeAll(replaceFuncs...),
	}
}

var (
	strongAndEm = newRule(`\*\*\*.*\*\*\*`, addTag("em"), addTag("strong"))
	strong      = newRule(`(?<=\*\*).*(?=\*\*)`, addTag("strong"))
)

func addTag(tag string) replaceFunc {
	return func(input string) string {
		return fmt.Sprintf("<%s>%s</%s>", tag, input, tag)
	}
}
