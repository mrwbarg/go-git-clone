package object

type Tag struct {
	baseObject
}

func (t *Tag) Serialize() ([]byte, error) {
	return nil, nil
}

func (t *Tag) Deserialize(data []byte) {
	t.objectType = TagType
	t.baseObject.Deserialize(data)
}
