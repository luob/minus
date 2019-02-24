package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type generator struct {
	workDir   string
	targetDir string
	tpls      map[string]*template.Template
	articles  []*article
}

type article struct {
	title      string
	date       *time.Time
	categories []string
}

func newGenerator(workDir string) *generator {
	targetDir := path.Join(workDir, "target")
	tpls := loadTemplates(workDir)
	articles := loadArticles(workDir, []string{"article"})
	return &generator{
		workDir:   workDir,
		targetDir: targetDir,
		tpls:      tpls,
		articles:  articles,
	}
}

func (g *generator) generate() {
	g.prepareTargetDir()
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

func loadTemplates(templateDir string) map[string]*template.Template {
	fileInfos, err := ioutil.ReadDir(templateDir)
	if err != nil {
		log.Fatal(err)
	}
	tpls := make(map[string]*template.Template)
	for _, fileInfo := range fileInfos {
		name := fileInfo.Name()
		ext := filepath.Ext(name)
		tplName := strings.TrimSuffix(name, ext)
		tplFileName := path.Join(templateDir, name)
		tpl, err := template.ParseFiles(tplFileName)
		if err != nil {
			log.Fatal(err)
		}
		tpls[tplName] = tpl
	}
	return tpls
}

func loadArticles(currentDir string, categories []string) []*article {
	articles := make([]*article, 0)
	fileInfos, err := ioutil.ReadDir(currentDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range fileInfos {
		name := fileInfo.Name()
		nameAbs := path.Join(currentDir, name)
		if fileInfo.IsDir() {
			articles = append(articles, loadArticles(nameAbs, append(categories, name))...)
		} else if filepath.Ext(name) == ".md" {
			title, date, err := extractArticle(name)
			if err != nil {
				continue
			}
			articles = append(articles, &article{
				title: title,
				date:  date,
			})
		}
	}
	return articles
}

func (g *generator) generateArticles() {
	// for _, article := range g.articles {

	// }
}

func (g *generator) generateIndex() {

}
