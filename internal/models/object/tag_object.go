package object

type Tag struct {
	baseObject
}

func (t *Tag) Deserialize(data []byte) {
	t.objectType = TagType
	t.baseObject.Deserialize(data)
}
