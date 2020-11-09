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
	if !Same(t1, t2) {
		t.Fatalf(`Same(t1, t2) is false for identical trees, want true`)
	}
}

func TestSameObject(t *testing.T) {
	t1 := tree.New(1)
	if !Same(t1, t1) {
		t.Fatalf(`Same(t1, t2) is false for same tree objects, want true`)
	}
}

func TestSameNilTrees(t *testing.T) {
	t1 := tree.New(1)
	if Same(nil, t1) {
		t.Fatalf(`Same(t1, t2) is true for left nil tree, want false`)
	}
	if Same(t1, nil) {
		t.Fatalf(`Same(t1, t2) is true for right nil tree, want false`)
	}
	if !Same(nil, nil) {
		t.Fatalf(`Same(t1, t2) is false for nil trees, want true`)
	}
}

func TestSameSizeDifferent(t *testing.T) {
	t1 := tree.New(1)
	t2 := tree.New(2)
	if Same(t1, t2) {
		t.Fatalf(`Same(t1, t2) is true for same size different value trees, want false`)
	}
}

func TestSameLeftTreeBigger(t *testing.T) {
	t1 := tree.New(1)
	Insert(t1, 11)
	t2 := tree.New(1)
	if Same(t1, t2) {
		t.Fatalf(`Same(t1, t2) is true for different value trees, left tree bigger, want false`)
	}
}

func TestSameRightTreeBigger(t *testing.T) {
	t1 := tree.New(1)
	Insert(t1, 11)
	t2 := tree.New(1)
	if Same(t2, t1) {
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
	if !Same(t2, t1) {
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
	if Same(t2, t1) {
		t.Fatalf(`Same(t1, t2) is true for two different big trees, want false`)
	}
}