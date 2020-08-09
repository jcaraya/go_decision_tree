package decisiontree

// DecisionFunction function used to determine the next node.
type DecisionFunction func(int) bool

// Node contains the information of a regular or leaf node.
type Node struct {
	isLeaf    bool
	leafValue int

	decisionFunction DecisionFunction
	left             *Node
	right            *Node
}

// NewNodeLeaf initializes the BinaryNode structure
// with its corresponding leafvalue. A leaf node receives
// its leaf value and the left and right nodes remain as nil.
func NewNodeLeaf(leafValue int) *Node {
	return &Node{
		isLeaf:    true,
		leafValue: leafValue,
	}
}

// NewNodeRegular initializes the BinaryNode structure
// with the decision function and the corresponding next nodes.
func NewNodeRegular(function DecisionFunction,
	left *Node, right *Node) *Node {

	node := &Node{
		decisionFunction: function,
		left:             left,
		right:            right,
	}

	return node
}

// IsLeaf returns whether the Node is leaf type or not.
func (b *Node) IsLeaf() bool {
	return b.isLeaf
}

// LeafValue returns the value of a leaf node.
func (b *Node) LeafValue() int {
	return b.leafValue
}

// NextNode receives a value and returns the next corresponding node.
func (b *Node) NextNode(value int) *Node {
	// Validate the desicion function was assigned.
	if b.decisionFunction == nil {
		return nil
	}

	// Use the decision function to determine the next node.
	if !b.decisionFunction(value) {
		return b.left
	}

	return b.right
}
