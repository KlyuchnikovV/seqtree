package seqtree

import (
	"fmt"

	"github.com/KlyuchnikovV/seqtree/node"
	"github.com/KlyuchnikovV/stack"
)

type SequentialAVLTree struct {
	useNew bool
	root   *node.Node
	size   int
}

func New(data interface{}, useNew bool) *SequentialAVLTree {
	return &SequentialAVLTree{
		root: node.New(data),
		size: 1,
	}
}

func (tree *SequentialAVLTree) Insert(position int, data interface{}) error {
	if position > tree.size || position < 0 {
		return fmt.Errorf("position \"%d\" is out of tree range: [%d - %d]", position, 0, tree.size+1)
	}

	tree.size++

	if tree.root == nil {
		tree.root = node.New(data)
		return nil
	}

	var (
		currentNode     = tree.root
		nodeStack       = stack.New(currentNode.Height())
		currentPosition = currentNode.Position(-1, false)
	)

	// TODO: rework
	// loop:
	for {
		currentNode.SetNumberOfChildren(currentNode.NumberOfChildren() + 1)
		switch {
		case currentPosition >= position && currentNode.HasLeft():
			nodeStack.Push(currentNode)
			currentNode = currentNode.LeftChild()
			currentPosition = currentNode.Position(currentPosition, true)
			continue
		case currentPosition < position && currentNode.HasRight():
			nodeStack.Push(currentNode)
			currentNode = currentNode.RightChild()
			currentPosition = currentNode.Position(currentPosition, false)
			continue
		case currentPosition >= position && !currentNode.HasLeft():
			currentNode.SetLeftChild(node.New(data))
			// break loop
		case currentPosition < position && !currentNode.HasRight():
			currentNode.SetRightChild(node.New(data))
			// break loop
		}
		break
	}

	return tree.Balance()
}

func (tree *SequentialAVLTree) GetNode(position int) *node.Node {
	if tree.size == 0 || position < 0 || position > tree.size {
		return nil
	}

	var (
		currentNode     = tree.root
		currentPosition = currentNode.Position(-1, false)
	)

	for currentNode != nil {
		switch {
		case position < currentPosition:
			currentNode = (*currentNode).LeftChild()
			currentPosition = currentNode.Position(currentPosition, true)
			continue
		case position > currentPosition:
			currentNode = (*currentNode).RightChild()
			currentPosition = currentNode.Position(currentPosition, false)
			continue
		}
		break
	}
	return currentNode
}

func (tree *SequentialAVLTree) Find(position int) (interface{}, bool) {
	var node = tree.GetNode(position)
	if node == nil {
		return nil, false
	}
	return node.Data(), true
}

func (tree *SequentialAVLTree) Size() int {
	return tree.size
}

// func (tree *SequentialAVLTree) Delete(position int) (interface{}, bool) {
// 	if position >= tree.size || position < 0 {
// 		return nil, false
// 	}

// 	var (
// 		currentNode     = tree.root
// 		nodeStack       = stack.New(currentNode.Height())
// 		currentPosition = currentNode.Position(-1, false)
// 	)

// 	for currentPosition != position {
// 		currentNode.SetNumberOfChildren(currentNode.NumberOfChildren() - 1)
// 		nodeStack.Push(currentNode)
// 		if currentPosition > position {
// 			currentNode = currentNode.LeftChild()
// 			currentPosition = currentNode.Position(currentPosition, true)
// 		} else if currentPosition < position {
// 			currentNode = currentNode.RightChild()
// 			currentPosition = currentNode.Position(currentPosition, false)
// 		}
// 	}

// 	var result = currentNode.Data()

// 	if currentNode.HasRight() {

// 		var (
// 			parentNode    = currentNode
// 			replacingNode = currentNode.RightChild()
// 		)
// 		for replacingNode.HasLeft() {
// 			parentNode = replacingNode
// 			replacingNode.SetNumberOfChildren(replacingNode.NumberOfChildren() - 1)
// 			replacingNode = replacingNode.LeftChild()
// 		}
// 		if parentNode != currentNode {
// 			parentNode.SetLeftChild(nil)
// 		}
// 		replacingNode.SetLeftChild(currentNode.LeftChild())

// 		if currentNode.RightChild() != replacingNode {
// 			replacingNode.SetRightChild(currentNode.RightChild())
// 		}
// 		*currentNode = *replacingNode
// 		currentNode.FixNumberOfChildren()
// 	} else if !currentNode.HasLeft() {
// 		v, ok := nodeStack.Peek()
// 		if !ok {
// 			tree.root = nil
// 		} else if currentNode.IsLeftOf(*v.(*node.Node)) {
// 			v.(*node.Node).SetLeftChild(nil)
// 		} else {
// 			v.(*node.Node).SetRightChild(nil)
// 		}
// 	} else {
// 		*currentNode = *currentNode.LeftChild()
// 	}

// 	for v, ok := nodeStack.Pop(); ok; v, ok = nodeStack.Pop() {
// 		v.(*node.Node).BalanceTree()
// 	}

// 	tree.size--
// 	return result, true
// }

// func (tree *SequentialAVLTree) ToList() []interface{} {
// 	return tree.root.toList()
// }

// func (tree *SequentialAVLTree) Visualize() {
// 	if tree.root == nil {
// 		return
// 	}
// 	tree.root.visualizeNodeSubtree(0, tree.root.height())
// }

func (tree *SequentialAVLTree) Balance() error {
	if tree.root == nil {
		return nil
	}
	if !tree.root.HasLeft() && !tree.root.HasRight() {
		return nil
	}

	var (
		root          = tree.root
		leftChildren  = root.LeftChild().NumberOfChildren()
		rightChildren = root.RightChild().NumberOfChildren()
		position      = root.Position(-1, false)
	)

	if leftChildren == rightChildren {
		return nil
	}

	if leftChildren > rightChildren {
		// удалить самую правую из левого
		var prev = root.ExctactPrev()
		tree.size--
		// поставить рутом
		tree.root = prev
		// if tree.root.HasLeft() {
		// 	panic("left")
		// }
		if tree.root != root.LeftChild() {
			tree.root.SetLeftChild(root.LeftChild())
		}
		if tree.root.HasRight() {
			// предыдущая для корня нода не может иметь правого
			panic("right")
		}
		root.SetLeftChild(nil)
		root.SetRightChild(nil)
		// вставить в самое левое правого
		return tree.Insert(position+1, root.Data())
	} else {
		// удалить самую левую из правого
		var prev = root.ExctactNext()
		tree.size--
		// поставить рутом
		tree.root = prev
		if tree.root.HasLeft() {
			// предыдущая для корня нода не может иметь левого
			panic("left")
		}
		tree.root.SetLeftChild(root.LeftChild())
		// if tree.root.HasRight() {
		// 	panic("right")
		// }
		if tree.root != root.RightChild() {
			tree.root.SetRightChild(root.RightChild())
		}
		root.SetLeftChild(nil)
		root.SetRightChild(nil)
		// вставить в самое правое левого
		if position == 0 {
			position = 1
		}
		return tree.Insert(position-1, root.Data())
	}
}
