package markdown

import "strings"

type defaultRenderer struct{}

func newRenderer() *Renderer {
	return &defaultRenderer{}
}

func (r *defaultRender) Render(input string) (output string) {
	return parseArticle(input).toHTML()
}

func parseArticle(input string) (ast *node) {
	ast = newRoot()
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		ast.addChild(input)
	}
}
