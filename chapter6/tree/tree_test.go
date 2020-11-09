package tree

import (
	"math/rand"
	"testing"
	"time"
	"golang.org/x/tour/tree"
)

func TestSameIdentical(t *testing.T) {
	t1 := tree.New(1)
	t2 := tree.New(1)
	same := Same(t1, t2)
	if !same {
		t.Fatalf(`Same(t1, t2) is false for identical trees, want true`)
	}
}

func TestSameSizeDifferent(t *testing.T) {
	t1 := tree.New(1)
	t2 := tree.New(2)
	same := Same(t1, t2)
	if same {
		t.Fatalf(`Same(t1, t2) is true for same size different value trees, want false`)
	}
}

func TestSameLeftTreeBigger(t *testing.T) {
	t1 := tree.New(1)
	Insert(t1, 11)
	t2 := tree.New(1)
	same := Same(t1, t2)
	if same {
		t.Fatalf(`Same(t1, t2) is true for different value trees, left tree bigger, want false`)
	}
}

func TestSameRightTreeBigger(t *testing.T) {
	t1 := tree.New(1)
	Insert(t1, 11)
	t2 := tree.New(1)
	same := Same(t2, t1)
	if same {
		t.Fatalf(`Same(t1, t2) is true for different value trees, left tree bigger, want false`)
	}
}

func TestSameBigTrees(t *testing.T) {
	t1 := tree.New(1)
	t2 := tree.New(1)
	for i := 11; i < 10000; i++ {
		Insert(t1, i)
		Insert(t2, i)
	}
	same := Same(t2, t1)
	if !same {
		t.Fatalf(`Same(t1, t2) is false for two identical big trees, want true`)
	}
}

func TestSameBigRandomTrees(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	t1 := tree.New(1)
	t2 := tree.New(2)
	for i := 11; i < 100000; i++ {
		Insert(t1, rand.Int())
		Insert(t2, rand.Int())
	}
	same := Same(t2, t1)
	if same {
		t.Fatalf(`Same(t1, t2) is true for two different big trees, want false`)
	}
}