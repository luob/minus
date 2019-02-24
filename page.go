package main

import (
	"html/template"
	"log"
	"os"
	"path"
	"strings"
)

type page struct {
	fileName       string
	targetFileName string
	tpl            *template.Template
	data           interface{}
}

func newPage(fileName string, tpl *template.Template, data interface{}) *page {
	targetFileName := strings.TrimSuffix(fileName, path.Ext(fileName)) + ".html"
	return &page{
		fileName:       fileName,
		targetFileName: targetFileName,
		tpl:            tpl,
		data:           data,
	}
}

func newArticlePage(fileName string, tpl *template.Template) *page {
	return newPage(fileName, tpl, &struct {
		Title string
		Date  string
	}{
		Title: fileName,
		Date:  fileName,
	})
}

func (p *page) generate() {
	targetFile, err := os.Create(p.targetFileName)
	if err != nil {
		log.Fatal(err)
	}
	p.tpl.Execute(targetFile, p.data)
	defer targetFile.Close()
}

// ArticlePageData is
type ArticlePageData struct {
	Title    string
	Date     string
	Category string
	Content  string
}

// IndexPageData is
type IndexPageData struct {
	Title   string
	Date    string
	Content string
}
