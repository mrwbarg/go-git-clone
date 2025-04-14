package object

type Tag struct {
	baseObject
}

func (t *Tag) Deserialize(data []byte) error {
	t.objectType = TagType
	return t.baseObject.Deserialize(data)
}
