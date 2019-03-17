package main

import (
	"log"
	"os"
	"path"
	"runtime"
	"text/template"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

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
			workDir = path.Join(workDir, arg)
		}
	}

	// load template
	tplPattern := path.Join(workDir, "template", "*")
	tpl, err := template.ParseGlob(tplPattern)
	if err != nil {
		log.Fatal(err)
	}

	// refresh target dir
	targetDir := path.Join(workDir, "target")
	err = os.RemoveAll(targetDir)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(targetDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	New(workDir, targetDir, Parser)

}
