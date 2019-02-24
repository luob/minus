package main

import (
	"errors"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var dateReg = regexp.MustCompile(`^\d{4}-\d{1,2}-\d{1,2}`)

func extractArticle(fileName string) (string, *time.Time, error) {
	dateString := dateReg.FindString(fileName)
	if dateString == "" {
		return "", nil, errors.New("irregular file name ")
	}
	date, err := time.Parse("2006-02-01", dateString)
	if err != nil {
		return "", nil, errors.New("irregular file name")
	}
	title := trimExt(strings.TrimPrefix(fileName, dateString+"-"))
	return title, &date, nil
}

func trimExt(fileName string) string {
	ext := filepath.Ext(fileName)
	return strings.TrimSuffix(fileName, ext)
}
