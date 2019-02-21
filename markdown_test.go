package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func Test_markdownToHTML(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "strongAndEm",
			args: args{
				input: "***dsadasdad***",
			},
			want: "<strong><em>dsadasdad</em></strong>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := markdownToHTML(tt.args.input); got != tt.want {
				t.Errorf("markdownToHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMarkdownToHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, err := ioutil.ReadFile("posts/post1.md")
		if err != nil {
			log.Fatal(err)
		}
		markdownToHTML(string(file[:]))
	}
}

func BenchmarkMarkdownToHTML2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			file, err := ioutil.ReadFile("posts/post1.md")
			if err != nil {
				log.Fatal(err)
			}
			markdownToHTML(string(file[:]))
		}
	})
}
