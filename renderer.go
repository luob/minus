package main

type Renderer interface {
	render(input string) (output string)
}
