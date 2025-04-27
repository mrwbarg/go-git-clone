package object

import (
	"fmt"

	"github.com/mrwbarg/go-git-clone/internal/utils"
)

type Commit struct {
	baseObject
	content *utils.OrderedMap[string, []string]
}

func (c *Commit) Deserialize(data []byte) error {
	c.objectType = CommitType
	c.content = utils.NewOrderedMap[string, []string]()
	err := utils.ParseKVLM(data, c.content)
	if err != nil {
		return fmt.Errorf("fatal: malformed commit data: %w", err)
	}
	return nil
}

func (c *Commit) Content() []byte {
	return []byte(utils.DumpKVLM(c.content))
}
