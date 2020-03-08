package main

import "github.com/gocookbook/ch1/templates"

func main() {
	if err := templates.RunTemplate(); err != nil {
		panic(err)
	}
	if err := templates.InitTemplate(); err != nil {
		panic(err)
	}
	if err := templates.HTMLDifferences(); err != nil {
		panic(err)
	}
}
