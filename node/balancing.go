package node

func (n *Node) BalanceTree() {
	switch n.getBalance() {
	case 2:
		if n.rightChild.getBalance() < 0 {
			n.rightChild.rotateRight()
		}
		n.rotateLeft()
	case -2:
		if n.leftChild.getBalance() > 0 {
			n.leftChild.rotateLeft()
		}
		n.rotateRight()
	}
}

func (n *Node) FixNumberOfChildren() {
	if n.leftChild == nil {
		(*n).numberOfChildren = 0
	} else {
		(*n).numberOfChildren = n.leftChild.numberOfChildren + 1
	}

	if n.rightChild != nil {
		(*n).numberOfChildren += n.rightChild.numberOfChildren + 1
	}
}

func (n Node) getBalance() int {
	result := 0
	if n.leftChild != nil {
		result -= n.leftChild.Height()
	}
	if n.rightChild != nil {
		result += n.rightChild.Height()
	}
	return result
}

func (n *Node) rotateRight() {
	if n.leftChild == nil {
		return
	}
	temp := *n

	*n = *n.leftChild
	temp.leftChild = (*n).rightChild
	(*n).rightChild = &temp

	n.rightChild.FixNumberOfChildren()
	n.FixNumberOfChildren()
}

func (n *Node) rotateLeft() {
	if n.rightChild == nil {
		return
	}
	temp := *n

	*n = *n.rightChild
	temp.rightChild = (*n).leftChild
	(*n).leftChild = &temp

	n.leftChild.FixNumberOfChildren()
	n.FixNumberOfChildren()
}
