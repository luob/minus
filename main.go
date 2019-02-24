package main

import (
	"log"
	"os"
	"path"
	"runtime"
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

	newGenerator(workDir).generate()
}
