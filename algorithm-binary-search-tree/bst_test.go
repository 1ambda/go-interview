package bst

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestBst_Min(t *testing.T) {
	tree1 := &Node{
		value: 1,
	}
	assert.Equal(t, 1, tree1.Min())

	tree2 := &Node{
		value: 5,
		left: &Node{
			value: 3,
			left: &Node{
				value: 1,
			},
		},
	}
	assert.Equal(t, 1, tree2.Min())

	var tree3 *Node = nil
	assert.Equal(t, math.MinInt64, tree3.Min())
}

func TestBst_Max(t *testing.T) {
	tree1 := &Node{
		value: 1,
	}
	assert.Equal(t, 1, tree1.Max())

	tree2 := &Node{
		value: 5,
		right: &Node{
			value: 8,
			right: &Node{
				value: 11,
			},
		},
	}
	assert.Equal(t, 11, tree2.Max())

	var tree3 *Node = nil
	assert.Equal(t, math.MaxInt64, tree3.Max())
}

func TestBst_Find(t *testing.T) {
	var tree1 *Node = nil
	assert.False(t, tree1.Find(1))
	assert.False(t, tree1.Find(3))

	tree2 := &Node{
		value: 1,
	}
	assert.True(t, tree2.Find(1))
	assert.False(t, tree2.Find(0))

	tree3 := &Node{
		value: 5,
		left: &Node{
			value: 3,
			left: &Node{
				value: 1,
			},
		},
	}
	assert.True(t, tree3.Find(1))

	tree4 := &Node{
		value: 5,
		right: &Node{
			value: 8,
			right: &Node{
				value: 11,
			},
		},
	}
	assert.True(t, tree4.Find(11))
}

func TestBst_Valid(t *testing.T) {
	tree1 := &Node{
		value: 2,
		left: &Node{
			value: 1,
		},
		right: &Node{
			value: 3,
		},
	}

	assert.True(t, tree1.Valid())

	tree2 := &Node{
		value: 5,
		left: &Node{
			value: 1,
		},
		right: &Node{
			value: 4,
			left: &Node{
				value: 3,
			},
			right: &Node{
				value: 6,
			},
		},
	}
	assert.False(t, tree2.Valid())

	//    10
	//   5  15
	//     6 20
	tree3 := &Node{
		value: 10,
		left: &Node{
			value: 5,
		},
		right: &Node{
			value: 15,
			left: &Node{
				value: 6,
			},
			right: &Node{
				value: 20,
			},
		},
	}
	assert.False(t, tree3.Valid())

	tree4 := &Node{
		value: 3,
		left: &Node{
			value: 1,
			left: &Node{
				value: 0,
			},
			right: &Node{
				value: 2,
			},
		},
		right: &Node{
			value: 5,
			left: &Node{
				value: 4,
			},
			right: &Node{
				value: 6,
			},
		},
	}
	assert.True(t, tree4.Valid())
}

func TestBst_Insert(t *testing.T) {
	tree1 := &Node{
		value: 10,
		left: &Node{
			value: 5,
			right: &Node{
				value: 6,
			},
		},
		right: &Node{
			value: 15,
			right: &Node{
				value: 20,
			},
		},
	}
	assert.True(t, tree1.Insert(21))
	assert.True(t, tree1.Valid())
	assert.True(t, tree1.Find(21))

	tree2 := &Node{
		value: 3,
		left: &Node{
			value: 1,
			left: &Node{
				value: 0,
			},
			right: &Node{
				value: 2,
			},
		},
		right: &Node{
			value: 5,
			left: &Node{
				value: 4,
			},
			right: &Node{
				value: 6,
			},
		},
	}
	assert.True(t, tree2.Insert(-1))
	assert.True(t, tree2.Valid())
	assert.True(t, tree2.Find(-1))
}

func TestBst_Delete(t *testing.T) {
	deleted := false

	tree1 := &Node{
		value: 1,
	}
	deleted, tree1 = tree1.Delete(0)
	assert.False(t, deleted)
	assert.True(t, tree1.Valid())

	deleted, tree1 = tree1.Delete(1)
	assert.True(t, deleted)
	assert.Nil(t, tree1)

	tree2 := &Node{
		value: 3,
		left: &Node{
			value: 1,
			left: &Node{
				value: 0,
			},
			right: &Node{
				value: 2,
			},
		},
		right: &Node{
			value: 5,
			left: &Node{
				value: 4,
			},
			right: &Node{
				value: 6,
			},
		},
	}
	deleted, tree2 = tree2.Delete(3)
	assert.True(t, deleted)
	assert.True(t, tree2.Valid())
	assert.Equal(t, 2, tree2.value)

	_, tree2 = tree2.Delete(1)
	assert.True(t, tree2.Valid())
	_, tree2 = tree2.Delete(2)
	assert.True(t, tree2.Valid())
	_, tree2 = tree2.Delete(0)
	assert.True(t, tree2.Valid())
	_, tree2 = tree2.Delete(5)
	assert.True(t, tree2.Valid())
	_, tree2 = tree2.Delete(4)
	assert.True(t, tree2.Valid())
	assert.Equal(t, 6, tree2.value)
}
