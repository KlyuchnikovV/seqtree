package seqtree

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
)

func TestSequentialAILTree_SequentialInsert(t *testing.T) {
	t.Parallel()
	tree := new(SequentialAVLTree)
	for i := 0; i < 100000; i++ {
		assert.NoError(t, tree.Insert(i, i))
	}
	assert.Equal(t, 100000, tree.Size())
}

func TestSequentialAVLTree_RandomInsert(t *testing.T) {
	t.Parallel()
	tree := new(SequentialAVLTree)
	for i := 0; i < 100000; i++ {
		r := rand.Intn(tree.Size() + 1)
		assert.NoError(t, tree.Insert(r, r))
	}
	assert.Equal(t, 100000, tree.Size())
}

func TestSequentialAVLTree_ToList(t *testing.T) {
	t.Parallel()
	tree := new(SequentialAVLTree)

	for i := 0; i < 8; i++ {
		assert.NoError(t, tree.Insert(i, i))
		assert.Equal(t, i+1, tree.size)
		assert.Equal(t, i+1, tree.root.numberOfChildren+1)
	}

	lines := tree.ToList()
	for i, line := range lines {
		assert.Equal(t, i, line.(int))
	}
}

func TestSequentialAVLTree_Delete(t *testing.T) {
	t.Parallel()
	tree := new(SequentialAVLTree)
	values := make(map[int]bool)

	for i := 0; i < 10; i++ {
		assert.NoError(t, tree.Insert(i, i))
		values[i] = false
	}

	for tree.Size() > 0 {
		data, ok := tree.Delete(rand.Intn(tree.Size()))
		assert.True(t, ok)
		assert.False(t, values[data.(int)])
		values[data.(int)] = true
	}
	for _, value := range values {
		assert.True(t, value)
	}
}

func TestLog2(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000000; i++ {
		assert.Equal(t, int(math.Log2(float64(i))), log2(i), "failed on %d", i)
	}
}

func BenchmarkSequentialInsertingIntoTree(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		tree := new(SequentialAVLTree)
		for j := 0; j < 100000; j++ {
			_ = tree.Insert(j, j)
		}
	}
}

func BenchmarkRandomInsertingIntoTree(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		tree := new(SequentialAVLTree)
		for j := 0; j < 100000; j++ {
			_ = tree.Insert(j, rand.Intn(tree.Size()+1))
		}
	}
}

func BenchmarkRandomInsertingDeleting(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		tree := new(SequentialAVLTree)
		for j := 0; j < 100000; j++ {
			if rand.Intn(1) == 1 || tree.size == 0 {
				_ = tree.Insert(j, rand.Intn(tree.Size()+1))
			} else {
				tree.Delete(rand.Intn(tree.Size()))
			}
		}
	}
}

func BenchmarkSequentialInsertSlice(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		var slice []interface{}
		for j := 0; j < 100000; j++ {
			slice = append(slice, j)
		}
	}
}

func BenchmarkRandomInsertSlice(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		var slice []interface{}
		for j := 0; j < 100000; j++ {
			pos := rand.Intn(len(slice) + 1)
			slice = append(append(slice[:pos], j), slice[pos:]...)
		}
	}
}
