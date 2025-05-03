package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Leaf_Parse(t *testing.T) {
	leaf := &leaf{}

	data := []byte("100644 path/to/file.txt\x007815696ecbf1c96e6894")

	read, err := leaf.Parse(data)
	assert.NoError(t, err)

	assert.Equal(t, "100644", leaf.mode)
	assert.Equal(t, "path/to/file.txt", leaf.path)
	assert.Equal(t, "3738313536393665636266316339366536383934", leaf.sha)
	assert.Equal(t, 44, read)
}

func Test_Leaf_Parse_ModeLength5(t *testing.T) {
	leaf := &leaf{}
	data := []byte("10064 path/to/file.txt\x007815696ecbf1c96e6894")
	read, err := leaf.Parse(data)
	assert.NoError(t, err)
	assert.Equal(t, read, 43)

	assert.Equal(t, "010064", leaf.mode)
}

func Test_Tree_Deserialize(t *testing.T) {
	tree := &Tree{}
	data := []byte("100644 path/to/file.txt\x007815696ecbf1c96e689410644 path/to/anotherfile.txt\x007815696ecbf1c96e6894")
	err := tree.Deserialize(data)
	assert.NoError(t, err)

	assert.Equal(t, TreeType, tree.objectType)
	assert.Equal(t, data, tree.content)
	assert.Len(t, tree.leafs, 2)

	assert.Equal(t, "100644", tree.leafs[0].mode)
	assert.Equal(t, "path/to/file.txt", tree.leafs[0].path)
	assert.Equal(t, "3738313536393665636266316339366536383934", tree.leafs[0].sha)
	assert.Equal(t, "010644", tree.leafs[1].mode)
	assert.Equal(t, "path/to/anotherfile.txt", tree.leafs[1].path)
	assert.Equal(t, "3738313536393665636266316339366536383934", tree.leafs[1].sha)
}
