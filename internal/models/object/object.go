package object

import (
	"crypto"
	"encoding/hex"
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
	Serialize() []byte
	Deserialize(data []byte)
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
	return []byte(fmt.Sprintf("%s %d\x00%s", string(o.objectType), len(o.content), o.content))
}

func (o *baseObject) Deserialize(data []byte) {
	// in each object that embeds this, the type must also be set.
	if o.objectType == "" {
		panic("object type not set")
	}
	o.content = data
}

func (o *baseObject) Hash() string {
	h := crypto.SHA1.New()
	h.Write(o.Serialize())
	return hex.EncodeToString(h.Sum(nil))
}
