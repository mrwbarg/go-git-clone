package object

import (
	"fmt"
)

type ObjectType string

const (
	CommitType ObjectType = "commit"
	TreeType   ObjectType = "tree"
	BlobType   ObjectType = "blob"
	TagType    ObjectType = "tag"
)

type Object interface {
	Serialize() ([]byte, error)
	Deserialize(data []byte)
	Size() int
	Content() []byte
	Type() ObjectType
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

func (o *baseObject) Serialize() ([]byte, error) {
	panic(fmt.Sprintf("Serialize not implemented for object of type %s", string(o.objectType)))
}

func (o *baseObject) Deserialize(data []byte) {
	// In each object that embeds this, the type must also be set.
	o.content = data
}
