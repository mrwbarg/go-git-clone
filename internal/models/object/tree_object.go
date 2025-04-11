package object

type Tree struct {
	baseObject
}

func (t *Tree) Serialize() ([]byte, error) {
	return nil, nil
}

func (t *Tree) Deserialize(data []byte) {
	t.objectType = TreeType
	t.baseObject.Deserialize(data)
}
