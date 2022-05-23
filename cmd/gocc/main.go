package main

import (
    "github.com/famasoon/gocc"
    "golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
    singlechecker.Main(gocc.Analyzer)
}