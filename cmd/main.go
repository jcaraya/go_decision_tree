package main

import (
	"fmt"
	"go_decision_tree/pkg/csvreader"
	"go_decision_tree/pkg/dataset"
	"go_decision_tree/pkg/parser"
	"log"
)

func main() {
	filePath := "./data/iris.data"
	lines, err := csvreader.ReadCsv(filePath)
	if err != nil {
		log.Printf("failed reading file: %s", filePath)
		return
	}

	x, y, classes, err := parser.ParseCsvLines(lines)
	if err != nil {
		log.Println("failed parsing the dataset lines with error: ", err)
		return
	}

	fmt.Println("Len:", len(x))
	x, y = dataset.ShuffleDataset(x, y)

	// Partition the data
	xTrain, yTrain := x[:90], y[:90]
	trainDataset := dataset.NewDataset(xTrain, yTrain)

	xValidation, yValidation := x[90:120], y[90:120]
	_ = xValidation
	_ = yValidation

	xTest, yTest := x[120:150], y[120:150]
	_ = xTest
	_ = yTest

	fmt.Println(x[:10])
	fmt.Println(y[:10])
	fmt.Println(classes)
}
