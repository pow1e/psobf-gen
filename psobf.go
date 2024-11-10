package main

import (
	"github.com/pow1e/psobf-gen/cmd"
	_ "github.com/pow1e/psobf-gen/cmd/gen"
	_ "github.com/pow1e/psobf-gen/cmd/obf"
)

func main() {
	cmd.Execute()
}
