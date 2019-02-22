package main

import (
	"flag"
	"log"
	"os"
	"path"
)

func main() {

	devMode := flag.Bool("serve", false, "serve site")
	port := flag.Int("port", 10066, "serve site at port")
	flag.Parse()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if flag.NArg() > 0 {
		wd = path.Join(wd, flag.Args()[0])
	}

	newTask(wd, *devMode, *port).Run()
}
