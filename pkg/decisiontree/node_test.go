package decisiontree

import (
	"testing"
)

// Validates thecorrect initialization of a leaf node.
func TestNode_Leaf(t *testing.T) {
	expected := 10

	// Function under test.
	node := NewNodeLeaf(expected)

	// Validations.
	if !node.IsLeaf() {
		t.Error("expected leaf node")
	}

	if got := node.LeafValue(); got != expected {
		t.Errorf("expected leaf value: %d, got: %d", expected, got)
	}

	if node.NextNode(0) != nil {
		t.Errorf("expected nil next node")
	}
}

// Validates thecorrect initialization of a leaf node.
func TestNode_Regular(t *testing.T) {

	left := &Node{}
	right := &Node{}
	decisionFunction := func(value interface{}) bool {
		return value.(int) > 10
	}

	// Function under test.
	node := NewNodeRegular(decisionFunction, left, right)

	// Validations.
	if node.IsLeaf() {
		t.Error("expected leaf node")
	}

	if node.NextNode(0) != left {
		t.Errorf("obtained incorrect node")
	}

	if node.NextNode(20) != right {
		t.Errorf("obtained incorrect node")
	}
}
