package union_find

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnionFind(t *testing.T) {
	u := NewUnion()

	u.Join(1, 2)
	assert.Equal(t, 1, u.Find(1))
	assert.Equal(t, 1, u.Find(2))
	assert.True(t, u.Connected(1, 2))

	u.Join(3, 4)
	u.Join(2, 4)
	assert.True(t, u.Connected(1, 3))
	assert.True(t, u.Connected(2, 4))
	assert.Equal(t, 1, u.Find(3))
	assert.Equal(t, 1, u.Find(4))
}
