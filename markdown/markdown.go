package markdown

import (
	"html/template"
	"regexp"
)

// Processor is das
type Processor interface {
	Markdown(string) template.HTML
}

// DefaultProcessor is
type DefaultProcessor struct{}

var _ Processor = &DefaultProcessor{}

// Markdown is
func (p *DefaultProcessor) Markdown(input string) template.HTML {
	return markdownToHTML(input)
}

var (
	splitRegexp = regexp.MustCompile(`\n{2,}`)
	langRegexp  = regexp.MustCompile(`\b.*\b`)
)

// MarkdownToHTML is
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
