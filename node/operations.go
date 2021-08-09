package node

func (n *Node) Insert(position int, data interface{}) {
	defer func() {
		n.numberOfChildren++
		n.Balance()
	}()

	var currentPosition = n.Position()
	if currentPosition >= position {
		if !n.HasLeft() {
			n.leftChild = New(data)
			return
		}
		n.leftChild.Insert(position, data)
	} else {
		if !n.HasRight() {
			n.rightChild = New(data)
			return
		}
		n.rightChild.Insert(position-currentPosition, data)
	}
}

func (n *Node) GetNode(position int) *Node {
	var currentPosition = n.Position()
	if currentPosition == position {
		return n
	}

	if currentPosition >= position {
		if !n.HasLeft() {
			return nil
		}
		return n.leftChild.GetNode(position)
	} else {
		if !n.HasRight() {
			return nil
		}
		return n.rightChild.GetNode(position - currentPosition)
	}
}

// func (n *Node) Delete(position int) interface{} {
// 	defer func() {
// 		if n != nil {
// 			n.numberOfChildren--
// 			n.Balance()
// 		}
// 	}()

// 	var (
// 		deletedNode *Node
// 		nodePosition = n.Position()
// 	)
// 	if nodePosition > position  && n.HasLeft(){
// 		if n.leftChild.Position() != position {
// 			return n.leftChild.Delete(position)
// 		}
// 		deletedNode = n.leftChild
// 	} else if nodePosition < position && n.HasRight() {
// 		if n.rightChild.Position() != position  {
// 			return n.rightChild.Delete(position-nodePosition)
// 		}
// 		deletedNode = n.rightChild
// 	}

// 	var (
// 		replaceNode = deletedNode.ExctactNext()
// 		result = deletedNode.data
// 	)


// 	return result
// }
