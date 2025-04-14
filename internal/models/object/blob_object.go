package object

type Blob struct {
	baseObject
}

func (b *Blob) Deserialize(data []byte) {
	b.objectType = BlobType
	b.baseObject.Deserialize(data)
}
