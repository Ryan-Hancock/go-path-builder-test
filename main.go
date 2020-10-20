package main

import (
	"fmt"

	"github.com/ryan-hancock/path-builder/builder"
)

func main() {
	var p builder.PathParser
	p.Separator = "/"

	path := p.ParseArray([]interface{}{"abc", "d/", nil})
	fmt.Println(path)
}
