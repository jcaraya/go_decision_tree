package dataset

import "math/rand"

// Dataset contains the (x, y) pairs that represent the examples
// of a dataset. It Implements several utilitary functions that
// are used by the decisiontree package.
type Dataset struct {
	x [][]float64
	y []int
}

// NewDataset receive the (x, y) pairs and initializes a new
// dataset struct.
func NewDataset(x [][]float64, y []int) (*Dataset, error) {
	dataset := &Dataset{
		x: x,
		y: y,
	}
	return dataset, nil
}

// IsEmpty validates if the dataset has at least one example.
func (d *Dataset) IsEmpty() bool {
	return len(d.x) == 0
}

// IsAllSameClass validates if the dataset has all examples
// in the same class.
func (d *Dataset) IsAllSameClass() (int, bool) {
	if len(d.y) == 0 {
		return 0, false
	}

	for _, class := range d.y {
		if class != d.y[0] {
			return 0, false
		}
	}

	return d.y[0], true
}

// BestAttribute returns the best attribute to partition the
// data and its value.
func (d *Dataset) BestAttribute() (int, float64) {
	// TODO: Compute entropy... in the mean time we can
	// have just a random tree.
	if d.IsEmpty() {
		return 0, 0.0
	}

	attrIndex := rand.Intn(len(d.x[0]))

	// Compute the mean of the value
	attrSum := 0.0
	for _, x := range d.x {
		attrSum = attrSum + x[attrIndex]
	}
	attrMean := attrSum / float64(len(d.x))

	return attrIndex, attrMean
}

// DecisionFunction returns the function used to partition the
// dataset.
func (d *Dataset) DecisionFunction(attributeIndex int, value float64) func(interface{}) bool {
	return func(xVal interface{}) bool {
		x := xVal.([]float64)
		return x[attributeIndex] <= value
	}
}

// Partition splits the dtaset into two based on the given decision function.
func (d *Dataset) Partition(decisionFunction func(interface{}) bool) (left, right *Dataset) {
	left, right = &Dataset{}, &Dataset{}

	for i, x := range d.x {
		if decisionFunction(x) {
			right.x = append(right.x, d.x[i])
			right.y = append(right.y, d.y[i])
		} else {
			left.x = append(left.x, d.x[i])
			left.y = append(left.y, d.y[i])
		}
	}

	return left, right
}

// MostCommonClass returns the most common class in the
// dataset.
func (d *Dataset) MostCommonClass() int {
	classCounter := initializeClassCounter(d.y)

	var classMode, classCount int
	for class, count := range classCounter {
		if count > classCount {
			classCount = count
			classMode = class
		}
	}

	return classMode
}

// initializeClassCounter initializes a map that is used to count
// the amount of times a class is used.
func initializeClassCounter(y []int) map[int]int {
	classCounter := make(map[int]int)
	for _, class := range y {
		if val, ok := classCounter[class]; !ok {
			classCounter[class] = 1
		} else {
			classCounter[class] = val + 1
		}
	}
	return classCounter
}
