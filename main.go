package main

import (
	"go/token"
	"go/types"
	"log"
	"strconv"
)

func main() {
	println("result: %v", calc("15*5"))
}

func calc(formula string) int {
	answer, err := types.Eval(token.NewFileSet(), types.NewPackage("main", "calc"), token.NoPos, formula)
	if err != nil {
		log.Fatal(err)
	}
	intAnswer, err := strconv.Atoi(answer.Value.ExactString())
	if err != nil {
		log.Fatal(err)
	}
	return intAnswer
}