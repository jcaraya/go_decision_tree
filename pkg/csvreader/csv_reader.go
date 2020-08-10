package csvreader

import (
	"encoding/csv"
	"os"
)

// ReadCsv returns the contents of a CSV file. Each line is
// represented as a list of string components.
func ReadCsv(fileName string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
