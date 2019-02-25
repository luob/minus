package main

import (
	"log"
	"os"
	"time"
)

type task struct {
	fileName       string
	targetFileName string
	template       *template
	data           interface{}
}

func newtask(fileName, targetFileName string, template *template, data interface{}) *task {
	return &task{
		fileName:       fileName,
		targetFileName: targetFileName,
		template:       template,
		data:           data,
	}
}

func newArticletask(fileName, targetFileName string, tpl *template) *task {
	return newtask(fileName, targetFileName, tpl, &struct {
		Title string
		Date  *time.Time
	}{
		Title: fileName,
		// Date:  fileName,
	})
}

func (p *task) run() {
	targetFile, err := os.Create(p.targetFileName)
	if err != nil {
		log.Fatal(err)
	}
	p.template.Execute(targetFile, p.data)
	defer targetFile.Close()
}
