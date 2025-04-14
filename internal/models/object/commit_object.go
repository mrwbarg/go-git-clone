package object

type Commit struct {
	baseObject
}

func (c *Commit) Deserialize(data []byte) error {
	c.objectType = CommitType
	return c.baseObject.Deserialize(data)
}
