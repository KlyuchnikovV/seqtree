package node

type Node struct {
	data                  interface{}
	numberOfChildren      int
	leftChild, rightChild *Node
}

func New(data interface{}) *Node {
	return &Node{
		data: data,
	}
}

/* Node's setters */

func (n *Node) SetData(data interface{})    { n.data = data }
func (n *Node) SetLeftChild(left *Node)     { n.leftChild = left }
func (n *Node) SetRightChild(right *Node)   { n.rightChild = right }
func (n *Node) SetNumberOfChildren(num int) { n.numberOfChildren = num }

/* Node's getters */

func (n Node) Data() interface{}           { return n.data }
func (n Node) Height() int                 { return 1 + log2(n.numberOfChildren+1) }
func (n Node) IsLeaf() bool                { return n.leftChild == nil && n.rightChild == nil }
func (n Node) HasLeft() bool               { return n.leftChild != nil }
func (n Node) HasRight() bool              { return n.rightChild != nil }
func (n Node) LeftChild() *Node            { return n.leftChild }
func (n Node) RightChild() *Node           { return n.rightChild }
func (n *Node) IsLeftOf(parent Node) bool  { return parent.leftChild == n }
func (n *Node) IsRightOf(parent Node) bool { return parent.rightChild == n }

func (n *Node) NumberOfChildren() int {
	if n == nil {
		return 0
	}
	return n.numberOfChildren
}

// TODO: rework
func (n Node) Position() int {
	if n.leftChild == nil {
		return 1
	}
	return n.leftChild.numberOfChildren + 2
}

func (n *Node) ExctactPrev() *Node {
	if !n.HasLeft() {
		return nil
	}

	var (
		node   = n.leftChild
		parent = n
	)
	parent.numberOfChildren--
	for node.HasRight() {
		parent = node
		parent.numberOfChildren--
		node = node.rightChild
	}
	if parent == n {
		parent.leftChild = nil
	} else if parent.HasRight() {
		parent.rightChild = nil
	}
	return node
}

func (n *Node) ExctactNext() *Node {
	if !n.HasRight() {
		return nil
	}

	var (
		node   = n.rightChild
		parent = n
	)
	parent.numberOfChildren--
	for node.HasLeft() {
		parent = node
		parent.numberOfChildren--
		node = node.leftChild
	}
	if parent == n {
		parent.rightChild = nil
	} else if parent.HasLeft() {
		parent.leftChild = nil
	}
	return node
}

func (n Node) ToList() []interface{} {
	var result = make([]interface{}, 0, n.numberOfChildren+1)

	if n.leftChild != nil {
		result = n.leftChild.ToList()
	}
	result = append(result, n.data)
	if n.rightChild != nil {
		result = append(result, n.rightChild.ToList()...)
	}
	return result
}

// func (n Node) visualizeNodeSubtree(currentLevel, treeHeight int) {
// 	if n.leftChild != nil {
// 		n.leftChild.visualizeNodeSubtree(currentLevel+1, treeHeight)
// 	}

// 	fmt.Printf("%s%#v%s\n", strings.Repeat("  ", currentLevel), n.data, strings.Repeat("--", treeHeight-currentLevel))

// 	if n.rightChild != nil {
// 		n.rightChild.visualizeNodeSubtree(currentLevel+1, treeHeight)
// 	}
// }
