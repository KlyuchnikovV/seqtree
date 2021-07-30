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

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func (n Node) Data() interface{} {
	return n.data
}

func (n Node) LeftChild() *Node {
	return n.leftChild
}

func (n *Node) SetLeftChild(left *Node) {
	n.leftChild = left
}

func (n Node) RightChild() *Node {
	return n.rightChild
}

func (n *Node) SetRightChild(right *Node) {
	n.rightChild = right
}

func (n Node) HasLeft() bool {
	return n.leftChild != nil
}

func (n Node) HasRight() bool {
	return n.rightChild != nil
}

func (n *Node) IsLeftOf(parent Node) bool {
	return parent.leftChild == n
}

func (n *Node) IsRightOf(parent Node) bool {
	return parent.rightChild == n
}

func (n Node) NumberOfChildren() int {
	return n.numberOfChildren
}

func (n *Node) SetNumberOfChildren(num int) {
	n.numberOfChildren = num
}

func (n Node) IsLeaf() bool {
	return n.leftChild == nil && n.rightChild == nil
}

func (n Node) Height() int {
	return 1 + log2(n.numberOfChildren+1)
}

// TODO: rework
func (n Node) Position(prevPosition int, goingLeft bool) int {
	if goingLeft {
		if n.rightChild != nil {
			return prevPosition - n.rightChild.numberOfChildren - 2
		}
		return prevPosition - 1
	}
	if n.leftChild != nil {
		return prevPosition + n.leftChild.numberOfChildren + 2
	}
	return prevPosition + 1
}

// func (n Node) toList() []interface{} {
// 	var result = make([]interface{}, 0, n.numberOfChildren+1)

// 	if n.leftChild != nil {
// 		result = n.leftChild.toList()
// 	}
// 	result = append(result, n.data)
// 	if n.rightChild != nil {
// 		result = append(result, n.rightChild.toList()...)
// 	}
// 	return result
// }

// func (n Node) visualizeNodeSubtree(currentLevel, treeHeight int) {
// 	if n.leftChild != nil {
// 		n.leftChild.visualizeNodeSubtree(currentLevel+1, treeHeight)
// 	}

// 	fmt.Printf("%s%#v%s\n", strings.Repeat("  ", currentLevel), n.data, strings.Repeat("--", treeHeight-currentLevel))

// 	if n.rightChild != nil {
// 		n.rightChild.visualizeNodeSubtree(currentLevel+1, treeHeight)
// 	}
// }
