package main

import "github.com/zenless-lab/sterna/pkg/hello"

func main() {
	name := "Zenless"
	text := hello.HelloWorld(name)
	println(text)
}
