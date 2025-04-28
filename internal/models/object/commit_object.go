package object

import (
	"fmt"
	"strings"

	"github.com/mrwbarg/go-git-clone/internal/utils"
)

type Commit struct {
	baseObject
	klvmData *utils.OrderedMap[string, []string]
}

func (c *Commit) Deserialize(data []byte) error {
	c.objectType = CommitType
	c.content = data
	c.klvmData = utils.NewOrderedMap[string, []string]()

	err := utils.ParseKVLM(data, c.klvmData)
	if err != nil {
		return fmt.Errorf("fatal: malformed commit data: %w", err)
	}
	return nil
}

func (c *Commit) Content() []byte {
	return []byte(utils.DumpKVLM(c.klvmData))
}

func (c *Commit) Log() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("commit %s\n", c.Hash()))
	builder.WriteString(fmt.Sprintf("Author: %s\n", c.Author()))
	builder.WriteString("\n")
	builder.WriteString("\t")

	messageLines := strings.Split(c.Message(false), "\n")
	for _, line := range messageLines[:len(messageLines)-1] {
		builder.WriteString(line)
		builder.WriteString("\n\t")
	}
	builder.WriteString(messageLines[len(messageLines)-1])
	return builder.String()
}

func (c *Commit) Parent() string {
	if parent, exists := c.klvmData.Get("parent"); exists {
		return parent[0]
	}

	return ""
}

func (c *Commit) Author() string {
	if author, exists := c.klvmData.Get("author"); exists {
		return strings.SplitAfter(author[0], ">")[0]
	}

	return ""
}

func (c *Commit) Message(firstLineOnly bool) string {
	if message, exists := c.klvmData.Get(""); exists {
		if !firstLineOnly {
			return message[0]
		}
		firstLine, _, _ := strings.Cut(message[0], "\n")
		return firstLine
	}

	return ""
}
