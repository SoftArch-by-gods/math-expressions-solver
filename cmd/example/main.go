package main

import (
	"flag"
	lab2 "github.com/SoftArch-by-gods/math-expressions-solver"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File with expression to compute")
	outputFile      = flag.String("o", "", "File for expression results")
)

func main() {
	flag.Parse()
	handler := &lab2.ComputeHandler{}

	if *outputFile == "" {
		handler.Output = os.Stdout
	} else {
		f, err := os.Create(*outputFile)
		if err != nil {
			panic("file creation error")
		}
		defer f.Close()
		handler.Output = f
	}

	switch {
	case *inputExpression != "" && *inputFile != "":
		panic("choose one source for expression")
	case *inputExpression != "":
		handler.Input = strings.NewReader(*inputExpression)
	case *inputFile != "":
		f, err := os.Open(*inputFile)
		if err != nil {
			panic("file opening error")
		}
		defer f.Close()
		handler.Input = f
	default:
		panic("no expression")
	}

	err := handler.Compute()
	if err != nil {
		panic(err)
	}
}
