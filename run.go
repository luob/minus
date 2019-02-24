package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

// PageData is
type PageData struct {
	PostInfoList []*PostInfo
	ContentFrom  string
	Content      template.HTML
}

// PostInfo is
type PostInfo struct {
	Title string
	Date  string

	filename       string
	targetFileName string
}

// func run(workDir string, limit int, devMode bool, port int) {
// 	// join path
// 	indexTplFileName := path.Join(workDir, "template", "index.html")
// 	postTplFileName := path.Join(workDir, "template", "post.html")
// 	postDirName := path.Join(workDir, "posts")
// 	targetDirName := path.Join(workDir, "public")
// 	targetIndexFileName := path.Join(targetDirName, "index.html")

// 	refreshTargetDir(targetDirName)

// 	// load template
// 	indexTpl := mustParseTemplate(indexTplFileName)
// 	postTpl := mustParseTemplate(postTplFileName)

// 	// get postList
// 	postFileInfos, err := ioutil.ReadDir(postDirName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	postInfoList := []*PostInfo{}
// 	for _, postFileInfo := range postFileInfos {
// 		title, date := extractPostInfo(postFileInfo.Name())
// 		if title == "" {
// 			continue
// 		}
// 		postInfoList = append(postInfoList, &PostInfo{
// 			filename:       path.Join(postDirName, postFileInfo.Name()),
// 			targetFileName: path.Join(targetDirName, title+".html"),
// 			Title:          title,
// 			Date:           date,
// 		})
// 	}

// 	// render

// 	renderIndexPage(targetIndexFileName, postInfoList, indexTpl)
// 	for _, postInfo := range postInfoList {
// 		renderPostPage(postInfo.filename, postInfo.targetFileName, postInfoList, postTpl)
// 	}
// }

func renderIndexPage(targetIndexFileName string, postInfoList []*PostInfo, tpl *template.Template) {
	targetFile, err := os.Create(targetIndexFileName)
	if err != nil {
		panic(err)
	}
	data := &PageData{
		PostInfoList: postInfoList,
	}
	tpl.Execute(targetFile, data)
	defer targetFile.Close()
}

func renderPostPage(fileName, targetFileName string, postInfoList []*PostInfo, tpl *template.Template) {

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}

	content := markdownToHTML(string(file[:]))

	targetFile, err := os.Create(targetFileName)
	if err != nil {
		log.Println(err)
	}

	data := &PageData{
		PostInfoList: postInfoList,
		Content:      content,
	}
	tpl.Execute(targetFile, data)
	defer targetFile.Close()
}
