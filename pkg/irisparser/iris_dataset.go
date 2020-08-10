package irisparser

import "strconv"

const attributeCount = 4
const categoryPosition = 4

// IrisDataset defines the features and classes that compose the
// iris dataset. It implements several utilitary function that will
// help with the data processing.
type IrisDataset struct {
	x [][]float64 // Features or attributes of the dataset.
	y []int       // Classes.

	categorySet     map[string]int // Set of unique classes.
	categoryCounter int            // Number of categories in the dataset.
}

// NewIrisDataset receives as a parameter the lines obtained through the
// CSV Reader. It parses the raw information into the IrisDataset struct
// which contains usable information of the dataset.
func NewIrisDataset(lines [][]string) (*IrisDataset, error) {

	ids := &IrisDataset{
		x:           make([][]float64, len(lines)),
		y:           make([]int, len(lines)),
		categorySet: make(map[string]int),
	}

	err := ids.parseLines(lines)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

// parseLines receives the lines read from the csv file and fills
// the corresponding information of the IrisDataset struct.
func (ids *IrisDataset) parseLines(lines [][]string) error {
	var err error
	for i, line := range lines {

		ids.x[i], err = ids.parseAttributes(line)
		if err != nil {
			return err
		}

		ids.y[i] = ids.parseCategory(line)
	}

	return nil
}

// parseAttributes returns a list of floats that correspond to the
// attributes or features of the Iris dataset.
func (ids *IrisDataset) parseAttributes(line []string) ([]float64, error) {
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

// parseCategory receives a line read from the csv file and determines
// if it contains a new unique category. In that case it proceeds to
// to create a new entry in the categorySet. Otherwise, it just returns
// an existing value.
func (ids *IrisDataset) parseCategory(line []string) int {
	category := line[categoryPosition]

	if _, ok := ids.categorySet[category]; !ok {
		ids.categorySet[category] = ids.categoryCounter
		ids.categoryCounter++
	}

	return ids.categorySet[category]
}
