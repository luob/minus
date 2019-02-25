package markdown

import "regexp"

type rule struct {
	reg         *regexp.Regexp
	replaceFunc replaceFunc
}

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
