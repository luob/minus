package main

type page struct {
	fileName       string
	targetFileName string
	data           interface{}
}

type ArticlePageData struct {
	Title    string
	Date     string
	Category string
	Content  string
}

type IndexPageData struct {
	Title   string
	Date    string
	Content string
}

func newPage(fileName) {

}
