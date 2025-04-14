package object

type Tree struct {
	baseObject
}

func (t *Tree) Deserialize(data []byte) {
	t.objectType = TreeType
	t.baseObject.Deserialize(data)
}
