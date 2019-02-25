package main

import (
	"errors"
	"log"
	"regexp"
	"time"
)

// Article is
type Article struct {
	Title      string
	Date       time.Time
	Categories []string
}

// func newArticle(title, string, date time.Time, categories []string]) {

// }

// articles is
type articles map[string]Article

var dateReg = regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2}`)

func (a articles) add(fileInfo *fileInfo, categories []string) {
	dateString := dateReg.FindString(fileInfo.Name())
	if dateString == "" {
		log.Println("%s is not an irregular article file", fileInfo.absFileName)
		return
	}

	date, err := time.Parse("2006-02-01", dateString)
	if err != nil {
		log.Println("%s is not an irregular article file", fileInfo.absFileName)
		return
	}

	title := fileInfo.nameWithoutExt()

	a[title] = Article{
		Title:      title,
		Date:       date,
		Categories: categories,
	}

}

func newArticle(fileInfo *fileInfo, categories []string) (*Article, error) {

	dateString := dateReg.FindString(fileInfo.name())
	if dateString == "" {
		return nil, errors.New("irregular article file name ")
	}

	date, err := time.Parse("2006-02-01", dateString)
	if err != nil {
		return nil, errors.New("irregular article file name")
	}

	title := fileInfo.nameWithoutExt()

	return &Article{
		Title:      title,
		Date:       date,
		Categories: categories,
	}, nil
}
