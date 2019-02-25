package main

import (
	htmlTemplate "html/template"
	"log"
)

type templates map[string]template

type template struct {
	*htmlTemplate.Template
	used bool
}

func (t templates) add(fileInfo *fileInfo) {
	absFileName := fileInfo.absFileName()
	key := fileInfo.nameWithoutExt()
	tpl, err := htmlTemplate.ParseFiles(absFileName)
	if err != nil {
		log.Printf("%s is not an irregular template file", absFileName)
		return
	}
	t[key] = template{
		Template: tpl,
	}
}

func (t templates) use() {

}
