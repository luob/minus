package main

import (
	"errors"
	"regexp"
	"strings"
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
type articles map[string]*Article

var dateReg = regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2}`)

func (a *articles) add(absFileName string, categories []string) {
	dateString := dateReg.FindString(fileName)
	if dateString == "" {
		return
	}

	date, err := time.Parse("2006-02-01", dateString)
	if err != nil {
		return
	}

	title := trimExt(strings.TrimPrefix(fileName, dateString+"-"))
	a[title] = &article{}

}

func newArticle(fileName string, categories []string) (*Article, error) {

	dateString := dateReg.FindString(fileName)
	if dateString == "" {
		return nil, errors.New("irregular file name ")
	}

	date, err := time.Parse("2006-02-01", dateString)
	if err != nil {
		return nil, errors.New("irregular file name")
	}

	title := trimExt(strings.TrimPrefix(fileName, dateString+"-"))

	return &Article{
		Title:      title,
		Date:       date,
		Categories: categories,
	}, nil
}
