package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

// PageData is
type PageData struct {
	UserConfig   *UserConfig
	PostInfoList []*PostInfo
	Content      string
}

// UserConfig is
type UserConfig struct {
	Author string
}

// PostInfo is
type PostInfo struct {
	Title string
	Date  string

	filename       string
	targetFileName string
}

func run(workDir string, limit int, devMode bool, port int) {
	// join path
	userConfigFileName := path.Join(workDir, "config.json")
	indexTplFileName := path.Join(workDir, "template", "index.html")
	postTplFileName := path.Join(workDir, "template", "post.html")
	postDirName := path.Join(workDir, "posts")
	targetDirName := path.Join(workDir, "public")
	targetIndexFileName := path.Join(targetDirName, "index.html")

	err := os.RemoveAll(targetDirName)
	if err != nil {
		log.Println(err)
	}
	os.Mkdir(targetDirName, os.ModePerm)

	// load template
	indexTpl, err := template.ParseFiles(indexTplFileName)
	if err != nil {
		log.Fatal(err)
	}
	postTpl, err := template.ParseFiles(postTplFileName)
	if err != nil {
		log.Fatal(err)
	}

	// load user config
	userConfigFile, err := ioutil.ReadFile(userConfigFileName)
	if err != nil {
		log.Println(err)
	}
	userConfig := &UserConfig{}
	err = json.Unmarshal(userConfigFile, userConfig)
	if err != nil {
		log.Println(err)
	}

	// get postList
	postFileInfos, err := ioutil.ReadDir(postDirName)
	if err != nil {
		log.Println(err)
	}
	postInfoList := []*PostInfo{}
	for _, postFileInfo := range postFileInfos {
		title, date := extractFromFileName(postFileInfo.Name())
		log.Println(title)
		log.Println(date)
		if title == "" {
			continue
		}
		postInfoList = append(postInfoList, &PostInfo{
			filename:       path.Join(postDirName, postFileInfo.Name()),
			targetFileName: path.Join(targetDirName, title+".html"),
			Title:          title,
			Date:           date,
		})
	}
	log.Println(postInfoList)

	renderIndexPage(targetIndexFileName, userConfig, postInfoList, indexTpl)
	for _, postInfo := range postInfoList {
		renderPostPage(postInfo.filename, postInfo.targetFileName, userConfig, postInfoList, postTpl)
	}
}

func renderIndexPage(targetIndexFileName string, userConfig *UserConfig, postInfoList []*PostInfo, tpl *template.Template) {
	targetFile, err := os.Create(targetIndexFileName)
	if err != nil {
		panic(err)
	}
	data := &PageData{
		UserConfig:   userConfig,
		PostInfoList: postInfoList,
		Content:      "",
	}
	tpl.Execute(targetFile, data)
	defer targetFile.Close()
}

func renderPostPage(fileName, targetFileName string, userConfig *UserConfig, postInfoList []*PostInfo, tpl *template.Template) {

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
		UserConfig:   userConfig,
		PostInfoList: postInfoList,
		Content:      content,
	}
	tpl.Execute(targetFile, data)
	defer targetFile.Close()
}

func extractFromFileName(fileName string) (title, date string) {
	dateReg := regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2}`)
	date = dateReg.FindString(fileName)
	title = strings.TrimPrefix(fileName, date+"-")
	title = strings.TrimSuffix(title, ".md")
	return
}
