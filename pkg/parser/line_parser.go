package parser

import "strconv"

// Constants realted to the expected components of a CSV line.
const (
	attributeCount   = 4
	categoryPosition = 4
)

// ParseCsvLines receives the lines read from the csv file and returns
// the corresponding features, classes and a mapping of the class to integer.
func ParseCsvLines(csvLines [][]string) ([][]float64, []int, map[string]int, error) {
	classes := make(map[string]int)

	x := make([][]float64, len(csvLines))
	y := make([]int, len(csvLines))

	err := parseLines(x, y, classes, csvLines)
	if err != nil {
		return nil, nil, nil, err
	}

	return x, y, classes, nil
}

// parseLines receives the empty (x, y classMap) and fills the
// corresponding information based on the csvLines.
func parseLines(x [][]float64, y []int, classes map[string]int, csvLines [][]string) error {
	var err error
	classCounter := 0

	for i, line := range csvLines {

		x[i], err = parseAttributes(line)
		if err != nil {
			return err
		}

		y[i] = parseClass(classes, &classCounter, line)
	}

	return nil
}

// parseAttributes returns a list of floats that correspond to the
// attributes or features of the Iris dataset.
func parseAttributes(line []string) ([]float64, error) {
	var err error
	parsedAttributes := make([]float64, attributeCount)

	for i := 0; i < attributeCount; i++ {
		parsedAttributes[i], err = strconv.ParseFloat(line[i], 64)
		if err != nil {
			return nil, err
		}
	}

	return parsedAttributes, nil
}

// parseClass receives a line read from the csv file and determines
// if it contains a new unique category. In that case it proceeds to
// to create a new entry in the classMap. Otherwise, it just returns
// an existing value from the classMap.
func parseClass(classes map[string]int, classCounter *int, line []string) int {
	class := line[categoryPosition]

	if _, ok := classes[class]; !ok {
		classes[class] = *classCounter
		*classCounter++
	}

	return classes[class]
}
