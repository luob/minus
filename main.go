package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {

	// get work dir
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		arg := os.Args[1]
		if path.IsAbs(arg) {
			workDir = arg
		} else {
			workDir = path.Join(workDir, flag.Args()[0])
		}
	}

	postsDirName := path.Join(workDir, "posts")
	targetDirName := path.Join(workDir, "public")

	for postList {

	}
	// load templates
	tpls := loadTemplates(path.Join(workDir, "template"))

	// readPostInfo
	postFileInfos, err := ioutil.ReadDir(postsDirName)
	if err != nil {
		log.Fatal(err)
	}

	// prepare target directory
	emptyDir(targetDirName)

	// run(wd, CPUnum, *devMode, *port)
}

func loadTemplates(dirName string) map[string]*template.Template {
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	tpls := make(map[string]*template.Template, 0)
	for _, fileInfo := range fileInfos {
		name := fileInfo.Name()
		ext := filepath.Ext(name)
		tplName := strings.TrimSuffix(name, ext)
		tplFileName := path.Join(dirName, name)
		tpl, err := template.ParseFiles(tplFileName)
		if err != nil {
			log.Fatal(err)
		}
		tpls[tplName] = tpl
	}
	return tpls
}

func emptyDir(dirName string) {
	err := os.RemoveAll(dirName)
	if err != nil {
		log.Fatal(err)
	}
	os.Mkdir(dirName, os.ModePerm)
}
