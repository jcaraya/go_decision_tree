package main

import (
	"fmt"
	"go_decision_tree/pkg/csvreader"
	"go_decision_tree/pkg/parser"
)

func main() {
	lines, err := csvreader.ReadCsv("./data/iris.data")
	if err != nil {
		panic(err)
	}

	x, y, classes, err := parser.ParseCsvLines(lines)
	fmt.Println(x[:10])
	fmt.Println(y[:10])
	fmt.Println(classes)
}
