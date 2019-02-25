package main

import (
	"log"
	"os"
	"path"
)

// Generator is
type Generator struct {
	workDir   string
	targetDir string
	tpls      tpls
	articles  []*Article
}

// NewGenerator is
func NewGenerator(workDir string) *Generator {
	targetDir := path.Join(workDir, "target")
	return &Generator{
		workDir:   workDir,
		targetDir: targetDir,
	}
}

// Generate is
func (g *Generator) Generate() {
	g.loadFiles(make([]string, 0))
	g.prepareTargetDir()
	g.generateArticles()
	g.generateIndex()
}

func (g *Generator) loadFiles(categories []string) {
	currentDir := path.Join(g.workDir, path.Join(categories...))
	fileInfos := getFileInfos(currentDir)
	for _, fileInfo := range fileInfos {
		name := fileInfo.Name()
		if fileInfo.IsDir() {
			g.loadFiles(append(categories, name))
		} else if path.Ext(name) == ".md" {
			article, err := newArticle(name, categories)
			if err != nil {
				continue
			}
			g.articles = append(g.articles, article)
		} else if path.Ext(name) == ".html" {
			g.tpls.add(path.Join(currentDir, name))
		}
	}
}

func (g *Generator) prepareTargetDir() {
	err := os.RemoveAll(g.targetDir)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(g.targetDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Generator) generateArticles() {
	// for _, article := range g.articles {

	// }
}

func (g *Generator) generateIndex() {

}
