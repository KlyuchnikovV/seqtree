package seqtree

import (
	"fmt"

	"github.com/KlyuchnikovV/seqtree/node"
)

// TODO: split and concat

type SequentialAVLTree struct {
	root *node.Node
	size int
}

func New(data interface{}, useNew bool) *SequentialAVLTree {
	return &SequentialAVLTree{
		root: node.New(data),
		size: 1,
	}
}

func (tree *SequentialAVLTree) Insert(position int, data interface{}) error {
	if position > tree.size+1 || position <= 0 {
		return fmt.Errorf("position \"%d\" is out of tree range: [%d - %d]", position, 1, tree.size+1)
	}

	defer func() {
		tree.size = tree.root.NumberOfChildren() + 1
	}()

	if tree.root == nil {
		tree.root = node.New(data)
		return nil
	}

	tree.root.Insert(position, data)
	return nil
}

func (tree *SequentialAVLTree) GetNode(position int) *node.Node {
	if tree.size == 0 || position <= 0 || position > tree.size {
		return nil
	}

	return tree.root.GetNode(position)
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

// func (tree *SequentialAVLTree) Delete(position int) interface{} {
// 	if position > tree.size || position <= 0 {
// 		return nil
// 	}

// 	defer func() {
// 		tree.size = tree.root.NumberOfChildren() + 1
// 	}()

// 	if tree.root == nil {
// 		return nil
// 	}
// 	if tree.root.Position() == position {
// 		var (
// 			result  = tree.root.Data()
// 			newRoot = tree.root.ExctactNext()
// 		)
// 	}

// 	return tree.root.Delete(position)
// }

func (tree *SequentialAVLTree) ToList() []interface{} {
	return tree.root.ToList()
}

// func (tree *SequentialAVLTree) Visualize() {
// 	if tree.root == nil {
// 		return
// 	}
// 	tree.root.visualizeNodeSubtree(0, tree.root.height())
// }
