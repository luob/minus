package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type generator struct {
	workDir         string
	targetDir       string
	tpls            map[string]*template.Template
	ArticleInfoList []*articleInfo
}

type articleInfo struct {
	title      string
	date       string
	categories []string
}

func newGenerator(workDir string) *generator {
	return &generator{
		workDir:         workDir,
		targetDir:       path.Join(workDir, "target"),
		tpls:            loadTemplates(workDir),
		ArticleInfoList: loadArticles(workDir, "article"),
	}
}

func (g *generator) generate() {
	g.prepareTargetDir()
	g.loadTemplates()
	g.loadArticles()
	g.generateArticles()
	g.generateIndex()
}

func (g *generator) prepareTargetDir() {
	err := os.RemoveAll(g.targetDir)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(g.targetDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func loadTemplates(workDir) map[string]*template.Template {
	fileInfos, err := ioutil.ReadDir(g.workDir)
	if err != nil {
		log.Fatal(err)
	}
	tpls := make(map[string]*template.Template)
	for _, fileInfo := range fileInfos {
		name := fileInfo.Name()
		ext := filepath.Ext(name)
		tplName := strings.TrimSuffix(name, ext)
		tplFileName := path.Join(g.workDir, name)
		tpl, err := template.ParseFiles(tplFileName)
		if err != nil {
			log.Fatal(err)
		}
		tpls[tplName] = tpl
	}
	return tpls
}

func loadArticles(currentDir string, categories []string) {
	// readCategoriesInfo
	fileInfos, err := ioutil.ReadDir(currentDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {

		}
	}
}

func (g *generator) generateIndex() {

}
