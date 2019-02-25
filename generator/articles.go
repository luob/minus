package generator

import (
	"log"
	"regexp"
	"time"
)

// articles is
type articles map[string]Article

// Article is
type Article struct {
	*fileInfo
	Title      string
	Date       time.Time
	Categories []string
}

var dateReg = regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2}`)

func (a articles) add(fileInfo *fileInfo, categories []string) {
	dateString := dateReg.FindString(fileInfo.Name())
	if dateString == "" {
		log.Printf("%s is not an irregular article file", fileInfo.absFileName())
		return
	}

	date, err := time.Parse("2006-02-01", dateString)
	if err != nil {
		log.Printf("%s is not an irregular article file", fileInfo.absFileName())
		return
	}

	title := fileInfo.nameWithoutExt()

	a[title] = Article{
		Title:      title,
		Date:       date,
		Categories: categories,
		fileInfo:   fileInfo,
	}

}

func (a *Article) absFileName() string {
	return a.fileInfo.absFileName()
}

func (a *Article) targetFileName() string {
	return ""
}
