package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	var (
		wd  string // work dir
		err error
	)
	_ = wd
	args := os.Args
	if len(args) < 2 {
		wd, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		wd, err = filepath.Abs(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}
}
