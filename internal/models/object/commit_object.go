package object

type Commit struct {
	baseObject
}

func (c *Commit) Serialize() ([]byte, error) {
	return nil, nil
}

func (c *Commit) Deserialize(data []byte) {
	c.objectType = CommitType
	c.baseObject.Deserialize(data)
}
