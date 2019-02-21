package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	file, err := ioutil.ReadFile("posts/post2.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(markdownToHTML(string(file[:])))
	// var (
	// 	wd  string // work dir
	// 	err error
	// )
	// _ = wd
	// args := os.Args
	// if len(args) < 2 {
	// 	wd, err = os.Getwd()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// } else {
	// 	wd, err = filepath.Abs(os.Args[1])
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

}
