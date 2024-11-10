package main

import (
	"github.com/pow1e/psobf-gent/cmd"
	_ "github.com/pow1e/psobf-gent/cmd/gen"
	_ "github.com/pow1e/psobf-gent/cmd/obf"
)

func main() {
	cmd.Execute()
}
