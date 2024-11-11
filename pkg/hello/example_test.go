package hello_test

import (
	"testing"

	"github.com/zenless-lab/sterna/pkg/hello"
)

func TestHelloworld(t *testing.T) {
	text := hello.HelloWorld("Zenless")
	if text != "Hello, Zenless" {
		t.Error("Expected 'Hello, Zenless', but got ", text)
	}
}
