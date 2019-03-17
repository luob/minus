package generator

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

// Generator is
type Generator struct {
	workDir   string
	targetDir string
	templates templates
	articles  articles
}

// New is
func New(workDir string) *Generator {
	targetDir := path.Join(workDir, "target")
	return &Generator{
		workDir:   workDir,
		targetDir: targetDir,
		templates: make(templates),
		articles:  make(articles),
	}
}

// Generate is
func (g *Generator) Generate() {
	g.loadFiles(make([]string, 0))
	g.prepareTargetDir()
	// g.generateArticles()
	g.generateIndex()
}

func (g *Generator) loadFiles(categories []string) {
	currentDir := path.Join(g.workDir, path.Join(categories...))
	fileInfos := getFileInfos(currentDir)
	for _, fi := range fileInfos {
		fileInfo := newFileInfo(currentDir, fi)
		if fileInfo.isDir() {
			g.loadFiles(append(categories, fileInfo.name()))
		} else if fileInfo.ext() == ".md" {
			g.articles.add(fileInfo, categories)
		} else if fileInfo.ext() == ".html" {
			g.templates.add(fileInfo)
		}
	}
}

func getFileInfos(dirName string) []os.FileInfo {
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Println("")
	}
	return fileInfos
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

// func (g *Generator) generateArticles() {
// 	for _, article := range g.articles {
// 		// newArticletask(article.)
// 	}
// }

func (g *Generator) generateIndex() {

}
