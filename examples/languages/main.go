package main

import (
	"fmt"

	gocloc "github.com/fniu/pocloc"
)

func main() {
	languages := gocloc.NewDefinedLanguages(false)
	options := gocloc.NewClocOptions()
	paths := []string{
		".",
	}

	processor := gocloc.NewProcessor(languages, options)
	result, err := processor.Analyze(paths)
	if err != nil {
		fmt.Printf("gocloc fail. error: %v\n", err)
		return
	}

	for _, lang := range result.Languages {
		fmt.Println(lang)
	}
	fmt.Println(result.Total)
	fmt.Printf("%+v", result)
}
