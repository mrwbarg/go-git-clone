package object

import (
	"bytes"
	"crypto"
	"encoding/hex"
	"fmt"
	"strconv"
)

type ObjectType string

const (
	CommitType ObjectType = "commit"
	TreeType   ObjectType = "tree"
	BlobType   ObjectType = "blob"
	TagType    ObjectType = "tag"
)

type Object interface {
	Serialize() []byte
	Deserialize(data []byte) error
	Size() int
	Content() []byte
	Type() ObjectType
	Hash() string
}

var _ Object = (*baseObject)(nil)

type baseObject struct {
	objectType ObjectType
	content    []byte
}

func (o *baseObject) FromData(data []byte) error {
	o.content = data
	return nil
}

func (o *baseObject) Size() int {
	return len(o.content)
}

func (o *baseObject) Content() []byte {
	return o.content
}

func (o *baseObject) Type() ObjectType {
	return o.objectType
}

func (o *baseObject) Serialize() []byte {
	return []byte(fmt.Sprintf("%s %d\x00%s", string(o.objectType), len(o.Content()), o.Content()))
}

func (o *baseObject) Deserialize(data []byte) error {
	o.content = data
	return nil
}

func New(data []byte) (*Object, error) {
	format, sizeAndData, _ := bytes.Cut(data, []byte(" "))
	size, content, _ := bytes.Cut(sizeAndData, []byte("\x00"))

	intSize, err := strconv.Atoi(string(size))
	if err != nil {
		return nil, fmt.Errorf("fatal: error parsing object size: %v", err)
	}

	if intSize != len(content) {
		return nil, fmt.Errorf("fatal: invalid object size. Expected: %d. Actual: %d", intSize, len(data))
	}

	var obj Object
	switch ObjectType(format) {
	case CommitType:
		obj = &Commit{}
	case TreeType:
		obj = &Tree{}
	case BlobType:
		obj = &Blob{}
	case TagType:
		obj = &Tag{}
	default:
		return nil, fmt.Errorf("fatal: unknown object type %s", format)
	}

	err = obj.Deserialize(content)
	if err != nil {
		return nil, fmt.Errorf("fatal: error deserializing object: %v", err)
	}
	return &obj, nil
}

func (o *baseObject) Hash() string {
	h := crypto.SHA1.New()
	h.Write(o.Serialize())
	return hex.EncodeToString(h.Sum(nil))
}
