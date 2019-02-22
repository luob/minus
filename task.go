package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"path"
	"regexp"
)

// Task is
type Task struct {
	userConfig *userConfig
	postInfos  []*postInfo
	indexTpl   *template.Template
	postTpl    *template.Template
}

type userConfig struct {
	Author string
	Index  string
	// TemplateDir string `json:`
}

type indexInfo struct {
	filename       string
	targetFileName string

	template *template.Template
	author   string
	title    string
	date     string
}

type postInfo struct {
	filename       string
	targetFileName string
	template       *template.Template
	author         string
	title          string
	date           string
}

func newTask(workDir string, devMode bool, port int) *Task {

	// join path
	userConfigFileName := path.Join(workDir, "config.json")
	indexTplFileName := path.Join(workDir, "template", "index.html")
	postTplFileName := path.Join(workDir, "template", "post.html")
	postDirName := path.Join(workDir, "posts")
	targetDirName := path.Join(workDir, "public")

	// load template
	indexTpl := template.New(indexTplFileName)
	postTpl := template.New(postTplFileName)

	// load user config
	userConfigFile, err := ioutil.ReadFile(userConfigFileName)
	userConfig := &userConfig{}
	err = json.Unmarshal(userConfigFile, userConfig)
	if err != nil {
		log.Println("config.json parse error")
		log.Println(err)
	}

	// load posts info
	postFileInfos, err := ioutil.ReadDir(postDirName)
	postInfos := []*postInfo{}
	if err != nil {
		log.Println("posts directory not found")
		log.Println(err)
	}
	for _, postFileInfo := range postFileInfos {
		title, date := extractFromFileName(postFileInfo.Name())
		if title == "" {
			continue
		}
		postInfos = append(postInfos, &postInfo{
			filename:       path.Join(postDirName, postFileInfo.Name()),
			targetFileName: path.Join(targetDirName, title+".html"),
			template:       postTpl,
			title:          title,
			date:           date,
		})
	}

	return &Task{
		indexTpl:  indexTpl,
		postTpl:   postTpl,
		postInfos: postInfos,
	}
}

// Run is
func (t *Task) Run() {
	t.processIndex(t.indexTpl, t.postInfos)
	// t.processPosts(t.postTpl, t.postInfoList)
}

func (t *Task) processIndex(indexTpl *template.Template, postInfoList []*postInfo) {

}

func (t *Task) processPosts(indexTpl *template.Template, postInfoList []*postInfo) {

}

// func build(wd string) {
// 	postsDir := path.Join(srcDir, "posts")
// }

// func serve(wd string, port int) {

// }
// func task(srcDir, targetDir string) {
// 	postsDir := path.Join(srcDir, "posts")
// 	files, err := ioutil.ReadDir(postsDir)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, file := range files {
// 		fileName := path.Join(postsDir, file.Name())
// 		processFile(fileName, targetDir)
// 	}
// }

// func processPost(fileName string, targetDir string) {
// 	file, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	content := string(file[:])
// 	page
// }

func extractFromFileName(fileName string) (title, date string) {
	titleReg := regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2}`)
	dateReg := regexp.MustCompile(`-.*\.md$`)
	title = titleReg.FindString(fileName)
	date = dateReg.FindString(fileName)
	return
}
