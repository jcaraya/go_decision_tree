package decisiontree

import "go_decision_tree/pkg/dataset"

// Tree TODO
type Tree struct {
	root *Node
}

// Learn TODO
func (t *Tree) Learn(dataset *dataset.Dataset) {
	// Validate that we have at least one example.
	if dataset.IsEmpty() {
		return
	}

	t.root = t.buildTree(dataset, nil)
}

func (t *Tree) buildTree(dataset *dataset.Dataset, parentDataset *dataset.Dataset) *Node {
	if dataset.IsEmpty() {
		return NewNodeLeaf(parentDataset.MostCommonClass())
	} else if class, ok := dataset.IsAllSameClass(); ok {
		return NewNodeLeaf(class)
	} else {
		attrIndex, attrValue := dataset.BestAttribute()
		decisionFunction := dataset.DecisionFunction(attrIndex, attrValue)

		leftDataset, rightDataset := dataset.Partition(decisionFunction)
		leftNode, rightNode := t.buildTree(leftDataset, dataset), t.buildTree(rightDataset, dataset)

		root := NewNodeRegular(decisionFunction, leftNode, rightNode)
		return root
	}
}
