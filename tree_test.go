package seqtree

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequentialAVLTree_SequentialInsert(t *testing.T) {
	t.Parallel()
	tree := new(SequentialAVLTree)
	for i := 1; i <= 100000; i++ {
		assert.NoError(t, tree.Insert(i, i))
	}
	assert.Equal(t, 100000, tree.Size())
	tree.Find(0)
}

func TestSequentialAVLTree_RandomInsert(t *testing.T) {
	t.Parallel()
	tree := new(SequentialAVLTree)
	for i := 1; i <= 100000; i++ {
		r := rand.Intn(tree.Size() + 1)
		assert.NoError(t, tree.Insert(r, r))
	}
	assert.Equal(t, 100000, tree.Size())
}

// func TestSequentialAVLTree_ToList(t *testing.T) {
// 	t.Parallel()
// 	tree := new(SequentialAVLTree)

// 	for i := 1; i <= 8; i++ {
// 		assert.NoError(t, tree.Insert(i, i))
// 		assert.Equal(t, i, tree.size)
// 		assert.Equal(t, i, tree.root.NumberOfChildren()+1)
// 	}

// 	tree.Delete(4)

// 	lines := tree.ToList()
// 	for i, line := range lines {
// 		assert.Equal(t, i, line.(int))
// 	}
// }

// func TestSequentialAVLTree_Delete(t *testing.T) {
// 	t.Parallel()
// 	tree := new(SequentialAVLTree)
// 	values := make(map[int]bool)

// 	for i := 1; i <= 10; i++ {
// 		assert.NoError(t, tree.Insert(i, i))
// 		values[i] = false
// 	}

// 	for tree.Size() > 0 {
// 		data := tree.Delete(rand.Intn(tree.Size()))
// 		assert.False(t, values[data.(int)])
// 		values[data.(int)] = true
// 	}
// 	for _, value := range values {
// 		assert.True(t, value)
// 	}
// }

func BenchmarkSequentialInsertingIntoTree(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		tree := new(SequentialAVLTree)
		for j := 1; j <= 100000; j++ {
			_ = tree.Insert(j, j)
		}
	}
}

func BenchmarkRandomInsertingIntoTree(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		tree := new(SequentialAVLTree)
		for j := 1; j <= 100000; j++ {
			_ = tree.Insert(j, rand.Intn(tree.Size()+1))
		}
	}
}

// func BenchmarkRandomInsertingDeleting(b *testing.B) {
// 	for i := 0; i <= b.N; i++ {
// 		tree := new(SequentialAVLTree)
// 		for j := 1; j <= 100000; j++ {
// 			if rand.Intn(1) == 1 || tree.size == 0 {
// 				_ = tree.Insert(j, rand.Intn(tree.Size()+1))
// 			} else {
// 				tree.Delete(rand.Intn(tree.Size()))
// 			}
// 		}
// 	}
// }

func BenchmarkSequentialInsertSlice(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		var slice []interface{}
		for j := 1; j <= 100000; j++ {
			slice = append(slice, j)
		}
	}
}

func BenchmarkRandomInsertSlice(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		var slice []interface{}
		for j := 1; j <= 100000; j++ {
			pos := rand.Intn(len(slice) + 1)
			slice = append(append(slice[:pos], j), slice[pos:]...)
		}
	}
}
