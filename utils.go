package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func trimExt(fileName string) string {
	ext := filepath.Ext(fileName)
	return strings.TrimSuffix(fileName, ext)
}

func getFileInfos(dirName string) []os.FileInfo {
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	return fileInfos
}

// func getCurrentDir(workDir string, categories []string) string {
// 	currentDir := workDir
// 	categoriesNum := len(categories)
// 	if categoriesNum != 0 {
// 		currentDir = path.Join(workDir, path.Join(categories...))
// 	}
// 	return currentDir
// }
