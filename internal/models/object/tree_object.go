package object

type Tree struct {
	baseObject
}

func (t *Tree) Deserialize(data []byte) error {
	t.objectType = TreeType
	return t.baseObject.Deserialize(data)
}
