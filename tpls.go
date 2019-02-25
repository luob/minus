package main

import (
	"html/template"
	"log"
	"path"
)

type tpls map[string]*template.Template

func (t tpls) add(absFileName string) {
	_, fileName := path.Split(absFileName)
	key := trimExt(fileName)
	tpl, err := template.ParseFiles(absFileName)
	if err != nil {
		log.Fatal(err)
	}
	t[key] = tpl
}

type tpl struct {
	tpl  *template.Template
	used bool
}
