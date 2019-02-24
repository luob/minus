package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

type task struct {
	fileName       string
	targetFileName string
	tpl            *template.Template
	data           interface{}
}

func newtask(fileName, targetFileName string, tpl *template.Template, data interface{}) *task {
	return &task{
		fileName:       fileName,
		targetFileName: targetFileName,
		tpl:            tpl,
		data:           data,
	}
}

func newArticletask(fileName, targetFileName string, tpl *template.Template) *task {
	return newtask(fileName, targetFileName, tpl, &struct {
		Title string
		Date  *time.Time
	}{
		Title: fileName,
		Date:  fileName,
	})
}

func (p *task) generate() {
	targetFile, err := os.Create(p.targetFileName)
	if err != nil {
		log.Fatal(err)
	}
	p.tpl.Execute(targetFile, p.data)
	defer targetFile.Close()
}
