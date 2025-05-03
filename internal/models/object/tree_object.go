package object

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type leaf struct {
	mode string
	path string
	sha  string
}

func (leaf *leaf) Parse(data []byte) (int, error) {
	read := 0
	mode, remaining, found := bytes.Cut(data, []byte(" "))
	if !found {
		return 0, fmt.Errorf("invalid leaf format")
	}
	read += len(mode) + 1 // read mode and space

	path, sha, found := bytes.Cut(remaining, []byte("\x00"))
	if !found {
		return 0, fmt.Errorf("invalid leaf format")
	}
	read += len(path) + 1 // read path, null terminator and sha

	leaf.mode = fmt.Sprintf("%06s", string(mode))
	leaf.path = string(path)

	decodedSha := make([]byte, len(sha))
	n, err := binary.Decode(sha, binary.BigEndian, decodedSha)
	if err != nil {
		return 0, err
	}
	read += n

	leaf.sha = fmt.Sprintf("%040x", decodedSha)

	return read, nil
}

type Tree struct {
	baseObject
	leafs []leaf
}

func (t *Tree) Deserialize(data []byte) error {
	t.objectType = TreeType
	t.content = data
	t.leafs = make([]leaf, 0)

	start := 0
	end := bytes.Index(data, []byte("\x00")) + 1 + 20 // null terminator + sha
	if end == -1 {
		return fmt.Errorf("invalid tree format")
	}

	for end <= len(data) {
		leaf := &leaf{}
		read, err := leaf.Parse(data[start:end])
		if err != nil {
			return err
		}
		t.leafs = append(t.leafs, *leaf)
		start = start + read
		end = start + bytes.Index(data[start:], []byte("\x00")) + 1 + 20
	}

	return nil
}
