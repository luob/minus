package markdown

import (
	"regexp"
)

// Processor is das
type Processor interface {
	Markdown(string) string
}

// DefaultProcessor is
type DefaultProcessor struct{}

var _ Processor = &DefaultProcessor{}

// Markdown is
func (p *DefaultProcessor) Markdown(input string) string {
	return markdownToHTML(input)
}

var (
	splitRegexp = regexp.MustCompile(`\n{2,}`)
	langRegexp  = regexp.MustCompile(`\b.*\b`)
)

// MarkdownToHTML is
func markdownToHTML(input string) string {
	s := ""
	for _, block := range splitRegexp.Split(input, -1) {
		s += parseBlock(block)
	}
	return s
}

func parseBlock(input string) string {
	for _, rule := range rules {
		if rule.reg.MatchString(input) {
			return rule.replaceFunc(input)
		}
	}
	return input
}
