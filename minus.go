package main

import (
	"log"
	"os"
	"path"
	"text/template"
)

// Minus is
type Minus struct {
	workDir  string
	template *template.Template
	parser   *Parser
}

// New returns a new  Minus instance.
func New(wd string, p *Parser, t *template.Template) *Minus {
	m := &Minus{
		workDir: wd,
		parser:  p,
	}
	m.initTemplate()
	m.prepareTargetDir()
	return m
}

// Run is
func (m *Minus) Run() {

}

// SetTemplate is
func (m *Minus) useTemplate(t *template.Template) {
	m.template = t
}

// SetParser is
func (m *Minus) SetParser(p *Parser) {
	m.parser = p
}

func (m *Minus) initTemplate() {
	m.template = template
}

func (m *Minus) prepareTargetDir() {
	targetDir := path.Join(m.workDir, "target")
	err := os.RemoveAll(targetDir)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(targetDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
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
