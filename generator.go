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
	title    string
	date     string
	category string
}

func newGenerator(workDir string) *generator {
	return &generator{
		workDir:   workDir,
		targetDir: path.Join(workDir, "target"),
	}
}

func (g *generator) generate() {
	g.prepareTargetDir()
	g.loadTemplates()
	g.generateArticles()
	g.generateIndex()
}

func (g *generator) generateArticles() {
	// readCategoriesInfo
	fileInfos, err := ioutil.ReadDir(g.workDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfos := range fileInfos {
		// if fileInfo
	}
}

func (g *generator) generateIndex() {

}

func (g *generator) loadTemplates() {
	fileInfos, err := ioutil.ReadDir(g.workDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range fileInfos {
		name := fileInfo.Name()
		ext := filepath.Ext(name)
		tplName := strings.TrimSuffix(name, ext)
		tplFileName := path.Join(g.workDir, name)
		tpl, err := template.ParseFiles(tplFileName)
		if err != nil {
			log.Fatal(err)
		}
		g.tpls[tplName] = tpl
	}
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
