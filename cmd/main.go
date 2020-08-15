package main

import (
	"fmt"
	"go_decision_tree/pkg/csvreader"
	"go_decision_tree/pkg/dataset"
)

func main() {
	lines, err := csvreader.ReadCsv("./data/iris.data")
	if err != nil {
		panic(err)
	}

	x, y, classes, err := dataset.ParseCsvLines(lines)
	fmt.Println(x[:10])
	fmt.Println(y[:10])
	fmt.Println(classes)
}
