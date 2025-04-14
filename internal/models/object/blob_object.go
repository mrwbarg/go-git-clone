package object

type Blob struct {
	baseObject
}

func (b *Blob) Deserialize(data []byte) error {
	b.objectType = BlobType
	return b.baseObject.Deserialize(data)
}
